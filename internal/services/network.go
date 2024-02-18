package services

import (
	"github.com/Victor90001/prod/internal/entity"
	"github.com/Victor90001/prod/internal/interfaces"
	"strconv"
)

type NetworkService struct {
	repo interfaces.NetworkRepository
}

func NewNetworkService(repo interfaces.NetworkRepository) (*NetworkService, error) {
	return &NetworkService{
		repo: repo,
	}, nil
}

func (s *NetworkService) GetNetworks() ([]map[string]string, error) {
	networks, err := s.repo.GetNetworks()
	if err != nil {
		return nil, err
	}
	networksSlice := []map[string]string{}
	for _, item := range networks {
		tmpMap := map[string]string{
			"id":   strconv.Itoa(item.Id),
			"name": item.Name,
		}
		networksSlice = append(networksSlice, tmpMap)
	}
	return networksSlice, nil
}

func (s *NetworkService) InsertNetwork(name string) error {
	return s.repo.InsertNetwork(entity.Network{Id: 0, Name: name})
}

func (s *NetworkService) UpdateNetwork(id int, name string) error {
	return s.repo.UpdateNetwork(entity.Network{Id: id, Name: name})
}

func (s *NetworkService) DeleteNetwork(id int) error {
	return s.repo.DeleteNetwork(id)
}
