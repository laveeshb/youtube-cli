package pkg

import (
	"github.com/laveeshb/youtube-cli/internal/api"
	"github.com/laveeshb/youtube-cli/internal/auth"
	"github.com/spf13/cobra"
)

func newClient(cmd *cobra.Command) (*api.Client, error) {
	ts, err := auth.TokenSource(cmd.Context())
	if err != nil {
		return nil, err
	}
	return api.New(cmd.Context(), ts)
}
