package cmd

import (
	"fmt"
	"lucit-cli/lucitapi"
	"lucit-cli/output"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes your Lucit CLI API Configuration",
	Long:  `This command will generate a configuration file that contains your Application ID, token and secret`,
	Run: func(cmd *cobra.Command, args []string) {

		output.Header("Lucit CLI API Configuration Setup")
		output.Info("Welcome to the Lucit CLI API Configuration setup wizard")
		output.Info("This wizard will help you configure your Application ID, token and secret")
		output.Break()
		output.Info("Before continuing, you should have an Application ID, a Token and a Secret ready")

		viper.SetDefault("lucit_api_url", "https://api.lucit.app/api/v3")

		output.Break()
		output.Notable("The Lucit API URL is set to " + viper.GetString("lucit_api_url"))
		output.Break()

		err := viper.WriteConfig()

		if err != nil {
			output.FailDescriptive("Error writing config file", err.Error())
			return
		} else {
			output.Pass("Config file written to " + viper.ConfigFileUsed())
		}

		output.Break()

		initV3()

	},
}

func initV3() {

	promptAndSetString("Enter your Lucit Application ID", "lucit_app_id")
	output.Break()
	promptAndSetString("Enter your Lucit Token", "lucit_app_token")
	output.Break()
	promptAndSetString("Enter your Lucit Token Secret", "lucit_app_secret")

	output.Break()
	output.Break()

	err := viper.WriteConfig()

	if err != nil {
		output.FailDescriptive("Error writing config file", err.Error())
		return
	} else {
		output.Pass("Config file written to " + viper.ConfigFileUsed())
	}

	output.Notable("Authenticating with Lucit API")

	authResponse, err := lucitapi.Auth()

	if err != nil {
		output.FailDescriptive("Error authenticating with Lucit API", err.Error())
		return
	}

	if authResponse.Ok == false {
		output.FailDescriptive("Error authenticating with Lucit API", authResponse.Message)
		return
	}

	output.Pass("Authenticated")

	viper.Set("lucit_oauth_token", authResponse.Token)

	err = viper.WriteConfig()

	if err != nil {
		output.FailDescriptive("Error writing auth token to file", err.Error())
		return
	} else {
		output.Pass("Auth token written to config file")
	}
}

func promptAndSetString(label string, key string) {

	fmt.Println(label + " : ")

	var result string

	fmt.Scanln(&result)

	viper.Set(key, result)

}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
