package movie

//This going to use Golang reflection for json convertion
type movie struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

//Implementing Stringer interface
func (movie *movie) String() string {
	return movie.Name + " " + movie.Director
}
