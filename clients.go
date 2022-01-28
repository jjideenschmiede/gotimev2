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

import "encoding/json"

// ClientsReturn is to decode the json data
type ClientsReturn struct {
	Total   int `json:"total"`
	Clients []struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Rate      string `json:"rate"`
		Lexoffice string `json:"lexoffice"`
	} `json:"clients"`
}

// Clients are to get all clients from timev2
func Clients(r Request) (ClientsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/v1/clients",
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ClientsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ClientsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ClientsReturn{}, err
	}

	// Return data
	return decode, err

}
