package cmd

import (
	"github.com/sjmshsh/HubuIM/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use: "client",
	Run: ClientHandle,
}

func ClientHandle(cmd *cobra.Command, args []string) {
	client.RunMain()
}
