package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"io"
)

// get logs of container
func (d *Docker) Logs(containerID string, opt types.ContainerLogsOptions) (io.Reader, error) {
	return d.ContainerLogs(context.Background(), containerID, opt)
}
