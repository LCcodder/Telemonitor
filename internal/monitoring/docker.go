package monitoring

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/k0kubun/pp/v3"
)

const (
	dockerErrorMessage string = "An <b>unknown error</b> was occured during metrics gaining...\nMake sure you running <i>Docker</i> on your machine"
)

type DockerMetrics struct {
	Client client.Client
}

func (d *DockerMetrics) GetRunningContainers(ctx context.Context, limit int) string {
	if limit > 20 || limit < 1 {
		return dockerErrorMessage
	}
	opts := container.ListOptions{Limit: limit}
	containers, err := d.Client.ContainerList(ctx, opts)
	if err != nil {
		return dockerErrorMessage
	}
	pp.Print(containers)
	return ""
}

func (d *DockerMetrics) GetAllContainers(ctx context.Context, limit int) string {
	if limit > 20 || limit < 1 {
		return dockerErrorMessage
	}
	opts := container.ListOptions{Limit: limit, Filters: filters.Args{}}
	containers, err := d.Client.ContainerList(ctx, opts)
	if err != nil {
		return dockerErrorMessage
	}

	var message string

	for _, container := range containers {
		message += fmt.Sprintf(
			"<b>[%s]</b>: <i>%s</i>\n",
			container.ID,
			container.Names[0],
		)
	}

	pp.Print(containers)

	return message
}

func (d *DockerMetrics) GetImages(ctx context.Context) string {
	images, err := d.Client.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return dockerErrorMessage
	}

	pp.Print(images)

	return ""
}
