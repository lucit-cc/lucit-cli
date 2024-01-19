package cmd

import (
	"lucit-cli/lucitapi"
	"lucit-cli/output"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Send a GET request to the Lucit API and output a JSON string",
	Long:  `Send raw GET requests to the Lucit API`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			output.ErrorDescriptive("Error Reading config file:", viper.ConfigFileUsed())
		}

		output.Notable("Sending GET request to " + args[0])

		start := time.Now()

		jsonResponse, _ := lucitapi.Endpoint(args[0])

		duration := time.Since(start)

		output.PrettyJSON(jsonResponse)

		output.KeyValue("TIME", duration.String())
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
