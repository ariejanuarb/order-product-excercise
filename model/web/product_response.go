package web

type ProductResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId int    `json:"category_id"`
}
