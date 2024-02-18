package requests

type InsertNetworkRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateNetworkRequest struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteNetworkRequest struct {
	Id int `json:"id" binding:"required"`
}
