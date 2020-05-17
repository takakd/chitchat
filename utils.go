package main

import "fmt"

type Configuration struct {
	Address string
	Static string
}

var config Configuration

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a)
}

// version
func version() string {
	return "0.1"
}
