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
	instanceList := flag.Bool("list-containers", false, "Do you want to list running containers?")
	serviceList := flag.Bool("list-services", false, "Do you want to list running services?")
	version := flag.Bool("version", false, "Version Info")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Version: %s\n", versionString)
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()

	cli, err := client.NewEnvClient()
	checkError(err)

	// decision logic based on user flags
	if *instanceList {
		outputRunningContainers(context.Background(), cli)
		os.Exit(0)
	} else if *serviceList {
		outputServices(context.Background(), cli)
		os.Exit(0)
	} else if *version {
		fmt.Fprintf(os.Stdout, "  Version: %s:\n", versionString)
		os.Exit(0)
	} else if *searchTerm == "" {
		flag.Usage()
	} else {
		_, err = getContainersForService(context.Background(), cli, *searchTerm)
		// containerList, err := searchRunningContainers(context.Background(), cli, *searchTerm)
		checkError(err)

		// container := getRandomContainer(containerList)
		// err = cli.ContainerKill(context.Background(), container.ID, "KILL")
		// checkError(err)
	}
}
