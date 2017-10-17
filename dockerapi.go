package main

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// outputRunningContainers outputs info to cli using defined methods
// currently only tabwriter is implemented
func outputRunningContainers(ctx context.Context, cli *client.Client) {
	ctxd, cancel := context.WithCancel(ctx)
	defer cancel()

	containers, err := cli.ContainerList(ctxd, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	tableOutput(containers)
}

// searchRunningContainers returns slice of the current running containers
func searchRunningContainers(ctx context.Context, cli *client.Client, term string) ([]types.Container, error) {
	nameFilters := filters.NewArgs()
	nameFilters.Add("name", term)
	containerListOptions := types.ContainerListOptions{
		Filters: nameFilters,
	}

	ctxd, cancel := context.WithCancel(ctx)
	defer cancel()

	containers, err := cli.ContainerList(ctxd, containerListOptions)
	if err != nil {
		return []types.Container{}, err
	}
	if len(containers) < 1 {
		return containers, errors.New("Cannot find Container")
	}

	return containers, nil
}
