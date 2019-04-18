package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = apiController.Ping(maxAttempts, time.Duration(waitTimeInSec)*time.Second, Verbose)
		if err != nil {
			return
		}
		fmt.Println("success")
		return
	},
}

var maxAttempts int
var waitTimeInSec int

func init() {
	rootCmd.AddCommand(pingCmd)

	pingCmd.Flags().IntVarP(&maxAttempts, "maxAttempts", "m", 5, "maximum numbers of retries (default: 5)")
	pingCmd.Flags().IntVarP(&waitTimeInSec, "waitTime", "w", 5, "time to wait between tries in seconds (default: 5)")
}
