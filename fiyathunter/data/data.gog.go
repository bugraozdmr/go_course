package data

import (
	"encoding/json"
	"fiyathunter/helpers"
	response "fiyathunter/response/gog"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func FetchDataGog(c *gin.Context) {
	keyword := c.Param("keyword")
	//! No replacing
	keyword = strings.ToLower(keyword)

	fmt.Println(keyword)

	params := url.Values{}
	params.Add("limit", "48")
	params.Add("query", fmt.Sprintf("like:%s", keyword))
	params.Add("order", "desc:score")
	params.Add("productType", "in:game,pack,dlc,extras")
	params.Add("page", "1")
	params.Add("countryCode", "TR")
	params.Add("locale", "en-US")
	params.Add("currencyCode", "USD")

	fullURL := "https://catalog.gog.com/v1/catalog?" + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "tr-TR,tr;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")

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

	var jsonData response.GogFetch
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		log.Fatal("JSON parse error:", err)
	}

	GogResponse := helpers.GogBindData(jsonData)

	c.JSON(http.StatusOK, gin.H{
		"data": GogResponse,
	})
}
