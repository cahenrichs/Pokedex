package pokeapi


type LocationAreaList struct {
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

type LocationAreaResult struct {
	Name string `json:"name"`
}









type Location struct {

}