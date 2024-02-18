package requests

//type GetDealersRequest struct {
//	Id int `json:"id" binding:"required"`
//}

type InsertDealerRequest struct {
	NetworkId int    `json:"sectionID" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

type UpdateDealerRequest struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteDealerRequest struct {
	Id int `json:"id" binding:"required"`
}
