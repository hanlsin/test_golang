package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var printLog bool
var logPath string
var debug bool

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start command: sub command of server command",
	Long: `A longer description of start command:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("serverStartCmd: PersistentPreRunE")
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("serverStartCmd: Run")
		if printLog {
			fmt.Println("serverStartCmd: Server will print log")
		}
		if logPath != "./log.dat" {
			fmt.Println("serverStartCmd: New log file path: ", logPath)
		}
		if debug {
			fmt.Println("serverStartCmd: debug mode on")
		}
		return nil
	},
}

func startCmd() *cobra.Command {
	flags := serverStartCmd.Flags()

	flags.BoolVarP(&printLog, "print-log", "p", false,
		"Whether to print log")
	flags.StringVarP(&logPath, "log-path", "l", "./log.dat",
		"Set a file path for log file")
	flags.BoolVarP(&debug, "debug", "d", false,
		"Whether to start debug mode")

	return serverStartCmd
}
