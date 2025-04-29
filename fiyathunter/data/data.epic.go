package data

import (
	"encoding/json"
	"fiyathunter/helpers"
	"fiyathunter/response/epic"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func FetchDataEpic(c *gin.Context) {
	keyword := c.Param("keyword")
	keyword = strings.ToLower(strings.ReplaceAll(keyword, " ", "+"))

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

	extensions := `{
		"persistedQuery": {
			"version": 1,
			"sha256Hash": "7d58e12d9dd8cb14c84a3ff18d360bf9f0caa96bf218f2c5fda68ba88d68a437"
		}
	}`

	params := url.Values{}
	params.Add("operationName", "searchStoreQuery")
	params.Add("variables", variables)
	params.Add("extensions", extensions)

	fullURL := helpers.GetUrl("EPIC_FULL_URL") + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "tr-TR,tr;q=0.9")
	req.Header.Set("User-Agent", "Go-http-client")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result response.EpicFetch
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("JSON error:", err)
	}

	epicResponse := helpers.EpicBindData(result)

	c.JSON(http.StatusOK, gin.H{
		"data": epicResponse,
	})
}
