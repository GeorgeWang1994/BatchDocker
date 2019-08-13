package docker

import (
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"

	"github.com/docker/docker/go-connections/nat"
	"io"
	"strings"
	"time"
)

type CreateContainerOptions struct {
	Config        *container.Config
	HostConfig    *container.HostConfig
	NetworkConfig *network.NetworkingConfig
	Name          string
}

// generate create options for create container
func (d *Docker) NewCreateOptions(config map[string]string) (CreateContainerOptions, error) {
	options := &CreateContainerOptions{
		Config:        &container.Config{},
		HostConfig:    &container.HostConfig{},
	}

	if image := config["Image"]; image == "" {
		return *options, errors.New("config's image name is empty")
	}
	options.Config.Image = config["Image"]

	if name := config["Name"]; name == "" {
		return *options, errors.New("config's name is empty")
	}
	options.Name = config["Name"]

	if user := config["User"]; user == "" {
		return *options, errors.New("config's name is empty")
	}
	options.Config.User = config["User"]

	if cmd := config["Cmd"]; cmd != "" {
		cmds := strings.Split(cmd, ",")
		for _, c := range cmds {
			options.Config.Cmd = append(options.Config.Cmd, c)
		}
	}

	port := config["Port"]
	hostPort := config["HostPort"]
	ip := config["HostIP"]

	if ip == "" {
		ip = "0.0.0.0"
	}

	if port != "" && hostPort != "" {
		options.HostConfig.PortBindings = nat.PortMap {
			nat.Port(port + "/tcp"): {
				HostIP:   ip,
				HostPort: hostPort,
		}}
	}

	options.Config.AttachStderr = true
	options.Config.AttachStdout = true
	options.Config.AttachStdin = true

	return *options, nil
}

// create container
func (d *Docker) CreateContainer(opt CreateContainerOptions) (container.ContainerCreateCreatedBody, error) {
	return d.ContainerCreate(context.TODO(), opt.Config, opt.HostConfig, opt.NetworkConfig, opt.Name)
}

// container list
func (d *Docker) ListContainer(opts types.ContainerListOptions) ([]types.Container, error) {
	return d.ContainerList(context.TODO(), opts)
}

// get container info
func (d *Docker) InspectContainer(containerID string) (types.ContainerJSON, error) {
	res, _, err := d.ContainerInspectWithRaw(context.TODO(), containerID, false)
	return res, err
}

// start container
func (d *Docker) StartContainer(containerID string) error {
	return d.ContainerStart(context.TODO(), containerID, types.ContainerStartOptions{})
}

// stop container, default timeout is 1 minute
func (d *Docker) StopContainer(containerID string) error {
	var timeout = time.Minute
	return d.ContainerStop(context.TODO(), containerID, &timeout)
}

// remove container
func (d *Docker) RemoveContainer(containerID string) error {
	return d.ContainerRemove(context.TODO(), containerID, types.ContainerRemoveOptions{})
}

// copy something to container
func (d *Docker) CopyContainer(containerID, dstPath string, content io.Reader) error {
	return d.CopyToContainer(context.TODO(), containerID, dstPath, content, types.CopyToContainerOptions{})
}

// from container copy something
func (d *Docker) FromContainer(containerID, srcPath string) (io.ReadCloser, error) {
	fd, _, err := d.CopyFromContainer(context.TODO(), containerID, srcPath)
	return fd, err
}
