package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ironarachne/climategen"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	climate := climategen.Generate()

	fmt.Println(climate.Name)

	for _, resource := range climate.Resources {
		fmt.Println("Resource: " + resource.Name)
	}

	for _, need := range climate.Needs {
		fmt.Println("Need: " + need.Name)
	}
}
