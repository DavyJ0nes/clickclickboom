package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"

	"github.com/docker/docker/api/types"
)

// checkError is abstraction for checking error output from command
func checkError(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(2)
	}
}

// getRandomContainer
func getRandomContainer(containerList []types.Container) types.Container {
	// set pseudo random generator
	rand.Seed(time.Now().Unix())
	return containerList[rand.Intn(len(containerList))]
}

//tableOutput
func tableOutput(containers []types.Container) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tIMAGE\tSTATUS")
	for _, container := range containers {
		infoLine := fmt.Sprintf("%s\t%s\t%s\t%s", container.ID[:10], container.Names[0], container.Image, container.Status)
		fmt.Fprintln(w, infoLine)
	}
	w.Flush()
}
