package output

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/TylerBrock/colorjson"
)

var VerboseOutput bool

func SetVerbose() {
	VerboseOutput = true
	Break()
	Notable("Verbose output enabled")
	Break()
}

func Info(message string) {
	fmt.Println("  " + message)
}

func InfoIfVerbose(message string) {
	if !VerboseOutput {
		return
	}

	fmt.Println("  " + message)
}

func Notable(message string) {
	fmt.Println(" ", color.OverBlue(" NOTE "), message)
}

func KeyValue(key string, value string) {
	fmt.Println(" ", color.OverBlue(" "+key+" "), value)
}

func Header(message string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(" ", color.OverPurple("        "+strings.ToUpper(message)+"        "))
	fmt.Println("")

}

func Warn(message string) {
	fmt.Println(" ", color.OverYellow(" WARN "), message)
}

func Error(message string) {
	fmt.Println(" ", color.OverRed(" ERROR "), message)
}

func ErrorDescriptive(message string, moreInformation string) {
	fmt.Println("")
	fmt.Println(" ", color.OverRed(" ERROR "), message)
	Info("   " + moreInformation)
	fmt.Println("")
}

func Pass(message string) {
	fmt.Println(" ", color.OverGreen(" PASS "), message)
}

func Fail(message string) {
	fmt.Println(" ", color.OverRed(" FAIL "), message)
}

func FailDescriptive(message string, moreInformation string) {
	fmt.Println("")
	fmt.Println(" ", color.OverRed(" FAIL "), message)
	Info("   " + moreInformation)
	fmt.Println("")
}

func PrettyJSON(str string) {

	fmt.Println("")

	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)

	f := colorjson.NewFormatter()
	f.Indent = 2

	s, _ := f.Marshal(obj)
	fmt.Println(string(s))

	Break()
}

func Break() {
	fmt.Println("")
	fmt.Println("")
}
