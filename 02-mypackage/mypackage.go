package mypackage

var bar = "This is not exported."
var Bar = "This variable IS exported."

func barfunc(x int) int {
	return x - 1
}

func Barfunc(x int) int {
	return x + 1
}
