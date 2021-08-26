package movie

// swagger:parameters getMovieById
type GetByIdParam struct {

	// id of a movie
	//
	// In: path
	Id string `json:"id"`
}

// swagger:parameters saveMovie
type SaveRequest struct {

	// in: body
	// required: true
	Mov *Movie `json:"movie"`
}
