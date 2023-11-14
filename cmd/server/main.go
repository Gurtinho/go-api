package main

import "github.com/gurtinho/go/api/configs"

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}