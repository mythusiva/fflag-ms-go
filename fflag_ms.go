// Package fflag_ms provides functionality to fetch and manage feature flags.
package fflag_ms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiConfigParameters struct {
	Key       string
	Namespace string
	RefreshMS uint   `default:"60000"`
	BaseUrl   string `default:"https://feature-flags2.p.rapidapi.com"`
}

type featureFlags struct {
	config ApiConfigParameters
	data   map[string]any
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

	json.Unmarshal([]byte(string(data)), &f.data)
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

	ff.Fetch()

	return ff
}
