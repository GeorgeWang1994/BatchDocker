package docker

import (
	"github.com/docker/docker/client"
	"os"
)

// client for docker
type Docker struct {
	*client.Client
}

type Config struct {
	endpoint   string
	caPath     string
	certPath   string
	keyPath    string
	version    string
}

const Host = "DOCKER_HOST"

// create config for docker client
func NewConfig(endPoint, caPath, certPath, keyPath, version string) *Config {
	return &Config{
		endpoint: endPoint,
		caPath  : caPath,
		certPath: certPath,
		keyPath:  keyPath,
		version:  version,
	}
}

// create docker client
func NewDocker(config *Config) *Docker {
	var cli *client.Client
	var err error
	if os.Getenv(Host) != "" {
		cli, err = client.NewClientWithOpts(client.FromEnv, client.WithVersion(config.version))
		if err != nil {
			panic(err)
		}
	}

	if config.certPath != "" && config.keyPath != "" && config.caPath != "" {
		cli, err = client.NewClientWithOpts(client.WithTLSClientConfig(config.caPath, config.certPath, config.keyPath),
			client.WithHost(config.endpoint), client.WithVersion(config.version))
		if err != nil {
			panic(err)
		}
	}

	cli, err = client.NewClientWithOpts(client.WithHost(config.endpoint), client.WithVersion(config.version))
	if err != nil {
		panic(err)
	}
	return &Docker{cli}
}