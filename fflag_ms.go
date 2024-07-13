// Package fflag_ms provides functionality to fetch and manage feature flags.
package fflag_ms

// TODO: Auto-fetch functionality
// TODO: Put flags
// TODO: Delete flags
// TODO: Get CDN url
// TODO: Set exposed names
// TODO: Create README file for example usage

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiConfigParameters struct {
	Key       string
	Namespace string
	RefreshMS uint   `json:"60000"`
	BaseUrl   string `json:"https://feature-flags2.p.rapidapi.com/v1/flags"`
}

type featureFlags struct {
	config ApiConfigParameters
	data   map[string]interface{}
}

func (f *featureFlags) Fetch() {
	req, err := http.NewRequest("GET", f.config.BaseUrl, nil)
	req.Header.Set("x-rapidapi-key", f.config.Key)
	req.Header.Set("namespace", f.config.Namespace)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	var response map[string]interface{}
	json.Unmarshal([]byte(string(data)), &response)

	castedMap, ok := response["data"].(map[string]interface{})

	if !ok {
		fmt.Println("Error: feature flags are malformed.")
		return
	}

	f.data = castedMap
}

func (f featureFlags) Get(name string, fallback any) any {
	result := f.data[name]

	if result == nil {
		result = fallback
	}

	return result
}

func (f featureFlags) GetAll() map[string]any {
	return f.data
}

func New(params ApiConfigParameters) *featureFlags {
	ff := new(featureFlags)
	ff.config = params

	// ff.config.RefreshMS = 60000
	// ff.config.BaseUrl = "https://feature-flags2.p.rapidapi.com/v1/flags"

	ff.Fetch()

	return ff
}
