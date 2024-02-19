package services

import (
	"github.com/Victor90001/prod/internal/entity"
	"github.com/Victor90001/prod/internal/interfaces"
	"github.com/Victor90001/prod/internal/requests"
	"strconv"
)

type ListService struct {
	repo interfaces.ListRepository
}

func NewListService(repo interfaces.ListRepository) (*ListService, error) {
	return &ListService{
		repo: repo,
	}, nil
}

func (s *ListService) GetLists() ([]map[string]string, error) {
	lists, err := s.repo.GetLists()
	if err != nil {
		return nil, err
	}
	listsSlice := []map[string]string{}
	for _, item := range lists {
		tmpMap := map[string]string{
			"id":             strconv.Itoa(item.Id),
			"groupID":      strconv.Itoa(item.DealerId),
			"name":           item.Name,
			"price":          strconv.Itoa(item.Price),
			"code":         strconv.Itoa(item.Amount),
			"prodDate":     item.CreatedAt,
			"describe":           item.Info,
			"size":        item.Carrier,
			"country": item.ContactPerson,
			"addParam":           item.Note,
		}
		listsSlice = append(listsSlice, tmpMap)
	}
	return listsSlice, nil
}

func (s *ListService) InsertList(request requests.InsertListRequest) error {
	return s.repo.InsertList(entity.List(request))
}

func (s *ListService) UpdateList(request requests.UpdateListRequest) error {
	return s.repo.UpdateList(entity.List(request))
}

func (s *ListService) DeleteList(id int) error {
	return s.repo.DeleteList(id)
}
