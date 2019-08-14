package batchdocker

import (
	"batchdocker/docker"
	"batchdocker/product"
	"context"
	"flag"
	"fmt"
)

var (
	endpoint = flag.String("endpoint", "unix:///var/run/docker.sock", "endpoint")
	caPath = flag.String("capath", "", "capath")
	certPath = flag.String("certpath", "", "certpath")
	keyPath = flag.String("keypath", "", "keypath")
	version = flag.String("version", "1.0.0", "version")
)

func main() {
	config := docker.NewConfig(*endpoint, *caPath, *certPath, *keyPath, *version)
	client := docker.NewDocker(config)

	if _, err := client.Ping(context.TODO()); err != nil {
		panic(err)
	}

	imageNames := []string{product.WindowsImageName, product.UbuntuImageName}
	agentPaths := []string{product.UbuntuAgentPath, product.WindowsAgentPath}
	for idx, imageName := range imageNames {
		err := product.PullSystem(client, imageName)
		if err != nil {
			fmt.Printf("pull system has some error...\n")
			continue
		}

		successIDs, err := product.CreateSystem(client, "ubuntu-test", imageName, 5, agentPaths[idx])
		if err != nil {
			fmt.Printf("create container has some error...\n")
		}

		fmt.Printf("success container ID:[%v]\n", successIDs)
	}
}
