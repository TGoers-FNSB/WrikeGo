# WrikeGo
WrikeGo is a Golang Wrapper for the [Wrike.com](https://www.wrike.com/) API.

## Getting Started

### Import WrikeGo

    import (
	    wrike "github.com/TGoers-FNSB/WrikeGo"
	    params "github.com/TGoers-FNSB/WrikeGo/parameters"
	)
### Initialize API
Every Wrike API call requires a configuration parameter which defines the base URL, the permanent access token, and SSL Verify. This function returns a configuration struct required for all Wrike API calls.

    config := wrike.NewConfig(base_url string, perm_access_token string, ssl_verify bool)
-   The  **base_url**  parameter is the base URL for every Wrike API request. The default value is "[https://www.wrike.com/api/v4](https://www.wrike.com/api/v4)".
-   The  **perm_access_token**  parameter is the API key that you generate from your Wrike account (Apps & Integrations > API).
-   The  **ssl_verify**  parameter is a security feature that comes with secure URLs.

### Wrike API
The format of WrikePy mirrors that of the [Wrike API](https://developers.wrike.com/) documentation methods. WrikePy includes a class for each API method and a function for each API call type. As an example, the following will get a folders and/or projects from within a folder...

