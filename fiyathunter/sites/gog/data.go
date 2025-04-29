package sites

import (
	"encoding/json"
	"fiyathunter/helpers"
	response "fiyathunter/response"
	fetch "fiyathunter/response/gog"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchData(keyword string) response.GogResponse {
	params := BuildVariables(keyword)

	fullURL := helpers.GetUrl("GOG_FULL_URL") + params

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

	var jsonData fetch.GogFetch
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		log.Fatal("JSON parse error:", err)
	}

	GogResponse := helpers.GogBindData(jsonData)

	return GogResponse
}
