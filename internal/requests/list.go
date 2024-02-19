package requests

//type GetListsRequest struct {
//	Id int `json:"id" binding:"required"`
//}

type InsertListRequest struct {
	Id            int    `json:"id"`
	DealerId      int    `json:"groupID" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	Amount        int    `json:"code" binding:"required"`
	CreatedAt     string `json:"prodDate" binding:"required"`
	Info          string `json:"describe" binding:"required"`
	Carrier       string `json:"size" binding:"required"`
	ContactPerson string `json:"country" binding:"required"`
	Note          string `json:"addParam" binding:"required"`
}

type UpdateListRequest struct {
	Id            int    `json:"id" binding:"required"`
	DealerId      int    `json:"groupID"`
	Name          string `json:"name" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	Amount        int    `json:"code" binding:"required"`
	CreatedAt     string `json:"prodDate" binding:"required"`
	Info          string `json:"describe" binding:"required"`
	Carrier       string `json:"size" binding:"required"`
	ContactPerson string `json:"country" binding:"required"`
	Note          string `json:"addParam" binding:"required"`
}

type DeleteListRequest struct {
	Id int `json:"id" binding:"required"`
}
