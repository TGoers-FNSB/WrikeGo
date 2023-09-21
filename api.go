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

	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("A")

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("B")

	return response, err
}
