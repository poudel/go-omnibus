package configlang

import (
	"fmt"
	"testing"
)

func TestParseConfig(t *testing.T) {
	c := `
name : str = "K Paudel"				// explicit type
website := "https://keshab.net"		// type inferred

// This is a dictionary
address = {
  	timezone = "Asia/Kathmandu",

	// this is a comment
	// or, is it?
	city = "Kathmandu",
}

// This is a list
phones = {
	9842002937,
	9483427344
}`
	fmt.Println(c)

}
