package cmd

import (
	"lucit-cli/lucitapi"
	"lucit-cli/output"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes your Lucit CLI API Configuration",
	Long:  `This command will generate a configuration file that contains the Lucit API url, tokens and secrets`,
	Run: func(cmd *cobra.Command, args []string) {

		output.Header("Lucit CLI API Configuration Setup")
		output.Info("Welcome to the Lucit CLI API Configuration setup wizard")
		output.Info("This wizard will help you configure your Lucit API url, tokens and secrets")
		output.Info("Before continuing, you should have an App ID, a Bot Token and a Bot Secret")

		output.Break()

		viper.SetDefault("lucit_api_url", "https://api.lucit.app/api")

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

	promptAndSetString("Enter your Lucit App ID", "lucit_app_id", viper.GetString("lucit_app_id"))
	promptAndSetString("Enter your Lucit Bot Token", "lucit_bot_token", viper.GetString("lucit_bot_token"))
	promptAndSetString("Enter your Lucit Bot Secret", "lucit_bot_secret", viper.GetString("lucit_bot_secret"))

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

	viper.Set("lucit_bot_auth_token", authResponse.Token)

	err = viper.WriteConfig()

	if err != nil {
		output.FailDescriptive("Error writing auth token to file", err.Error())
		return
	} else {
		output.Pass("Auth token written to config file")
	}
}

func promptAndSetChoice(label string, key string, choices []string) {

	prompt := promptui.Select{
		Label: label,
		Items: choices,
	}

	_, result, err := prompt.Run()

	if err != nil {
		output.Error(err.Error())
		return
	}

	viper.Set(key, result)

}

func promptAndSetString(label string, key string, defaultValue string) {

	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}

	result, err := prompt.Run()

	if err != nil {
		output.Error(err.Error())
		return
	}

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
