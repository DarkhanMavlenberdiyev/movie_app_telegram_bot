package movie

type MoviesDb interface {
	GetMovie(id int) (*Movie, error)
	CreateMovie(movie *Movie) (*Movie, error)
	DeleteMovie(id int) error


	GetMyMovie(user_id int) ([]*Movie, error)
	DeleteMyMovie(id int, user_id int) error
}




type Movie struct {
	Title            string  `json:"title"`
	ID               int     `json:"id"`
	Overview         string  `json:"overview"`
	Budget           int     `json:"budget"`
	Popularity       float64 `json:"popularity"`
	BackdropPath     string  `json:"backdrop_path"`
	Homepage         string  `json:"homepage"`
	PosterPath       string  `json:"poster_path"`
	OriginalLanguage string  `json:"original_language"`
	ReleaseDate      string  `json:"release_date"`
	Status           string  `json:"status"`
	VoteCount        int     `json:"vote_count"`
	UserID           int     `json:"user_id"` // fk
}


type MovieRate struct {
	Results []struct {
		Title        string  `json:"title"`
		ID           int     `json:"id"`
		Overview     string  `json:"overview"`
		Popularity   float64 `json:"popularity"`
		VoteCount    int     `json:"vote_count"`
		PosterPath   string  `json:"poster_path"`
		BackdropPath string  `json:"backdrop_path"`
		ReleaseDate  string  `json:"release_date"`
	} `json:"results"`
}



type ListGenres struct {
	Genres []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
}

type Person struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Birthday    string  `json:"birthday"`
	Deathay    string  `json:"deathday"`
	Gender      int     `json:"gender"`
	Biography   string  `json:"biography"`
	Popularity  float64 `json:"popularity"`
	ProfilePath string  `json:"profile_path"`
	Homepage    string  `json:"homepage"`
}

type TrendingPersons struct {
	Results []struct {
		ID          int     `json:"id"`
		Gender      int     `json:"gender"`
		Name        string  `json:"name"`
		ProfilePath string  `json:"profile_path"`
		PosterPath  string  `json:"poster_path"`
		Popularity  float64 `json:"popularity"`
		KnownFor    []Movie `json:"known_for"`
	} `json:"results"`
}
