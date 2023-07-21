package main

import (
	"fmt"
	"search_engine/app/linkgpraph/graph"
	"time"

	"github.com/google/uuid"
)

func main() {
	link := graph.Link{
		ID:          uuid.New(),
		URL:         "https://google.com",
		RetrievedAt: time.Now(),
	}

	fmt.Printf("%+v\n", link)
}
