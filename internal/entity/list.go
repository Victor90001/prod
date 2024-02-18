package entity

type List struct {
	Id            int    `json:"id"`
	DealerId      int    `json:"groupID"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Amount        int    `json:"code"`
	CreatedAt     string `json:"prodDate"`
	Info          string `json:"desc"`
	Carrier       string `json:"size"`
	ContactPerson string `json:"country"`
	Note          string `json:"addParam"`
}
