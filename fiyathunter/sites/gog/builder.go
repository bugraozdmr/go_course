package sites

import (
	"fmt"
	"net/url"
)

func BuildVariables(keyword string) string {
	params := url.Values{}
	params.Add("limit", "48")
	params.Add("query", fmt.Sprintf("like:%s", keyword))
	params.Add("order", "desc:score")
	params.Add("productType", "in:game,pack,dlc,extras")
	params.Add("page", "1")
	params.Add("countryCode", "TR")
	params.Add("locale", "en-US")
	params.Add("currencyCode", "USD")

	return params.Encode()
}