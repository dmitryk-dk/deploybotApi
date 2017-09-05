package structs

type Meta struct {
	Next    int    `json:"next"`
	NextUri string `json:"next_uri"`
	Total   int    `json:"total"`
}
