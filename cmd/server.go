package cmd

import (
	"github.com/spf13/cobra"
	"grpc-demo/server"
	"log"
)

var StartCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC gateway server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover error:%v \n", err)
			}
		}()
		err := server.Server()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	StartCmd.Flags().StringVarP(&server.Port, "port", "p", "50001", "server port")
	StartCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "./certs/server.pem", "cert pem path")
	StartCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "./certs/server.key", "cert key path")
	StartCmd.Flags().StringVarP(&server.CertName, "cert-name", "", "grpc.demo", "server's hostname")

}
