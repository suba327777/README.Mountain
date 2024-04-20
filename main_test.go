package main

import (
	"flag"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	userTestCases := []struct {
		caseName      string
		userName      string
		errRaisedFlag bool
	}{
		{"case1", "niwaniwa", false},
		{"case2", "niwaniwa2", true},
	}

	for _, tt := range userTestCases {
		t.Run(tt.caseName, func(t *testing.T) {
			flag.CommandLine.Set("userName", tt.userName)

			err := validateArgs()
			if errRaised(err) != tt.errRaisedFlag {
				t.Errorf("Expected error raised to be %v, got %v, error: %#v", tt.errRaisedFlag, !tt.errRaisedFlag, err)
			}
		})
	}
}

func errRaised(err error) bool {
	return err != nil
}
