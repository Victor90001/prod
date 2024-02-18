package entity

type Dealer struct {
	Id        int    `json:"id"`
	NetworkId int    `json:"sectionID"`
	Name      string `json:"name"`
}
