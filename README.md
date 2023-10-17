# WrikeGo ##
WrikeGo is a Go library API wrapper for the [Wrike.com](https://www.wrike.com/) API.

## Getting Started ##

### Import WrikeGo ##
```Go
import (
    wrike "github.com/TGoers-FNSB/WrikeGo"
    params "github.com/TGoers-FNSB/WrikeGo/parameters"
)
```
### Initialize API ##
Every Wrike API call requires a configuration parameter which defines the base URL, the permanent access token, and SSL Verify. This function returns a configuration struct required for all Wrike API calls.
```Go
config := wrike.NewConfig(base_url string, perm_access_token string, ssl_verify bool)
```
-   The  **base_url**  parameter is the base URL for every Wrike API request.
-   The  **perm_access_token**  parameter is the API key that you generate from your Wrike account (Apps & Integrations > API).
-   The  **ssl_verify**  parameter is a security feature that comes with secure URLs.

### Create Payload ##
To create a payload, use parameters. Each parameter struct starts with the keyword "Query," "Create," "Modify," "Delete," or "Download." These keywords follow the structure of the Wrike API. The second keyword used is the name of the method such as "Tasks" or "Folders." The final keyword in the naming convention is if any Ids are used in the URL such as "taskId" or "folderId." You can [view all parameters in the parameters folder](https://github.com/TGoers-FNSB/WrikeGo/tree/main/parameters).
```Go
payload := params.QueryFolders{
		CustomField: params.CustomField{
			Id:    "ABC123",
			Value: "val1",
		},
	}
```
### Identify Id(s) ##
Some requests require an Id or multiple Ids. In this case, the pathId field is used. This field is either a ```string``` variable or a ```[]string``` variable. The later is used when multiple ids may be used in the request.

### Send Request ##
The following is an example of a request which includes the three variables, config, payload, and an id.
```Go
response, httpresponse := wrike.QueryFoldersByFolder(config, payload, "IEACSABC123")
```
### Response ##
There are 2 variables that each request responds with. The first is the struct that is formatted from the json response. The second is the "*http.Response" object that is returned from each request using Go's "net/http" library. These variables are denoted above by the "response" and "httpresponse" variables.

### Errors ##
Any http status codes >= 400 are returned as errors by the Wrike API. These errors are coded into the "response" variable as "ErrorDescription" and "Error". Any other errors are "panicked" by the WrikeGo library.
