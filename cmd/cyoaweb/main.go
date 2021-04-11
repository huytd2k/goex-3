package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	port := flag.Int("port", 3000, "the port to start CYOA app on")
	flag.Parse()
	fmt.Printf("Using the story in %s \n", *filename)
	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)

	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)

	fmt.Printf("Starting server on port: %d", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
