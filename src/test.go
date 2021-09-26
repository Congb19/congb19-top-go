package test

import (
	"fmt"
)

var ga int  = 123
var gb = 123

func main() {
	fmt.Println("start");
	var a string = "123"
	var b int = 2
	var c bool
	var d float64 = 0.00
	e := true
	fmt.Println(ga, gb, a, b, c, d, e);
	fmt.Println(&ga, &gb, &a, &b, &c, &d, &e);
	
}