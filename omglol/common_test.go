package omglol

import (
	"fmt"
	"os"
	"time"
)

var testEmail = os.Getenv("OMGLOL_USER_EMAIL")
var testKey = os.Getenv("OMGLOL_API_KEY")
var testName = os.Getenv("OMGLOL_USERNAME")
var testOwnedDomain = os.Getenv("OMGLOL_TEST_OWNED_DOMAIN") // some tests will only work if you own the domain you are testing against

func setHostURL() string {
	if hostURL, exists := os.LookupEnv("OMGLOL_API_HOST"); exists {
		return hostURL
	} else {
		return "https://api.omg.lol"
	}
}

var testHostURL = setHostURL()

// Generates a UID from the Github Workflow if present, otherwise generates a random string. This UID can then be used to prevent collision between test runs.
func generateRunUID() string {
	RunUID := os.Getenv("GITHUB_RUN_ID") + os.Getenv("GITHUB_RUN_ATTEMPT")
	if RunUID == "" {
		RunUID = fmt.Sprintf("ts%d", time.Now().Unix())
	}
	return RunUID
}

var RunUID = generateRunUID()
