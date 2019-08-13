package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"io"
)

// batch create container
func (d *Docker) BatchCreateContainer(opts []CreateContainerOptions) []container.ContainerCreateCreatedBody {
	var containers []container.ContainerCreateCreatedBody
	for _, opt := range opts {
		info, err := d.CreateContainer(opt)
		if err != nil {
			continue
		}
		containers = append(containers, info)
	}
	return containers
}

// batch get container's info
func (d *Docker) BatchInspectContainer(containerIDs []string) []types.ContainerJSON {
	var containers []types.ContainerJSON
	for _, containerID := range containerIDs {
		info, err := d.InspectContainer(containerID)
		if err != nil {
			continue
		}
		containers = append(containers, info)
	}
	return containers
}

// batch start container
func (d *Docker) BatchStartContainer(containerIDs []string) ([]string, []error) {
	var errList []error
	var successIDs []string
	for _, containerID := range containerIDs {
		err := d.StartContainer(containerID)
		if err != nil {
			errList = append(errList, err)
		} else {
			successIDs = append(successIDs, containerID)
		}
	}
	return successIDs, errList
}

// batch stop container
func (d *Docker) BatchStopContainer(containerIDs []string) ([]string, []error) {
	var errList []error
	var successIDs []string
	for _, containerID := range containerIDs {
		err := d.StopContainer(containerID)
		if err != nil {
			errList = append(errList, err)
		} else {
			successIDs = append(successIDs, containerID)
		}
	}
	return successIDs, errList
}

// batch remove container
func (d *Docker) BatchRemoveContainer(containerIDs []string) ([]string, []error) {
	var errList []error
	var successIDs []string
	for _, containerID := range containerIDs {
		err := d.RemoveContainer(containerID)
		if err != nil {
			errList = append(errList, err)
		} else {
			successIDs = append(successIDs, containerID)
		}
	}
	return successIDs, errList
}

// batch copy to container
func (d *Docker) BatchCopyContainer(containerIDs []string, dstPath string, content io.Reader) ([]string, []error) {
	var errList []error
	var successIDs []string
	for _, containerID := range containerIDs {
		err := d.CopyContainer(containerID, dstPath, content)
		if err != nil {
			errList = append(errList, err)
		} else {
			successIDs = append(successIDs, containerID)
		}
	}
	return successIDs, errList
}

// batch from container copy
func (d *Docker) BatchFromContainer(containerIDs []string, srcPath string) []io.ReadCloser {
	var fds []io.ReadCloser
	for _, containerID := range containerIDs {
		fd, err := d.FromContainer(containerID, srcPath)
		if err != nil {
			continue
		}
		fds = append(fds, fd)
	}
	return fds
}
