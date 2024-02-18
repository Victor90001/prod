package interfaces

import "github.com/Victor90001/prod/internal/entity"

type NetworkRepository interface {
	GetNetworks() ([]entity.Network, error)
	UpdateNetwork(Network entity.Network) error
	InsertNetwork(Network entity.Network) error
	DeleteNetwork(id int) error
}
