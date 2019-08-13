package product

import (
	"batchdocker/docker"
	"errors"
	"os"
)

// check agent's path
func checkPath(path string) error {
	if err, _ := os.Stat(path); err != nil {
		return errors.New("文件路径不存在")
	}
	return nil
}

// move local agent's file to honeypot
func MoveAgentFile(client *docker.Docker, containerID, srcPath, dstPath string) error {
	if err := checkPath(srcPath); err != nil {
		return err
	}
	fd, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	err = client.CopyContainer(containerID, dstPath, fd)
	if err != nil {
		return err
	}
	return nil
}
