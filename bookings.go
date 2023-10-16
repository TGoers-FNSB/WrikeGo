package wrikego

import (
	"fmt"
	"net/http"

	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryBookingsByIds(config Config, pathId []string) (resp.Bookings, *http.Response) {
	path := fmt.Sprintf("/bookings/%s", strings.Join(pathId, ","))
	response, httpResponse, err := Get(config, path, nil)
	ErrorCheck(err)
	json, err := resp.BookingsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryBookings(config Config, params params.QueryBookings) (resp.Bookings, *http.Response) {
	path := "/bookings"
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.BookingsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func QueryBookingsByFolder(config Config, params params.QueryBookings, pathId string) (resp.Bookings, *http.Response) {
	path := fmt.Sprintf("/folders/%s/bookings", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Get(config, path, body)
	ErrorCheck(err)
	json, err := resp.BookingsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func CreateBookingsByFolder(config Config, params params.CreateBookings, pathId string) (resp.Bookings, *http.Response) {
	path := fmt.Sprintf("/folders/%s/bookings", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Post(config, path, body)
	ErrorCheck(err)
	json, err := resp.BookingsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func ModifyBookingsById(config Config, params params.ModifyBookings, pathId string) (resp.Bookings, *http.Response) {
	path := fmt.Sprintf("/bookings/%s", pathId)
	body, err := query.Values(params)
	ErrorCheck(err)
	response, httpResponse, err := Put(config, path, body)
	ErrorCheck(err)
	json, err := resp.BookingsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}

func DeleteBookingsById(config Config, pathId string) (resp.Bookings, *http.Response) {
	path := fmt.Sprintf("/bookings/%s", pathId)
	response, httpResponse, err := Delete(config, path, nil)
	ErrorCheck(err)
	json, err := resp.BookingsFromJSON(response)
	ErrorCheck(err)
	return json, httpResponse
}