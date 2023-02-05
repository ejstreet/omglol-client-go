package omglol

import (
	"os"
)

var testEmail = os.Getenv("OMGLOL_USER_EMAIL")
var testKey = os.Getenv("OMGLOL_API_KEY")
var testName = os.Getenv("OMGLOL_USERNAME")
var testOwnedDomain = os.Getenv("OMGLOL_TEST_OWNED_DOMAIN") // some tests will only work if you own the domain you are testing against
