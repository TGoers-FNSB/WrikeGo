package wrikego

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
)

//TODO: (1) Test list of custom structs.
//TODO: (2) Remove * from all parameter datatypes.
//TODO: (3) Add the keyword "-slice" to all slices.

func wrikeEncode(params url.Values) string {
	if params != nil {
		fmt.Println(params) //! Remove

		newObj := ""
		newObjAddCnt := 0
		for i, j := range params {
			fmt.Println(i, j, reflect.TypeOf(j[0]))

			//* Objects
			if strings.Contains(i, "[") {
				nameSplit := strings.Split(i, "[")
				newObj += `"` + strings.Split(nameSplit[1], "]")[0] + `":"` + strings.Join(j, "") + `",`
				params.Del(i)

				// If key exists, set value to new newObj value. Else, add a new key and value.
				if params.Has(nameSplit[0]) {
					params.Set(nameSplit[0], `{`+newObj[:len(newObj)-1]+`}`)
				} else {
					// Increate count by one. If count > 0 then a new key value pair will not be added.
					if newObjAddCnt == 0 {
						params.Add(nameSplit[0], `{`+newObj[:len(newObj)-1]+`}`)
						newObjAddCnt += 1
					} else {
						// If count > 0, reset newObj
						newObj = ""
					}
				}
			} else {
				// Reset newObj and newObjAddCnt upon completing object
				newObj = ""
				newObjAddCnt = 0
			}

			//* Lists
			if strings.Contains(i, "-sliceStruct") {
				params.Add(strings.Split(i, "-sliceStruct")[0], `[` + strings.Join(j, `,`) + `]`)
				params.Del(i)
			} else if strings.Contains(i, "-slice") {
				params.Add(strings.Split(i, "-slice")[0], `["` + strings.Join(j, `","`) + `"]`)
				params.Del(i)
			}
		}
		fmt.Println(params) //! Remove

		return fmt.Sprintf("?%s", params.Encode())
	} else {
		return ""
	}
}

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
	url := config.BaseUrl + path + "?" + params.Encode() // wrikeEncode(params)
	fmt.Println(url) //! Remove

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

func apiAttachment(method string, config Config, path string, params params.UploadAttachment) ([]byte, error) {
	url := config.BaseUrl + path
	if params.Url != nil {
		url += fmt.Sprintf("?url=%s", *params.Url)
	}

	client := http.Client{}
	req, err := http.NewRequest(method, url, params.DataBinary)
	req.Header.Add("Authorization", "Bearer "+config.PermAccessToken)
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
