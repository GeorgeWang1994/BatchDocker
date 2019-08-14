package product

import (
	"batchdocker/docker"
	"errors"
	"strconv"
)

const (
	UbuntuImageName = "ubuntu:18.04"
	UbuntuAgentPath = "/agent/"

	WindowsImageName = "mcr.microsoft.com/windows"
	WindowsAgentPath = "c:/Users/agent"
)

var LocalAgentPath string


func PullSystem(client *docker.Docker, imageName string) error {
	if err := client.PullImage(imageName); err != nil {
		return err
	}
	return nil
}

// create system
func CreateSystem(client *docker.Docker, name, imageName string, count int, agentPath string) ([]string, error) {
	opts := make([]docker.CreateContainerOptions, 0)
	for i := 0; i < count; i++ {
		opt, err := client.NewCreateOptions(map[string]string{
			"Image": imageName,
			"Name": name + "-" + strconv.Itoa(i),
		})
		if err != nil {
			continue
		}
		opts = append(opts, opt)
	}

	if len(opts) == 0 {
		return nil, errors.New("create configs is empty")
	}

	containers := client.BatchCreateContainer(opts)
	if len(containers) == 0 {
		return nil, errors.New("all container create failed")
	}

	containerIDs := make([]string, 0)
	for _, container := range containers {
		containerIDs = append(containerIDs, container.ID)
	}

	successIDs, _ := client.BatchStartContainer(containerIDs)
	if len(successIDs) == 0 {
		return nil, errors.New("all container start failed")
	}

	containerIDs = make([]string, 0)
	for _, containerID := range successIDs {
		err := MoveAgentFile(client, containerID, LocalAgentPath, agentPath)
		if err != nil {
			continue
		}
		containerIDs = append(containerIDs, containerID)
	}

	return containerIDs, nil
}
