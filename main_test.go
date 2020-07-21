package main

import "testing"


type testpair struct {
	value string
	isInt bool
}
type sizepair struct {
	value int
	isLarge bool
}

var intTests = []testpair{
	{ "55", true },
	{ "hello", false },
	{ "12.1", false },
	{ "-4", false},
	{ "-4.1", false},
}

var bigTests = []sizepair{
	{ 41, false },
	{ 43, true },
	{ -42, false },
}

func TestIsInteger(t *testing.T) {
	for _, pair := range intTests {
		v := IsInteger(pair.value)
		if v!= pair.isInt{
			t.Error(
				"For", pair.value,
				"expected", pair.isInt,
				"got", v,
			  )
		}
	}
}

func TestTooBig(t *testing.T) {
	for _, pair := range bigTests {
		v := TooBig(pair.value)
		if v!= pair.isLarge{
			t.Error(
				"For", pair.value,
				"expected", pair.isLarge,
				"got", v,
			  )
		}
	}
}




