package monitoring

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

const (
	dockerErrorMessage string = "An <b>unknown error</b> was occured during metrics gaining...\nMake sure you running <i>Docker</i> on your machine"
)

type DockerMetrics struct {
	Client client.Client
}

func (d *DockerMetrics) GetAllContainers(ctx context.Context, limit int) string {
	if limit > 20 || limit < 1 {
		return dockerErrorMessage
	}

	opts := container.ListOptions{Limit: limit, All: true}
	containers, err := d.Client.ContainerList(ctx, opts)
	if err != nil {
		return dockerErrorMessage
	}

	var message string

	for _, container := range containers {
		var ports string
		for _, port := range container.Ports {
			ports += fmt.Sprintf("[IP: %s, Type: %s] ", port.IP, port.Type)
		}

		var mounts string
		for _, mount := range container.Mounts {
			mounts += fmt.Sprintf("[%s <b>%s</b> to <b>%s</b>] ",
				mount.Type,
				mount.Source,
				mount.Destination,
			)
		}

		message += fmt.Sprintf(
			"<b>%s</b>: <i>%s</i>\n- ID: %s\n- Image: %s\n- Network mode: %s\n- Ports: %s\n- Mounts: %s\n- Created at: %s\n- Command to launch: <b>%s</b>\n\n",
			container.Names[0][1:len(container.Names[0])],
			container.Status,
			container.ID,
			container.Image,
			container.HostConfig.NetworkMode,
			ports,
			mounts,
			time.Unix(container.Created, 0).String(),
			container.Command,
		)
	}

	return message
}

func (d *DockerMetrics) GetImages(ctx context.Context) string {
	images, err := d.Client.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return dockerErrorMessage
	}

	var message string

	for _, image := range images {
		tag := image.RepoTags[0]
		if strings.Contains(tag, "<") {
			tag = "none"
		}
		message += fmt.Sprintf(
			"<b>%s</b>\n<i>%s</i>\nCreated at: %s\nSize (mb): %d\n\n",
			tag,
			image.ID,
			time.Unix(image.Created, 0).String(),
			image.Size/1024/1024,
		)
	}

	return message
}
