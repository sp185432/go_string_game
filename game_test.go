package main

import (
	"log"
	"testing"
)

var testcases = []struct {
	guessword, userword string
	output              int
}{
	{guessword: "SHARANYA", userword: "sharanya", output: 0},
	{guessword: "pineapple", userword: "apple", output: 0},
	{guessword: "shar123", userword: "shar123", output: 0},
	{guessword: "interns", userword: "interns", output: 0},
	{guessword: "123456", userword: "123456", output: 0},
	{guessword: "shar12", userword: "password", output: 0},
	{guessword: "shar12", userword: "SHAR12", output: 0},
}

func TestCompareString(t *testing.T) {

	for _, tc := range testcases {

		actual := CompareString(tc.guessword, tc.userword)

		expected := tc.output

		if actual != expected {

			log.Printf("Strings are not equal ")
			log.Printf("Guessword: %s User entered word: %s Actual: %d Expected:%d", tc.guessword, tc.userword, actual, expected)

		} else {

			t.Log("Both the strings are equal")
			t.Logf("Guessword: %s User entered word: %s Actual:%d Expected:%d", tc.guessword, tc.userword, actual, expected)

		}

	}

}

var values = []struct {
	input, output string
}{

	{input: "pineapple", output: "aeeilnppp"},
	{input: "pineapple", output: "pppnlieea"},
	{input: "aeeilnppp", output: "aeeilnppp"},
	{input: "pineapple", output: "apple"},
	{input: "pineapple12", output: "aeeilnppp12"},
	{input: "pineapple12", output: "12aeeilnppp"},
}

func TestReArrange(t *testing.T) {

	for _, val := range values {

		actual := ReArrange(val.input)

		expected := val.output

		if actual != expected {

			log.Printf("The entered and rearranged strings are not equal ")
			log.Printf("Actual word: %s Rearranged word: %s Actual:%s Expected:%s", val.input, val.output, actual, expected)

		} else {

			t.Log("Both the strings are equal")
			t.Logf("Actual word: %s Rearranged word: %s Actual:%s Expected:%s", val.input, val.output, actual, expected)

		}

	}
}
