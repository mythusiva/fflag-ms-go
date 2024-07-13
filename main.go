package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiConfigParameters struct {
  key string 
  namespace string
  refreshMS uint `default:"60000"`
  baseUrl string `default:"https://feature-flags2.p.rapidapi.com"`
}

type FeatureFlags struct {
  config ApiConfigParameters
  data map[string]any
}

func (f *FeatureFlags) fetch() {  
  req, err := http.NewRequest("GET", f.config.baseUrl, nil)
  req.Header.Set("x-rapidapi-key", f.config.key)
  req.Header.Set("namespace", f.config.namespace)

  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()

  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  data, err := io.ReadAll(resp.Body)

  if err != nil {
    fmt.Println("Error", err)
    return;
  }

  json.Unmarshal([]byte(string(data)), &f.data)
}

func (f FeatureFlags) get(name string, fallback any) any {
  return f.data[name]
}

func (f FeatureFlags) getAll() map[string]any {
  return f.data
}

func New(params ApiConfigParameters) *FeatureFlags {
  ff := new(FeatureFlags)
  ff.config = params  

  ff.fetch()

  return ff
}

func main() {}

