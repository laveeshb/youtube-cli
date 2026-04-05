package pkg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var playlistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "Manage playlists",
}

var playlistListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all playlists",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Playlist list: not yet implemented")
		return nil
	},
}

var playlistAddCmd = &cobra.Command{
	Use:   "add <playlist-id> <video-id>",
	Short: "Add a video to a playlist",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Playlist add %s -> %s: not yet implemented\n", args[1], args[0])
		return nil
	},
}

func init() {
	playlistCmd.AddCommand(playlistListCmd)
	playlistCmd.AddCommand(playlistAddCmd)
}
