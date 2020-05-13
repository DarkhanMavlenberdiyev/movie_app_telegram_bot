package endpoints



type MoviesDb interface {
	GetMovie(id int) (*Movie,error)
	CreateMovie(movie *Movie) (*Movie,error)
	DeleteMovie(id int) error

	GetUser(id int)	(*UserMovies,error)
	CreateUser(user *UserMovies) (*UserMovies,error)
	UpdateUser(id int,user *UserMovies) (*UserMovies,error)


}

type UserMovies struct {
	ID 				int 	`json:"id"`
	MoviesList			string	`json:"movies_list"`
}


type Movie struct {
	Title 				string 	`json:"title"`
	ID 					int 	`json:"id"`
	Overview 			string 	`json:"overview"`
	Budget				int 	`json:"budget"`
	Popularity			float64	`json:"popularity"`
	BackdropPath		string	`json:"backdrop_path"`
	Homepage			string 	`json:"homepage"`
	PosterPath			string	`json:"poster_path"`
	OriginalLanguage	string 	`json:"original_language"`
	ReleaseDate			string 	`json:"release_date"`
	Status				string 	`json:"status"`
	VoteCount			int 	`json:"vote_count"`
}

type TV struct {
	Name 				string	`json:"name"`
	ID 					int 	`json:"id"`
	Overview			string	`json:"overview"`
	Homepage			string 	`json:"homepage"`
	FirstAirDate		string	`json:"first_air_date"`
	NumberOfEpisodes	int		`json:"number_of_episodes"`
	NumberOfSeasons		int 	`json:"number_of_seasons"`
	OriginalLanguage	string	`json:"original_language"`
	Popularity			float64	`json:"popularity"`
	PosterPath			string	`json:"poster_path"`
	BackdropPath		string	`json:"backdrop_path"`
	Status				string	`json:"status"`
	VoteCount			int		`json:"vote_count"`

}

type MovieRate struct {
	Results []struct{
		Title 			string 	`json:"title"`
		ID 				int 	`json:"id"`
		Overview 		string 	`json:"overview"`
		Popularity 		float64	`json:"popularity"`
		VoteCount 		int 	`json:"vote_count"`
		PosterPath 		string 	`json:"poster_path"`
		BackdropPath 	string 	`json:"backdrop_path"`
		ReleaseDate 	string 	`json:"release_date"`
	} `json:"results"`
}


type TVRate struct {
	Results []struct{
		Name 			string 	`json:"name"`
		ID 				int 	`json:"id"`
		Overview 		string 	`json:"overview"`
		Popularity 		float64 `json:"popularity"`
		VoteCount 		int 	`json:"vote_count"`
		PosterPath 		string 	`json:"poster_path"`
		BackdropPath 	string 	`json:"backdrop_path"`
		FirstAirDate 	string 	`json:"first_air_date"`
	} `json:"results"`
}



type ListGenres struct {
	Genres []struct {
		ID		int 	`json:"id"`
		Name	string 	`json:"name"`
	}	`json:"genres"`
}

type Person struct {
	ID 			int 	`json:"id"`
	Name		string 	`json:"name"`
	Birthday	string	`json:"birthday"`
	Deathday	string 	`json:"deathday"`
	Gender		int		`json:"gender"`
	Biography	string 	`json:"biography"`
	Popularity	float64	`json:"popularity"`
	ProfilePath	string	`json:"profile_path"`
	Homepage	string 	`json:"homepage"`
}

type TrendingPersons struct {
	Results []struct{
		ID 			int 	`json:"id"`
		Gender		int 	`json:"gender"`
		Name		string	`json:"name"`
		ProfilePath	string	`json:"profile_path"`
		PosterPath	string	`json:"poster_path"`
		Popularity	float64	`json:"popularity"`
		KnownFor	[]Movie	`json:"known_for"`
	} `json:"results"`
}

