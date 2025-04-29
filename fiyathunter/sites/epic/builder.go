package sites

import (
	"fmt"
	"net/url"
)

func BuildVariables(keyword string) string {
	variables := fmt.Sprintf(`{
		"allowCountries":"TR",
		"category":"games/edition/base|bundles/games|games/edition|editors|addons|games/demo|software/edition/base|games/experience|subscription",
		"count":40,
		"country":"TR",
		"keywords":"%s",
		"locale":"en-US",
		"sortBy":"relevancy,viewableDate",
		"sortDir":"DESC,DESC",
		"tag":"",
		"withPrice":true
	}`, keyword)

	params := url.Values{}
	params.Add("operationName", "searchStoreQuery")
	params.Add("variables", variables)
	params.Add("extensions", buildExtensions())

	return params.Encode()
}

func buildExtensions() string {
	return `{
		"persistedQuery": {
			"version": 1,
			"sha256Hash": "7d58e12d9dd8cb14c84a3ff18d360bf9f0caa96bf218f2c5fda68ba88d68a437"
		}
	}`
}