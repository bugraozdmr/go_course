package sites

import (
	"encoding/json"
	"fiyathunter/helpers"
	response "fiyathunter/response"
	fetch "fiyathunter/response/epic"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func FetchData(keyword string) response.EpicResponse {
	keyword = strings.ToLower(strings.ReplaceAll(keyword, " ", "+"))

	params := BuildVariables(keyword)

	fullURL := helpers.GetUrl("EPIC_FULL_URL") + params

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

	var result fetch.EpicFetch
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("JSON error:", err)
	}

	epicResponse := helpers.EpicBindData(result)

	return epicResponse
}
