package main

import (
	"encoding/json"
	"fmt"
)

func FormatMap(mapVar interface{}) (string, error) {

	mapFormatted, err := json.MarshalIndent(mapVar, "", "  ")

	if err != nil {

		return "", fmt.Errorf("error trying to format map variable: %s", err.Error())
	}

	return string(mapFormatted), nil

}

// fmt.Printf("\nmap transfers: %+v\n", s.transfers)

// 	fmt.Println("\nStorage Transfers: ", string(b))
