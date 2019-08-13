package docker

import (
	"github.com/docker/docker/api/types"
)

// batch get network's info
func (d *Docker) BatchInspectNetwork(networkIDs []string) []types.NetworkResource {
	var networks []types.NetworkResource
	for _, networkID := range networkIDs {
		network, err := d.InspectNetwork(networkID)
		if err != nil {
			networks = append(networks, network)
		}
	}
	return networks
}

// batch remove network
func (d *Docker) BatchRemoveNetwork(networkIDs []string) []error {
	var errList []error
	for _, networkID := range networkIDs {
		err := d.RemoveNetwork(networkID)
		if err != nil {
			errList = append(errList, err)
		}
	}
	return errList
}
