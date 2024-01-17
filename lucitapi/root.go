package lucitapi

import (
	"encoding/json"
	"io/ioutil"
	"lucit-cli/output"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// AuthResponse to /auth
type AuthV2Response struct {
	Success struct {
		Token string `json:"token"`
	}
	Error string `json:"error"`
	Code  string `json:"code"`
}

// AuthResponse to /auth
type AuthResponse struct {
	Ok      bool   `json:"ok"`
	Token   string `json:"token"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// PublicStatusResponse /public/status
type PublicStatusResponse struct {
	Ok            bool   `json:"ok"`
	Timestamp     int    `json:"timestamp"`
	APIVersion    string `json:"api_version"`
	AuthRequired  bool   `json:"auth_required"`
	AppIDRequired bool   `json:"app_id_required"`
	V3AppIDSent   string `json:"v3_app_id_sent"`
	Message       string `json:"message"`
}

// StatusResponse to /status
type StatusResponse struct {
	Ok             bool   `json:"ok"`
	Timestamp      int    `json:"timestamp"`
	APIVersion     string `json:"api_version"`
	AuthRequired   bool   `json:"auth_required"`
	AppIDRequired  bool   `json:"app_id_required"`
	V3AppIDSent    string `json:"v3_app_id_sent"`
	AuthUserIDSent string `json:"auth_user_id_sent"`
	Message        string `json:"message"`
	User           struct {
		Lcuid string `json:"lcuid"`
		Name  string `json:"name"`
	} `json:"user"`
	App struct {
		ID          int         `json:"id"`
		CreatedAt   time.Time   `json:"created_at"`
		UpdatedAt   time.Time   `json:"updated_at"`
		DeletedAt   interface{} `json:"deleted_at"`
		Name        string      `json:"name"`
		Description string      `json:"description"`
		Permissions struct {
			Allowed []string `json:"allowed"`
		} `json:"permissions"`
		Status           int    `json:"status"`
		CreatedByUserID  int    `json:"created_by_user_id"`
		ApplicationClass string `json:"application_class"`
		Options          struct {
			Init []interface{} `json:"_init"`
		} `json:"options"`
		Lcuid  string `json:"lcuid"`
		Slug   string `json:"slug"`
		HashID string `json:"hash_id"`
		Lid    string `json:"lid"`
	} `json:"app"`
}

func getApiUrl(endpoint string) string {

	lucitUrl := viper.GetString("lucit_api_url")

	return lucitUrl + endpoint
}

func getApiClientGetResponse(endpoint string) (responseData []byte, err error) {

	apiUrl := getApiUrl(endpoint)

	client := &http.Client{}

	req, err := http.NewRequest("GET", apiUrl, nil)

	output.InfoIfVerbose("Using v3 token")
	req.Header.Add("AppIdV3", viper.GetString("lucit_app_id"))

	if viper.GetString("lucit_oauth_token") != "" {
		req.Header.Add("Authorization", "Bearer "+viper.GetString("lucit_oauth_token"))
	}

	response, err := client.Do(req)

	if err != nil {
		output.Error(err.Error())
		return
	}

	OutputCommonErrors(response)

	responseData, err = ioutil.ReadAll(response.Body)

	if err != nil {
		output.Error(err.Error())
		return
	}

	return
}

func getApiClientPostResponse(endpoint string, params url.Values) (responseData []byte, err error) {

	apiUrl := getApiUrl(endpoint)

	client := &http.Client{}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(params.Encode()))

	output.InfoIfVerbose("Using v3 token")
	req.Header.Add("AppIdV3", viper.GetString("lucit_app_id"))

	if viper.GetString("lucit_oauth_token") != "" {
		req.Header.Add("Authorization", "Bearer "+viper.GetString("lucit_oauth_token"))
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(req)

	if err != nil {
		output.Error(err.Error())
		return
	}

	OutputCommonErrors(response)

	responseData, err = ioutil.ReadAll(response.Body)

	if err != nil {
		output.Error(err.Error())
		return
	}

	return responseData, nil
}

func OutputCommonErrors(response *http.Response) {

	output.InfoIfVerbose("Status Code : " + response.Status)

	if response.StatusCode == 401 {

		responseData401, _ := ioutil.ReadAll(response.Body)
		output.ErrorDescriptive(response.Status, string(responseData401))
		output.Info(response.Request.URL.RequestURI())
	}

	if response.StatusCode == 404 {
		output.ErrorDescriptive(response.Status, response.Request.URL.RequestURI())
	}

	if response.StatusCode == 405 {
		output.ErrorDescriptive(response.Status, response.Request.URL.RequestURI())
	}

	if response.StatusCode >= 300 {
		output.ErrorDescriptive(response.Status, "Run with --verbose for more information")
		output.Info(response.Request.URL.RequestURI())
	}

}

// Endpoint Returns an API Response as a JSON string
func Endpoint(endpoint string) (responseJSONString string, err error) {

	responseData, err := getApiClientGetResponse(endpoint)

	if err != nil {
		return
	}

	if err != nil {
		output.Error(err.Error())
		return
	}

	output.InfoIfVerbose("Response Data : " + string(responseData))

	responseJSONString = string(responseData)

	return
}

func PublicStatus() (responseObject PublicStatusResponse, err error) {

	responseData, err := getApiClientGetResponse("/public/status")

	if err != nil {
		return
	}

	if err != nil {
		output.Error(err.Error())
		return
	}

	output.InfoIfVerbose("Response Data : " + string(responseData))

	json.Unmarshal(responseData, &responseObject)

	return

}

func Auth() (responseObject AuthResponse, err error) {

	params := url.Values{}
	params.Add("token", viper.GetString("lucit_app_token"))
	params.Add("secret", viper.GetString("lucit_app_secret"))

	responseData, err := getApiClientPostResponse("/auth", params)

	if err != nil {
		return
	}

	if err != nil {
		output.Error(err.Error())
		return
	}

	output.InfoIfVerbose("Response Data : " + string(responseData))

	json.Unmarshal(responseData, &responseObject)

	return

}

func AuthV2(password string) (responseObject AuthV2Response, err error) {

	params := url.Values{}
	params.Add("email", viper.GetString("lucit_app_v2_email"))
	params.Add("password", password)

	output.Warn("AuthV2 is using /login-admin, this will be deprecated in the future. Please use /login instead")

	responseData, err := getApiClientPostResponse("/login-admin", params)

	if err != nil {
		return
	}

	if err != nil {
		output.Error(err.Error())
		return
	}

	output.InfoIfVerbose("Response Data : " + string(responseData))

	json.Unmarshal(responseData, &responseObject)

	return

}

func Status() (responseObject StatusResponse, err error) {

	responseData, err := getApiClientGetResponse("/status")

	if err != nil {
		return
	}

	if err != nil {
		output.Error(err.Error())
		return
	}

	output.InfoIfVerbose("Response Data : " + string(responseData))

	json.Unmarshal(responseData, &responseObject)

	return

}
