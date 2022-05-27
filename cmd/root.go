package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run the gRPC gateway server",
}

func Execute() {
	rootCmd.AddCommand(StartCmd)
	// rootCmd：表示在没有任何子命令的情况下的基本命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
