package wrikego

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get(config Config, path string, params url.Values) ([]byte, error) {
	url := config.BaseUrl + path
	if params != nil { url += fmt.Sprintf("?%s", params.Encode()) }

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer " + config.PermAccessToken)
	res, err := client.Do(req)
	if err != nil { fmt.Println("Request Error:", err) }
	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil { fmt.Println("io.ReadAll Error:", err) }

	return response, err
}
