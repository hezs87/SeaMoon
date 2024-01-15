package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/DVKunion/SeaMoon/pkg/client"
	"github.com/DVKunion/SeaMoon/pkg/consts"
	"github.com/DVKunion/SeaMoon/server"
)

var (
	debug   bool
	verbose bool

	// server params
	addr  string
	port  string
	proto string

	rootCommand = &cobra.Command{}

	serverCommand = &cobra.Command{
		Use:   "server",
		Short: "seaMoon server mod",
		RunE:  Server,
	}

	proxyCommand = &cobra.Command{
		Use:   "proxy",
		Short: "seaMoon proxy mod",
		Run:   Proxy,
	}

	versionCommand = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(consts.Version)
		},
	}
)

func Proxy(cmd *cobra.Command, args []string) {
	client.Serve(cmd.Context(), verbose, debug)
}

func Server(cmd *cobra.Command, args []string) error {
	s, err := server.New(
		server.WithHost("0.0.0.0"),
		server.WithPort(port),
		server.WithProto(proto),
	)

	if err != nil {
		return err
	}

	return s.Serve(cmd.Context())
}

func init() {
	proxyCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "proxy detail log")
	proxyCommand.Flags().BoolVarP(&debug, "debug", "d", false, "proxy detail log")

	serverCommand.Flags().StringVarP(&addr, "addr", "a", "0.0.0.0", "server listen addr")
	serverCommand.Flags().StringVarP(&port, "port", "p", "9000", "server listen port")
	serverCommand.Flags().StringVarP(&proto, "proto", "t", "websocket", "server listen proto: (websocket/grpc)")

	rootCommand.AddCommand(versionCommand)
	rootCommand.AddCommand(proxyCommand)
	rootCommand.AddCommand(serverCommand)
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}