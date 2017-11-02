package main

import (
	"context"
	"errors"
	"fmt"

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
	tableOutput("containers", containers, nil)
}

func outputServices(ctx context.Context, cli *client.Client) {
	ctxd, cancel := context.WithCancel(ctx)
	defer cancel()

	services, err := cli.ServiceList(ctxd, types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}
	tableOutput("services", nil, services)
}

// searchRunningContainers returns slice of the current running containers
func searchRunningContainers(ctx context.Context, cli *client.Client, id string) ([]types.Container, error) {
	nameFilters := filters.NewArgs()
	nameFilters.Add("Id", id)
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

func getContainersForService(ctx context.Context, cli *client.Client, term string) ([]types.Container, error) {
	nameFilters := filters.NewArgs()
	nameFilters.Add("name", term)
	taskListOptions := types.TaskListOptions{
		Filters: nameFilters,
	}

	ctxd, cancel := context.WithCancel(ctx)
	defer cancel()

	tasks, err := cli.TaskList(ctxd, taskListOptions)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		fmt.Println(task.Status.ContainerStatus.ContainerID)
	}

	return []types.Container{}, nil
}
