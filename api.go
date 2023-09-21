package wrikego

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get(config Config, path string, params url.Values) ([]byte, error) {
	url := config.BaseUrl + path
	if params != nil {
		url += fmt.Sprintf("?%s", params.Encode())
	}

	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil { fmt.Println("Request Error:", err) }
	defer res.Body.Close()

	fmt.Println("A")
	response, err := io.ReadAll(res.Body)
	if err != nil { fmt.Println("io.ReadAll Error:", err) }

	fmt.Println(string(response))

	return response, err
}
