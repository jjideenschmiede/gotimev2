//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2022 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gotimev2.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gotimev2

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// TimesByMonthReturn is to decode the json data
type TimesByMonthReturn struct {
	Total int     `json:"total"`
	Sum   float64 `json:"sum"`
	Times []struct {
		Id          string `json:"id"`
		Client      string `json:"client"`
		Project     string `json:"project"`
		Time        string `json:"time"`
		Cost        string `json:"cost"`
		Rate        string `json:"rate"`
		Description string `json:"description"`
		Timestamp   string `json:"timestamp"`
		User        string `json:"user"`
	} `json:"times"`
}

// TimesByMonth is to get all times by month
func TimesByMonth(month string, r Request) (TimesByMonthReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/v1/times",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return TimesByMonthReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return TimesByMonthReturn{}, err
	}

	newUrl.Add("month", month)

	// Set new url
	parse.RawQuery = newUrl.Encode()
	c.Path = fmt.Sprintf("%s", parse)

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return TimesByMonthReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode TimesByMonthReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return TimesByMonthReturn{}, err
	}

	// Return data
	return decode, err

}
