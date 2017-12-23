package docker

import (
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/network"
	"golang.org/x/net/context"
)

func (d *Docker) ListNetworks() ([]types.NetworkResource, error) {
	return d.client.NetworkList(context.Background(), types.NetworkListOptions{})
}

func (d *Docker) NetworkExists(name string) (bool, error) {
	networks, err := d.ListNetworks()
	if err != nil {
		return false, err
	}
	for _, network := range networks {
		if network.Name == name {
			return true, nil
		}
	}
	return false, nil
}

func (d *Docker) NetworkID(name string) (string, error) {
	networks, err := d.ListNetworks()
	if err != nil {
		return "", err
	}
	for _, network := range networks {
		if network.Name == name {
			return network.ID, nil
		}
	}
	return "", nil
}

func (d *Docker) CreateNetwork(name string, overlay bool, attachable bool, subnets []string) (string, error) {
	ipamCfg := []network.IPAMConfig{}
	for _, s := range subnets {
		ipamCfg = append(ipamCfg, network.IPAMConfig{Subnet: s, AuxAddress: map[string]string{}})
	}
	spec := types.NetworkCreate{
		CheckDuplicate: true,
		Attachable:     attachable,
	}
	if ipamCfg != nil {
		spec.IPAM = &network.IPAM{
			Driver: "default",
			Config: ipamCfg,
		}
	}
	if overlay {
		spec.Driver = "overlay"
	}
	network, err := d.client.NetworkCreate(context.Background(), name, spec)
	if err != nil {
		return "", err
	}
	return network.ID, nil
}

func (d *Docker) RemoveNetwork(name string) error {
	return d.client.NetworkRemove(context.Background(), name)
}
