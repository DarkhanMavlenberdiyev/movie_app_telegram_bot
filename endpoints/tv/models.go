package tv

type TV struct {
	Name             string  `json:"name"`
	ID               int     `json:"id"`
	Overview         string  `json:"overview"`
	Homepage         string  `json:"homepage"`
	FirstAirDate     string  `json:"first_air_date"`
	NumberOfEpisodes int     `json:"number_of_episodes"`
	NumberOfSeasons  int     `json:"number_of_seasons"`
	OriginalLanguage string  `json:"original_language"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	BackdropPath     string  `json:"backdrop_path"`
	Status           string  `json:"status"`
	VoteCount        int     `json:"vote_count"`
	UserID			 int 	 `json:"user_id"`//fk
}


type TvDb interface {
	GetTv(id int) (*TV,error)
	CreateTv(tv *TV) (*TV,error)
	DeleteTv (id int) error

	GetMyTv(user_id int) ([]*TV,error)
	DeleteMyTv(id int,user_id int) error
}

type TVRate struct {
	Results []struct {
		Name         string  `json:"name"`
		ID           int     `json:"id"`
		Overview     string  `json:"overview"`
		Popularity   float64 `json:"popularity"`
		VoteCount    int     `json:"vote_count"`
		PosterPath   string  `json:"poster_path"`
		BackdropPath string  `json:"backdrop_path"`
		FirstAirDate string  `json:"first_air_date"`
	} `json:"results"`
}