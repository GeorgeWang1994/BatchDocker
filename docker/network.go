package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

// network list
func (d *Docker) ListNetwork() ([]types.NetworkResource, error) {
	return d.NetworkList(context.TODO(), types.NetworkListOptions{})
}

// get info of network
func (d *Docker) InspectNetwork(networkID string) (types.NetworkResource, error) {
	res, _, err := d.NetworkInspectWithRaw(context.TODO(), networkID, types.NetworkInspectOptions{})
	return res, err
}

// remove network
func (d *Docker) RemoveNetwork(networkID string) error {
	return d.NetworkRemove(context.TODO(), networkID)
}
