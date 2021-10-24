package todos

type Todo struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"Cuci Baju"`
	Complete bool   `json:"complete" example:"false"`
}
