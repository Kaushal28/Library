package entities

type Book struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Cost      float64 `json:"cost"`
	Year      int     `json:"year"`
	Publisher string  `json:"publisher"`
}
