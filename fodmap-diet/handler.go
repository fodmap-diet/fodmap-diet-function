package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	sdk "github.com/fodmap-diet/go-sdk"
)

type Input struct {
	Items []string `json:"items"`
}

const help = `i/p data format:
{
    "items": [
        "banana",
	"chai",
	"cream"
    ]
}`
const DEFAULT_URL = "https://fodmap-diet-238401.appspot.com"

// Handle locally
func handleLocal(ip Input) string {
	items := make(map[string]interface{})

	for _, key := range ip.Items {

		key = strings.ToLower(key)
		if len(key) == 0 {
			log.Printf("Invalid item, key empty")
			continue
		}

		if _, added := items[key]; added {
			log.Printf("skipping duplicate item %s", key)
			continue
		}

		item, err := sdk.SearchItem(key)
		if err != nil {
			items[key] = struct {
				Error string `json: "error"`
			}{
				err.Error(),
			}
			continue
		}
		items[key] = item
	}

	js, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		log.Printf(err.Error())
		return fmt.Sprintf("Error: Failed to marshal output : %v", err)
	}

	return string(js)
}

// Handle with remote API
func handleRemote(ip Input) string {
	url := os.Getenv("api_url")
	if len(url) == 0 {
		url = DEFAULT_URL
	}

	req, _ := http.NewRequest("GET", url+"/search?", nil)

	q := req.URL.Query()
	for _, key := range ip.Items {
		q.Add("item", key)
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error: Failed to request, %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("Error: Failed to request %s, status %d", req.URL.String(), resp.StatusCode)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	return string(responseData)
}

// Check if readonly Fs
func isReadOnly() bool {
	readonly := true
	str := os.Getenv("read_only_fs")
	if strings.ToLower(str) == "false" {
		readonly = false
	}
	return readonly
}

// Handle a serverless request
func Handle(req []byte) string {
	var ip Input

	err := json.Unmarshal(req, &ip)
	if err != nil {
		log.Printf("Invalid input, error %v", err)
		return fmt.Sprintf("Error: Failed to parse input : %v \n %s", err, help)
	}

	js := ""

	if isReadOnly() {
		js = handleRemote(ip)
	} else {
		js = handleLocal(ip)
	}

	return string(js)
}
