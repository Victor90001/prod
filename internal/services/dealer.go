package services

import (
	"github.com/Victor90001/prod/internal/entity"
	"github.com/Victor90001/prod/internal/interfaces"
	"strconv"
)

type DealerService struct {
	repo interfaces.DealerRepository
}

func NewDealerService(repo interfaces.DealerRepository) (*DealerService, error) {
	return &DealerService{
		repo: repo,
	}, nil
}

func (s *DealerService) GetDealers() ([]map[string]string, error) {
	dealers, err := s.repo.GetDealers()
	if err != nil {
		return nil, err
	}
	dealersSlice := []map[string]string{}
	for _, item := range dealers {
		tmpMap := map[string]string{
			"id":         strconv.Itoa(item.Id),
			"sectionID": strconv.Itoa(item.NetworkId),
			"name":       item.Name,
		}
		dealersSlice = append(dealersSlice, tmpMap)
	}
	return dealersSlice, nil
}

func (s *DealerService) InsertDealer(networkId int, name string) error {
	return s.repo.InsertDealer(entity.Dealer{Id: 0, NetworkId: networkId, Name: name})
}

func (s *DealerService) UpdateDealer(id int, name string) error {
	return s.repo.UpdateDealer(entity.Dealer{Id: id, NetworkId: 0, Name: name})
}

func (s *DealerService) DeleteDealer(id int) error {
	return s.repo.DeleteDealer(id)
}
