//package ch2

package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {

	sliceOfSlices := [][]byte{}
	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		sliceOfSlices = append(sliceOfSlices, data)

	}
	return sliceOfSlices, nil

}
