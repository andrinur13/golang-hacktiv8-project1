package todos

type Todo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}
