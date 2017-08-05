package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCountLines(t *testing.T) {
	var tests = map[string]int{
		"a.txt": 3,
		"b.txt": 6,
		"c.txt": 7,
		"d.txt": 5,
		"e.txt": 1,
		"f.txt": 0,
		"g.txt": 4,
	}
	for k, v := range tests {
		t.Run(fmt.Sprintf("_filename:%s", k), func(t *testing.T) {
			data, err := ioutil.ReadFile("testdata/" + k)
			if err != nil {
				t.Fatal(err)
			}
			if r := CountLines(data); r != v {
				t.Errorf("expected %d, got %d", v, r)
			}
		})
	}
}
