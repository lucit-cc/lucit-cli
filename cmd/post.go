package cmd

import (
	"lucit-cli/lucitapi"
	"lucit-cli/output"
	"net/url"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Send a POST request to the Lucit API and output a JSON string",
	Long:  `Send raw POST requests to the Lucit API`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			output.ErrorDescriptive("Error Reading config file:", viper.ConfigFileUsed())
		}

		// Get the endpoint from the command line
		endpoint := args[0]

		// Get the parameters from the command line
		params := args[1:]

		// Convert params to url.Values
		urlValues := make(url.Values)
		for _, param := range params {
			// Split each parameter into 1key and value using '='
			kv := lucitapi.SplitParam(param)
			if len(kv) == 2 {
				urlValues.Add(kv[0], kv[1])
			} else {
				output.Error("Invalid parameter format: " + param)
				return
			}
		}

		output.Notable("Sending POST request to " + endpoint)

		start := time.Now()

		// Send the POST request
		jsonResponse, _ := lucitapi.Post(endpoint, urlValues)

		duration := time.Since(start)

		output.PrettyJSON(jsonResponse)

		output.KeyValue("TIME", duration.String())

	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
