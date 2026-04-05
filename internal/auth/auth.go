package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/laveeshb/youtube-cli/internal/config"
)

var scopes = []string{
	"https://www.googleapis.com/auth/youtube",
	"https://www.googleapis.com/auth/yt-analytics.readonly",
}

type credentials struct {
	Installed struct {
		ClientID     string   `json:"client_id"`
		ClientSecret string   `json:"client_secret"`
		RedirectURIs []string `json:"redirect_uris"`
	} `json:"installed"`
}

func loadOAuthConfig(port int) (*oauth2.Config, error) {
	credPath, err := config.CredentialsPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(credPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf(
				"credentials.json not found at %s\n\n"+
					"To set up:\n"+
					"  1. Go to https://console.cloud.google.com/\n"+
					"  2. Create a project and enable YouTube Data API v3 and YouTube Analytics API\n"+
					"  3. Create OAuth 2.0 credentials (Desktop app type)\n"+
					"  4. Download the JSON and save it to: %s",
				credPath, credPath,
			)
		}
		return nil, fmt.Errorf("reading credentials.json: %w", err)
	}

	var creds credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, fmt.Errorf("parsing credentials.json: %w", err)
	}

	return &oauth2.Config{
		ClientID:     creds.Installed.ClientID,
		ClientSecret: creds.Installed.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       scopes,
		RedirectURL:  fmt.Sprintf("http://localhost:%d/callback", port),
	}, nil
}

func Login() error {
	if err := config.EnsureDir(); err != nil {
		return fmt.Errorf("creating config dir: %w", err)
	}

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return fmt.Errorf("starting local server: %w", err)
	}
	port := listener.Addr().(*net.TCPAddr).Port

	oauthCfg, err := loadOAuthConfig(port)
	if err != nil {
		return err
	}

	authURL := oauthCfg.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	fmt.Printf("Opening browser for authentication...\nIf it doesn't open, visit:\n\n  %s\n\n", authURL)
	openBrowser(authURL)

	codeCh := make(chan string, 1)
	errCh := make(chan error, 1)

	server := &http.Server{}
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			errCh <- fmt.Errorf("no code in callback")
			http.Error(w, "Missing code", http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, "<html><body><h2>Authentication successful! You can close this tab.</h2></body></html>")
		codeCh <- code
	})

	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	var code string
	select {
	case code = <-codeCh:
	case err = <-errCh:
		return fmt.Errorf("OAuth callback error: %w", err)
	case <-time.After(5 * time.Minute):
		return fmt.Errorf("timed out waiting for authentication")
	}

	_ = server.Shutdown(context.Background())

	token, err := oauthCfg.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("exchanging auth code: %w", err)
	}

	if err := saveToken(token); err != nil {
		return fmt.Errorf("saving token: %w", err)
	}

	fmt.Println("Authenticated successfully.")
	return nil
}

func saveToken(token *oauth2.Token) error {
	tokenPath, err := config.TokenPath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return os.WriteFile(tokenPath, data, 0600)
}

func loadToken() (*oauth2.Token, error) {
	tokenPath, err := config.TokenPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("not logged in — run 'yt auth login' first")
		}
		return nil, err
	}
	var token oauth2.Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}
	return &token, nil
}

// persistingTokenSource wraps an oauth2.TokenSource and saves refreshed tokens to disk.
type persistingTokenSource struct {
	src oauth2.TokenSource
}

func (p *persistingTokenSource) Token() (*oauth2.Token, error) {
	token, err := p.src.Token()
	if err != nil {
		return nil, err
	}
	_ = saveToken(token) // best-effort persist
	return token, nil
}

func TokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	token, err := loadToken()
	if err != nil {
		return nil, err
	}

	// Use port 0 just to load config; redirect URI isn't used for refresh.
	oauthCfg, err := loadOAuthConfig(0)
	if err != nil {
		return nil, err
	}

	src := oauthCfg.TokenSource(ctx, token)
	return &persistingTokenSource{src: oauth2.ReuseTokenSource(token, src)}, nil
}

type TokenStatus struct {
	LoggedIn bool
	Expiry   time.Time
}

func Status() (*TokenStatus, error) {
	token, err := loadToken()
	if err != nil {
		return &TokenStatus{LoggedIn: false}, nil
	}
	return &TokenStatus{
		LoggedIn: token.Valid(),
		Expiry:   token.Expiry,
	}, nil
}

func Logout() error {
	tokenPath, err := config.TokenPath()
	if err != nil {
		return err
	}

	// Attempt token revocation
	token, err := loadToken()
	if err == nil && token.AccessToken != "" {
		req, _ := http.NewRequest("POST",
			"https://oauth2.googleapis.com/revoke?token="+token.AccessToken, nil)
		client := &http.Client{Timeout: 5 * time.Second}
		_, _ = client.Do(req) // best-effort
	}

	if err := os.Remove(tokenPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("removing token: %w", err)
	}
	return nil
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return
	}
	_ = cmd.Start()
}
