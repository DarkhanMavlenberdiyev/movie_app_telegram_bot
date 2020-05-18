package movie

import (
	"encoding/json"
	"fmt"
	//"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"net/http"
)

var (
	countMovie = 0

	page   = 1

	movies = MovieRate{}



	NextMovie    = tb.InlineButton{Text: "Next", Unique: "tt1"}
	PrevMovie    = tb.InlineButton{Text: "Previous", Unique: "tt2"}
	SaveMovie    = tb.InlineButton{Text: "Save", Unique: "tt3"}
	moviesInline = [][]tb.InlineButton{[]tb.InlineButton{PrevMovie, NextMovie},
		[]tb.InlineButton{SaveMovie},
	}

	//My movies
	NextMyMovie    = tb.InlineButton{Text: "Next", Unique: "m1"}
	PrevMyMovie    = tb.InlineButton{Text: "Previous", Unique: "m2"}
	DeleteMyMovie  = tb.InlineButton{Text: "Delete", Unique: "m3"}
	LinkHomepage   = tb.InlineButton{Text: "Link", Unique: "m4"}
	myMoviesInline = [][]tb.InlineButton{[]tb.InlineButton{PrevMyMovie, NextMyMovie}, []tb.InlineButton{DeleteMyMovie, LinkHomepage}}
	myMovies       = []*Movie{}
	myMoviesCount  = 0
)

const (
	API_KEY = "398aa9b4671e95dd4cc9f81eb2dcea03"
	IMG_URL = "https://image.tmdb.org/t/p/w500"
)

func NewEndpointsFactory(movie MoviesDb) *endpointsFactory {
	return &endpointsFactory{movies: movie}
}


type endpointsFactory struct {
	movies MoviesDb

}


func (ef *endpointsFactory) GetPopularMovies(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {

		req, _ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/movie/popular?api_key=%v&language=en-US&page=%v", API_KEY, page))
		reqData, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqData, &movies)

		res := fmt.Sprintf("Place: #1\nTitle: %v\nPopularity: %v\nRelease date: %v", movies.Results[0].Title, movies.Results[0].Popularity, movies.Results[0].ReleaseDate)
		photo := &tb.Photo{File: tb.FromURL(IMG_URL + movies.Results[0].BackdropPath), Caption: res}

		b.Send(m.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: moviesInline})

	}
}

func (ef *endpointsFactory) NextPopularMovie(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		if countMovie < 9 {
			countMovie++
			res := fmt.Sprintf("Place: #%v\nTitle: %v\nPopularity: %v\nRelease date: %v", countMovie+1, movies.Results[countMovie].Title, movies.Results[countMovie].Popularity, movies.Results[countMovie].ReleaseDate)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL + movies.Results[countMovie].BackdropPath), Caption: res}
			b.Delete(c.Message)
			b.Send(c.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: moviesInline})

		}
	}
}
func (ef *endpointsFactory) PrevPopularMovie(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		if countMovie > 0 {
			countMovie--
			res := fmt.Sprintf("Place: #%v\nTitle: %v\nPopularity: %v\nRelease date: %v", countMovie+1, movies.Results[countMovie].Title, movies.Results[countMovie].Popularity, movies.Results[countMovie].ReleaseDate)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL + movies.Results[countMovie].BackdropPath), Caption: res}
			b.Delete(c.Message)
			b.Send(c.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: moviesInline})
		}
	}
}

func (ef *endpointsFactory) SaveMovie(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		req, _ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/movie/%v?api_key=%v&language=en-US", movies.Results[countMovie].ID, API_KEY))
		reqData, _ := ioutil.ReadAll(req.Body)
		movie := Movie{}
		movie.UserID = c.Sender.ID
		json.Unmarshal(reqData, &movie)

		_, er := ef.movies.CreateMovie(&movie)

		if er != nil {
			b.Respond(c, &tb.CallbackResponse{Text: "There is error:(", ShowAlert: true})
			return
		}

		b.Respond(c, &tb.CallbackResponse{Text: "Movie saved!"})
	}
}

func (ef *endpointsFactory) GetMyMovies(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {
		res, err := ef.movies.GetMyMovie(m.Sender.ID)
		myMovies = res
		if err != nil || len(myMovies) == 0 {
			b.Send(m.Sender, "There is error or your list is empty")
			return
		}
		response := fmt.Sprintf("Title: %v\nRelease-Date: %v\nLink: %v\nOverview: %v", myMovies[myMoviesCount].Title, myMovies[myMoviesCount].ReleaseDate, myMovies[myMoviesCount].Homepage, myMovies[myMoviesCount].Overview)
		photo := &tb.Photo{File: tb.FromURL(IMG_URL + myMovies[myMoviesCount].BackdropPath), Caption: response}
		LinkHomepage = tb.InlineButton{Text: "Link", Unique: "m4", URL: myMovies[myMoviesCount].Homepage}
		myMoviesInline[1][1] = LinkHomepage
		b.Send(m.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: myMoviesInline})

	}
}

func (ef *endpointsFactory) NextMyMovie(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		if myMoviesCount < len(myMovies)-1 {
			myMoviesCount++
			response := fmt.Sprintf("Title: %v\nRelease-Date: %v\nLink: %v\nOverview: %v", myMovies[myMoviesCount].Title, myMovies[myMoviesCount].ReleaseDate, myMovies[myMoviesCount].Homepage, myMovies[myMoviesCount].Overview)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL + myMovies[myMoviesCount].BackdropPath), Caption: response}
			LinkHomepage = tb.InlineButton{Text: "Link", Unique: "m4", URL: myMovies[myMoviesCount].Homepage}
			myMoviesInline[1][1] = LinkHomepage
			b.Delete(c.Message)
			b.Send(c.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: myMoviesInline})

		}
	}
}

func (ef *endpointsFactory) PrevMyMovie(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		if myMoviesCount > 0 {
			myMoviesCount--
			response := fmt.Sprintf("Title: %v\nRelease-Date: %v\nLink: %v\nOverview: %v", myMovies[myMoviesCount].Title, myMovies[myMoviesCount].ReleaseDate, myMovies[myMoviesCount].Homepage, myMovies[myMoviesCount].Overview)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL + myMovies[myMoviesCount].BackdropPath), Caption: response}
			LinkHomepage = tb.InlineButton{Text: "Link", Unique: "m4", URL: myMovies[myMoviesCount].Homepage}
			myMoviesInline[1][1] = LinkHomepage
			b.Delete(c.Message)
			b.Send(c.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: myMoviesInline})

		}
	}
}
func (ef *endpointsFactory) DeleteMyMovie(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		err := ef.movies.DeleteMyMovie(myMovies[myMoviesCount].ID, c.Sender.ID)
		if err != nil {
			b.Respond(c, &tb.CallbackResponse{Text: "Can't delete movie. Try again!", ShowAlert: true})
			return
		}
		myMovies[myMoviesCount].Title = fmt.Sprintf("Title: %v (Deleted from list)", myMovies[myMoviesCount].Title)
		b.Respond(c, &tb.CallbackResponse{Text: "Movie deleted successfully!"})

	}
}


