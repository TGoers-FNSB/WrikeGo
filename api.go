package wrikego

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get(config *Config, path string, params url.Values) ([]byte, error) {
	url := config.BaseUrl + path
	if params != nil { url += fmt.Sprintf("?%s", params.Encode())}

	res, err := http.NewRequest("GET", url, nil)
	if err != nil { return nil, err }

	return io.ReadAll(res.Body)
}