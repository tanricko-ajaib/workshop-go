package main

import (
	"fmt"

	. "github.com/tanricko-ajaib/workshop/practical/titlecase"
)

func main() {
	tests := []struct {
		in   string
		want string
	}{
		{in: "john doe", want: "John Doe"},
		{in: "alice", want: "Alice"},
		{in: "ångström unit", want: "Ångström Unit"},
		{in: "élise dupont", want: "Élise Dupont"},
	}

	for _, tc := range tests {
		got := TitleCase(tc.in)
		fmt.Printf("in:   %q\n", tc.in)
		fmt.Printf("got:  %q\n", got)
		fmt.Printf("want: %q\n\n", tc.want)
	}
}
