package main

import (
	"io"
	"log"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/context"
)

func main() {
	client, err := docker.NewEnvClient()
	if err != nil {
		log.Fatalf("unable to connect to Docker: %v", err)
	}

	options := types.EventsOptions{}

	events, errs := client.Events(context.Background(), options)

	for {
		select {
		case err := <-errs:
			if err != nil && err != io.EOF {
				log.Println("err:", err)
			}

		case event := <-events:
			spew.Dump(event)
		}
	}
}
