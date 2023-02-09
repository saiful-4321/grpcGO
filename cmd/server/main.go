package main

import "log"

func Run() error {
	// responsible for initializing and starting grpc server
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
