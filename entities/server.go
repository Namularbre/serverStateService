package entities

type Server struct {
	Id   int    `json:"id"`
	Ip   string `json:"ip"`
	Name string `json:"name"`
}
