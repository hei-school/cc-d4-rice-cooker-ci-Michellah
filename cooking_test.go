package main

import (
	"errors"
	"testing"
)

func TestCheckPower(t *testing.T){
	tests := []struct {
		hasPower   bool
		expectErr  bool
		expectedErr error
	}{
		{true, false, nil},
		{false, true, errors.New("No power, unable to cook!")},
	}

	for _, test := range tests {
		err := CheckPower(test.hasPower)

		if (err != nil) != test.expectErr {
			t.Errorf("Expected error: %v, Got error: %v", test.expectErr, err != nil)
		}

		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("Expected error message: %v, Got error message: %v", test.expectedErr, err)
		}
	}
}

func TestAddWater(t *testing.T) {
	tests := []struct {
		water         string
		expectErr     bool
		expectedErr   error
	}{
		{"", true, errors.New("Please specify the amount of water!")},
		{"250ml", false, nil},
	}

	for _, test := range tests {
		err := AddWater(test.water)

		if (err != nil) != test.expectErr {
			t.Errorf("Expected error: %v, Got error: %v", test.expectErr, err != nil)
		}

		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("Expected error message: %v, Got error message: %v", test.expectedErr, err)
		}
	}
}

func TestCheckItemAndTimer(t *testing.T) {
	tests := []struct {
		itemToCook   string
		timer        int
		expectErr    bool
		expectedErr  error
	}{
		{"", 5, true, errors.New("Please specify the item to cook!")},
		{"rice", 2, true, errors.New("Time to cook is insufficient")},
		{"rice", 4, false, nil},
		{"pasta", 3, false, nil},
	}

	for _, test := range tests {
		err := CheckItemAndTimer(test.itemToCook, test.timer)

		if (err != nil) != test.expectErr {
			t.Errorf("Expected error: %v, Got error: %v", test.expectErr, err != nil)
		}

		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("Expected error message: %v, Got error message: %v", test.expectedErr, err)
		}
	}
}

func TestStartCooking(t *testing.T) {
	tests := []struct {
		itemToCook   string
		water        string
		timer        int
		expectErr    bool
		expectedErr  error
	}{
		{"rice", "2", 5, false, nil},
		{"pasta", "3", -1, true, errors.New("Timer must be greater than 0!")},
	}

	for _, test := range tests {
		err := StartCooking(test.itemToCook, test.water, test.timer)

		if (err != nil) != test.expectErr {
			t.Errorf("Expected error: %v, Got error: %v", test.expectErr, err != nil)
		}

		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("Expected error message: %v, Got error message: %v", test.expectedErr, err)
		}
	}
}
