package main

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	if !strings.EqualFold("FizzBuzz", FizzBuzz(15)) {
		t.Errorf("correct")
	}

}
