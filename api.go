package wrikego

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get(config Config, path string, params url.Values) ([]byte, error) {
	url := config.BaseUrl + path
	// if params != nil {
	// 	url += fmt.Sprintf("?%s", params.Encode())
	// }

	//TODO: Header authroization

	fmt.Println(url)


	client := http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(params.Encode()))
	req.Header.Add("Authorization", "Bearer " + config.PermAccessToken)
	res, err := client.Do(req)
	if err != nil { fmt.Println("Request Error:", err) }
	defer res.Body.Close()


	
	fmt.Println("A")
	response, err := io.ReadAll(res.Body)
	if err != nil { fmt.Println("io.ReadAll Error:", err) }

	fmt.Println(string(response))

	return response, err
}
