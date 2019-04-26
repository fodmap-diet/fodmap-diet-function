package function

import (
	"encoding/json"
	"fmt"
	"log"
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

// Handle a serverless request
func Handle(req []byte) string {
	var ip Input

	err := json.Unmarshal(req, &ip)
	if err != nil {
		log.Printf("Invalid input, error %v", err)
		return fmt.Sprintf("Error: Failed to parse input : %v \n %s", err, help)
	}

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
