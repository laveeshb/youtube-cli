package pkg

import (
	"fmt"

	"github.com/laveeshb/youtube-cli/internal/auth"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with your Google account",
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to your Google account via OAuth",
	RunE: func(cmd *cobra.Command, args []string) error {
		return auth.Login()
	},
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current authentication status",
	RunE: func(cmd *cobra.Command, args []string) error {
		status, err := auth.Status()
		if err != nil {
			return err
		}
		if !status.LoggedIn {
			fmt.Println("Not logged in. Run 'yt auth login' to authenticate.")
			return nil
		}
		fmt.Printf("Logged in\nToken expires: %s\n", status.Expiry.Local().Format("2006-01-02 15:04:05"))
		return nil
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out and remove stored credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := auth.Logout(); err != nil {
			return err
		}
		fmt.Println("Logged out.")
		return nil
	},
}

func init() {
	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authStatusCmd)
	authCmd.AddCommand(authLogoutCmd)
}
