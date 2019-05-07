package main

import "fmt"

func main() {
	resp, err := GetDefaultResponse()

	if err != nil {
		fmt.Println("Could not retrieve information")
		return
	}

	fmt.Println(resp)
}