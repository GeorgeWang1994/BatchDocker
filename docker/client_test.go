package docker

import (
	"context"
	"github.com/docker/docker/client"
	"reflect"
	"testing"
)

const (
	endpoint   = "unix:///var/run/docker.sock"
	caPath     = "tls/ca.pem"
	certPath   = "tls/cert.pem"
	keyPath    = "tls/key.pem"
	version = "1.0.0"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig(endpoint, caPath, certPath, keyPath, version)
	if config.endpoint != endpoint {
		t.Fatal("endpoint error")
	}

	if config.caPath != caPath {
		t.Fatal("caPath error")
	}

	if config.certPath != certPath {
		t.Fatal("certPath error")
	}

	if config.keyPath != keyPath {
		t.Fatal("keyPath error")
	}

	if config.version != version {
		t.Fatal("version error")
	}
}

func TestNewDockerWithTls(t *testing.T) {
	config := NewConfig(endpoint, caPath, certPath, keyPath, version)
	Client1 := NewDocker(config)

	client2, err := client.NewClientWithOpts(client.WithTLSClientConfig(caPath, certPath, keyPath), client.WithHost(endpoint),
		client.WithVersion(version))
	if err != nil {
		t.Fatal(err)
	}
	Client2 := &Docker{client2}

	if reflect.ValueOf(Client1).Elem().FieldByName("endpoint").String() !=
		reflect.ValueOf(Client2).Elem().FieldByName("endpoint").String() {
		t.Fatal("new docker with tls type error")
	}
}

func TestNewDockerCommon(t *testing.T) {
	config := NewConfig(endpoint, "", "", "", "")
	Client := NewDocker(config)
	_, err := Client.Ping(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
}
