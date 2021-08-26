package movie

// swagger:model Movie
type Movie struct {
	//id
	Id       int    `json:"id"`
	//name
	Name     string `json:"name"`
	//year
	Year     int    `json:"year"`
	//director
	Director string `json:"director"`
}

//Implementing Stringer interface
func (movie *Movie) String() string {
	return movie.Name + " " + movie.Director
}

