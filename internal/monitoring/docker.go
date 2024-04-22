package monitoring

import (
	"github.com/docker/docker/client"
)

type DockerMetrics struct {
	client client.Client
}

func (d *DockerMetrics) GetRunningContainers() string {
	return ""
}
func (d *DockerMetrics) GetAllContainers() string {
	return ""
}
func (d *DockerMetrics) GetImages() string {
	return ""
}
