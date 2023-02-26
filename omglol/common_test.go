package omglol

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

// time in seconds to sleep before running each test. Avoids hitting rate limits
const initPause int64 = 1

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

// Add sleep to tests to avoid hitting rate limits when accessing the API
func sleep() {
    time.Sleep(time.Duration(initPause) * time.Second)
}

// Generates a UID from the Github Workflow if present, otherwise generates a random string. This UID can then be used to prevent collision between test runs.
func generateRunUID() string {
	RunUID := os.Getenv("GITHUB_RUN_ID") + os.Getenv("GITHUB_RUN_ATTEMPT")
	if RunUID == "" {
		RunUID = fmt.Sprintf("ts%d", time.Now().Unix())
	}
	return RunUID
}

var RunUID = generateRunUID()

func isOneOf(target string, list []string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}

func testTimestamps(t *testing.T, unix int64, iso8601, rfc2822, relative string) {
	const RFC2822 = "Mon, 02 Jan 2006 15:04:05 -0700"

	u := time.Unix(unix, 0)
	if u.IsZero() {
		t.Errorf("Invalid UnixEpochTime: %d", unix)
	}
	i8601, err := time.Parse(time.RFC3339, iso8601)
	if err != nil {
		t.Errorf("Invalid Iso8601Time: %s", iso8601)
	}
	if u.Unix() != i8601.Unix() {
		t.Errorf("UnixEpochTime: %d, does not match Iso8601Time: %d", u.Unix(), i8601.Unix())
	}
	r2822, err := time.Parse(RFC2822, rfc2822)
	if err != nil {
		t.Errorf("Invalid Rfc2822Time: %s, %e", rfc2822, err)
	}
	if u.Unix() != r2822.Unix() {
		t.Errorf("UnixEpochTime: %d, does not match Rfc2822Time: %d", u.Unix(), i8601.Unix())
	}
	if len(relative) <= 0 {
		t.Errorf("Invalid RelativeTime: %s", relative)
	}
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}