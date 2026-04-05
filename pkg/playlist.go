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
		client, err := newClient(cmd)
		if err != nil {
			return err
		}

		playlists, err := client.ListPlaylists(cmd.Context())
		if err != nil {
			return err
		}

		if len(playlists) == 0 {
			fmt.Println("No playlists found.")
			return nil
		}

		for _, p := range playlists {
			fmt.Printf("%s\t%s\t(%d videos)\n", p.Id, p.Snippet.Title, p.ContentDetails.ItemCount)
		}
		return nil
	},
}

var playlistAddCmd = &cobra.Command{
	Use:   "add <playlist-id> <video-id>",
	Short: "Add a video to a playlist",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient(cmd)
		if err != nil {
			return err
		}

		if err := client.AddToPlaylist(cmd.Context(), args[0], args[1]); err != nil {
			return err
		}

		fmt.Printf("Added video %s to playlist %s\n", args[1], args[0])
		return nil
	},
}

func init() {
	playlistCmd.AddCommand(playlistListCmd)
	playlistCmd.AddCommand(playlistAddCmd)
}
