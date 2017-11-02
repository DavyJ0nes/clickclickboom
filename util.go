package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"text/tabwriter"
	"time"

	"github.com/docker/docker/api/types/swarm"

	"github.com/docker/docker/api/types"
)

type output interface {
	tableOutput()
}

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
func tableOutput(service string, conts []types.Container, servs []swarm.Service) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	if service == "containers" {
		fmt.Fprintln(w, "ID\tNAME\tIMAGE\tSTATUS")
		for _, container := range conts {
			infoLine := fmt.Sprintf("%s\t%s\t%s\t%s", container.ID[:10], container.Names[0], container.Image, container.Status)
			fmt.Fprintln(w, infoLine)
		}
	} else if service == "services" {
		fmt.Fprintln(w, "ID\tNAME\tIMAGE")
		for _, service := range servs {
			infoLine := fmt.Sprintf("%s\t%s\t%v", service.ID[:10], service.Spec.Name, service.Spec.TaskTemplate.ContainerSpec.Image)
			fmt.Fprintln(w, infoLine)
		}
	} else {
		fmt.Fprintln(w, "unknown data")
	}

	w.Flush()
}

func convertData(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
