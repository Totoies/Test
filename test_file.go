package test

import (
	"fmt"
)

var Massage = "This is for only Testing Perpose"

type Routes map[string]func() string

func PrintRoutes(_routes Routes) {
	for r, function := range _routes {
		v := function()
		fmt.Println(r, v)
	}
}
