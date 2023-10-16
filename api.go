package wrikego

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
)

func Get(config Config, path string, params url.Values) ([]byte, *http.Response, error) {
	return api("GET", config, path, params)
}

func Post(config Config, path string, params url.Values) ([]byte, *http.Response, error) {
	return api("POST", config, path, params)
}

func Put(config Config, path string, params url.Values) ([]byte, *http.Response, error) {
	return api("PUT", config, path, params)
}

func Delete(config Config, path string, params url.Values) ([]byte, *http.Response, error) {
	return api("DELETE", config, path, params)
}

func Download(config Config, path string, params url.Values) ([]byte, *http.Response, error) {
	return api("GET", config, path, params)
}

func Upload(config Config, path string, params params.UploadAttachment) ([]byte, *http.Response, error) {
	return apiAttachment("POST", config, path, params)
}

func Update(config Config, path string, params params.UploadAttachment) ([]byte, *http.Response, error) {
	return apiAttachment("PUT", config, path, params)
}

func api(method string, config Config, path string, params url.Values) ([]byte, *http.Response, error) {
	url := config.BaseUrl + path + "?" + params.Encode()
	fmt.Println(url)

	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	ErrorCheck(err)
	req.Header.Add("Authorization", "Bearer "+config.PermAccessToken)
	res, err := client.Do(req)
	ErrorCheck(err)
	defer res.Body.Close()
	response, err := io.ReadAll(res.Body)
	ErrorCheck(err)

	fmt.Println(string(response))

	return response, res, err
}

func apiAttachment(method string, config Config, path string, params params.UploadAttachment) ([]byte, *http.Response, error) {
	url := config.BaseUrl + path
	if params.Url != "" {
		url += fmt.Sprintf("?url=%s", params.Url)
	}

	client := http.Client{}
	req, err := http.NewRequest(method, url, params.DataBinary)
	ErrorCheck(err)
	req.Header.Add("Authorization", "Bearer "+config.PermAccessToken)
	req.Header.Add("content-type", "application/octet-stream")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("X-File-Name", params.FileName)
	res, err := client.Do(req)
	ErrorCheck(err)
	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	ErrorCheck(err)

	return response, res, err
}