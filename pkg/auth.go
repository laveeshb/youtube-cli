package pkg

import (
	"fmt"

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
		fmt.Println("Auth login: not yet implemented")
		return nil
	},
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current authentication status",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Auth status: not yet implemented")
		return nil
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out and remove stored credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Auth logout: not yet implemented")
		return nil
	},
}

func init() {
	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authStatusCmd)
	authCmd.AddCommand(authLogoutCmd)
}
