package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Tries to login to the rest-api.",
	Long: `Tries to execute the rest-api login flow, using email and password to authenticate. 

Mind the rate limiting of rocketchat. Use the waitTime argument to extend time between tries.`,
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
	RootCmd.AddCommand(pingCmd)

	pingCmd.Flags().IntVarP(&maxAttempts, "maxAttempts", "m", 5, "maximum numbers of retries (default: 5)")
	pingCmd.Flags().IntVarP(&waitTimeInSec, "waitTime", "w", 5, "time to wait between tries in seconds (default: 5)")
}
