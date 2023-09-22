package wrikego

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get(config Config, path string, params url.Values) ([]byte, error) {
	return api("GET", config, path, params)
}

func Post(config Config, path string, params url.Values) ([]byte, error) {
	return api("POST", config, path, params)
}

func Put(config Config, path string, params url.Values) ([]byte, error) {
	return api("PUT", config, path, params)
}

func Delete(config Config, path string) ([]byte, error) {
	return api("DELETE", config, path, nil)
}

func api(method string, config Config, path string, params url.Values) ([]byte, error) {
	url := config.BaseUrl + path
	if params != nil {
		url += fmt.Sprintf("?%s", params.Encode())
	}

	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "Bearer "+config.PermAccessToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request Error:", err)
	}
	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("io.ReadAll Error:", err)
	}
	return response, err
}
