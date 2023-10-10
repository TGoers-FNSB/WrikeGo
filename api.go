package wrikego

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
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

func Delete(config Config, path string, params url.Values) ([]byte, error) {
	return api("DELETE", config, path, params)
}

func Download(config Config, path string, params url.Values) ([]byte, error) {
	return api("GET", config, path, params)
}

func Upload(config Config, path string, params params.UploadAttachment) ([]byte, error) {
	return apiAttachment("POST", config, path, params)
}

func Update(config Config, path string, params params.UploadAttachment) ([]byte, error) {
	return apiAttachment("PUT", config, path, params)
}

func api(method string, config Config, path string, params url.Values) ([]byte, error) {
	url := config.BaseUrl + path
	if params != nil {
		url += fmt.Sprintf("?%s", params.Encode())
	}

	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", "Bearer " + config.PermAccessToken)
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

func apiAttachment(method string, config Config, path string, params params.UploadAttachment) ([]byte, error) {
	url := config.BaseUrl + path
	if params.Url != nil {
		url += fmt.Sprintf("?url=%s", *params.Url)
	}

	client := http.Client{}
	req, err := http.NewRequest(method, url, params.DataBinary)
	req.Header.Add("Authorization", "Bearer " + config.PermAccessToken)
	req.Header.Add("content-type", "application/octet-stream")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("X-File-Name", params.FileName)
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

// func Download(config Config, path string, params url.Values) ([]byte, error) {
// 	url := config.BaseUrl + path
// 	if params != nil {
// 		url += fmt.Sprintf("?%s", params.Encode())
// 	}

// 	client := http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	req.Header.Add("Authorization", "Bearer " + config.PermAccessToken)
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Request Error:", err)
// 	}
// 	defer res.Body.Close()

// 	response, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("io.ReadAll Error:", err)
// 	}

// 	return response, err
// }