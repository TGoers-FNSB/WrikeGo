package wrikego

import (
	"fmt"
	"log"
	"strings"

	params "github.com/TGoers-FNSB/WrikeGo/parameters"
	resp "github.com/TGoers-FNSB/WrikeGo/response"
	query "github.com/TGoers-FNSB/go-querystring-wrike/query"
)

func QueryBookingsByIds(config Config, pathId []string) (resp.Bookings, error) {
	path := fmt.Sprintf("/bookings/%s", strings.Join(pathId, ","))
	response, _ := Get(config, path, nil)
	return resp.BookingsFromJSON(response)
}

func QueryBookings(config Config, params params.QueryBookings) (resp.Bookings, error) {
	path := "/bookings"
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.BookingsFromJSON(response)
}

func QueryBookingsByFolder(config Config, params params.QueryBookings, pathId string) (resp.Bookings, error) {
	path := fmt.Sprintf("/folders/%s/bookings", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Get(config, path, body)
	return resp.BookingsFromJSON(response)
}

func CreateBookingsByFolder(config Config, params params.CreateBookings, pathId string) (resp.Bookings, error) {
	path := fmt.Sprintf("/folders/%s/bookings", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Post(config, path, body)
	return resp.BookingsFromJSON(response)
}

func ModifyBookingsById(config Config, params params.ModifyBookings, pathId string) (resp.Bookings, error) {
	path := fmt.Sprintf("/bookings/%s", pathId)
	body, err := query.Values(params)
	if err != nil {
		log.Println(err)
	}
	response, _ := Put(config, path, body)
	return resp.BookingsFromJSON(response)
}

func DeleteBookingsById(config Config, pathId string) (resp.Bookings, error) {
	path := fmt.Sprintf("/bookings/%s", pathId)
	response, _ := Delete(config, path, nil)
	return resp.BookingsFromJSON(response)
}