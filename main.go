package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/docker/docker/client"
)

var (
	version = "No Version Provided"
	date    = ""
	gitHash = ""
)

func main() {
	versionString := fmt.Sprintf("%s (%s) %s", version, gitHash, date)
	searchTerm := flag.String("name", "", "the name of the container you want to kill")
	list := flag.Bool("list", false, "Do you want to just list running containers?")
	version := flag.Bool("version", false, "Version Info")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  Version: %s:\n", versionString)
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()

	cli, err := client.NewEnvClient()
	checkError(err)

	// decision logic based on user flags
	if *list {
		outputRunningContainers(context.Background(), cli)
		os.Exit(0)
	} else if *version {
		fmt.Fprintf(os.Stdout, "  Version: %s:\n", versionString)
	} else {
		flag.Usage()
	}

	containerList, err := searchRunningContainers(context.Background(), cli, *searchTerm)
	checkError(err)

	container := getRandomContainer(containerList)
	err = cli.ContainerKill(context.Background(), container.ID, "KILL")
	checkError(err)
}
