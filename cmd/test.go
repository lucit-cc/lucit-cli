package cmd

import (
	"lucit-cli/lucitapi"
	"lucit-cli/output"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test your Lucit CLI API Configuration",
	Long:  `Runs a test suite against your Lucit CLI API Configuration.`,
	Run: func(cmd *cobra.Command, args []string) {

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			output.ErrorDescriptive("Error Reading config file:", viper.ConfigFileUsed())
		}

		testConfigFileExists()
		testConfigFileParams()
		testApiIsReachable()
		testApiIsReachableAndAuthenticated()
		//testOutputFunctions()

		output.Break()
		output.Notable("Test Complete")
		output.Break()

	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func testApiIsReachable() {

	output.Header("Lucit API Reachability")

	publicStatus, err := lucitapi.PublicStatus()

	if err != nil {
		output.FailDescriptive("Error getting public status", err.Error())
		return
	}

	output.Pass("Public Status : " + publicStatus.Message)

}

func testApiIsReachableAndAuthenticated() {

	output.Header("Lucit Authenticated API Test")

	status, err := lucitapi.Status()

	if err != nil {
		output.FailDescriptive("Error getting status", err.Error())
		return
	}

	output.KeyValue("api_version", status.APIVersion)
	output.KeyValue("user.lcuid", status.User.Lcuid)
	output.KeyValue("user.name", status.User.Name)
	output.KeyValue("app.lcuid", status.App.Lcuid)
	output.KeyValue("app.slug", status.App.Slug)
	output.KeyValue("app.name", status.App.Name)
	output.KeyValue("app.description", status.App.Description)
	output.KeyValue("app.application_class", status.App.ApplicationClass)
	output.KeyValue("app.status", strconv.Itoa(status.App.Status))
	output.KeyValue("app.permissions", strings.Join(status.App.Permissions.Allowed, ", "))
}

func testConfigFileParams() {

	output.Header("Config File Params")

	params := []string{"lucit_api_url", "lucit_app_id", "lucit_app_token", "lucit_app_secret", "lucit_oauth_token"}

	for i := 0; i < len(params); i++ {
		if viper.GetString(params[i]) == "" {
			output.Fail(params[i])
		} else {
			output.Pass(params[i])
		}
	}

}

func testConfigFileExists() {

	output.Header("Config File Existence")

	if _, err := os.Stat(viper.ConfigFileUsed()); err == nil {
		output.Pass("config file : " + viper.ConfigFileUsed())
	} else {
		output.Fail("config file : " + viper.ConfigFileUsed())
		output.Info("Create a new config file by running the 'lucit-cli init' command")
	}

}

func testOutputFunctions() {
	output.Header("Testing some Output Functions")
	output.Info("Testing the Lucit API")
	output.Warn("This is something to optionally worry about")
	output.Error("THis is bad")
	output.ErrorDescriptive("This is really bad", "Here is a more detailed description of the error that has just occurred")
	output.Pass("This test is good")
	output.Fail("This one did not work")
	output.PrettyJSON("{ \"test\": \"json\"}")
}
