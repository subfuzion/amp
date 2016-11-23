package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func (g *swarmClient) cli() {

	RootCmd.PersistentFlags().BoolVarP(&g.verbose, "verbose", "v", false, `Verbose output`)
	cobra.OnInitialize(func() {
		if err := client.init(); err != nil {
			fmt.Printf("Init error: %v\n", err)
			os.Exit(1)
		}
	})

	// versionCmd represents the agrid version
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Display the version number of swarm-server",
		Long:  `Display the version number of swarm-server`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("swarm-server version: %s, build: %s)\n", Version, Build)
		},
	}
	RootCmd.AddCommand(versionCmd)

	//Execute commad
	cmd, _, err := RootCmd.Find(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error during: %s: %v\n", cmd.Name(), err)
		os.Exit(1)
	}

	os.Exit(0)
}
