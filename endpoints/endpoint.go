package endpoints

import (
	api "../../api/endpoints"
	"encoding/json"
	"fmt"
	"strings"

	//"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"net/http"
	"strconv"
)
var (
	countTv = 0
	countMovie = 0
	page    = 1
	tv      = api.TVRate{}
	movies	= api.MovieRate{}

	PopularTvKey = tb.ReplyButton{Text:"10 popular TV",}
	PopularMovieKey = tb.ReplyButton{Text:"10 popular Movies",}
	FindKey = tb.ReplyButton{Text:"Find",}

	startReply = [][]tb.ReplyButton{[]tb.ReplyButton{PopularMovieKey,PopularTvKey},
	[]tb.ReplyButton{FindKey},
	}



	NextTV = tb.InlineButton{Text:"Next",Unique:"t1"}
	PrevTV = tb.InlineButton{Text:"Previous",Unique:"t2"}
	SaveTV = tb.InlineButton{Text:"Save",Unique:"t3"}
	tvInline = [][]tb.InlineButton{[]tb.InlineButton{PrevTV,NextTV},
		[]tb.InlineButton{SaveTV},}

	NextMovie = tb.InlineButton{Text:"Next",Unique:"t1"}
	PrevMovie = tb.InlineButton{Text:"Previous",Unique:"t2"}
	SaveMovie = tb.InlineButton{Text:"Save",Unique:"t3"}
	moviesInline = [][]tb.InlineButton{[]tb.InlineButton{PrevMovie,NextMovie},
		[]tb.InlineButton{SaveMovie},
		}

)
const (
	API_KEY = "398aa9b4671e95dd4cc9f81eb2dcea03"
	IMG_URL = "https://image.tmdb.org/t/p/w500"
)
func NewEndpointsFactory(movie MoviesDb) *endpointsFactory {
	return &endpointsFactory{movies: movie}
}

type endpointsFactory struct {
	movies 		MoviesDb
}



func (ef *endpointsFactory) Start(b *tb.Bot) func (m *tb.Message){
	return func(m *tb.Message) {
		user := UserMovies{
			ID:     m.Sender.ID,
			MoviesList: "",
		}
		ef.movies.CreateUser(&user)
		b.Send(m.Sender,"Welcome to Movie/TV bot",&tb.ReplyMarkup{ReplyKeyboard:startReply})
	}
}

func (ef *endpointsFactory) GetTVGenres(b *tb.Bot) func(m *tb.Message){
	return func(m *tb.Message) {
		genres := ListGenres{}
		req,_ := http.Get("https://api.themoviedb.org/3/genre/tv/list?api_key=%v&language=en-US")
		reqData,_ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqData,&genres)
		res:=""
		for i:=0;i<len(genres.Genres);i++  {
			res+=strconv.Itoa(i+1)+") "+genres.Genres[i].Name+"\n"
		}
		b.Send(m.Sender,res)
	}
}

func (ef *endpointsFactory) GetPopularTv(b *tb.Bot) func(m *tb.Message){
	return func(m *tb.Message) {

		countTv =0
		req,_ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/tv/popular?api_key=%v&language=en-US&page=%v",API_KEY,page))
		reqData,_ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqData,&tv)

		res := fmt.Sprintf("Place: #1\nTitle: %v\nPopularity: %v\nRelease date: %v",tv.Results[0].Name,tv.Results[0].Popularity,tv.Results[0].FirstAirDate)
		photo := &tb.Photo{File: tb.FromURL(IMG_URL+tv.Results[0].BackdropPath),Caption:res}

		b.Send(m.Sender,photo,&tb.ReplyMarkup{InlineKeyboard:tvInline})

	}
}


func (ef *endpointsFactory) NextPopularTv(b *tb.Bot) func (c *tb.Callback){
	return func(c *tb.Callback) {
		if countTv <9 {
			countTv++
			res := fmt.Sprintf("Place: #%v\nTitle: %v\nPopularity: %v\nRelease date: %v", countTv+1,tv.Results[countTv].Name,tv.Results[countTv].Popularity,tv.Results[countTv].FirstAirDate)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL+tv.Results[countTv].BackdropPath),Caption:res}
			b.Send(c.Sender,photo,&tb.ReplyMarkup{InlineKeyboard:tvInline})

		}
	}
}
func (ef *endpointsFactory) PrevPopularTv(b *tb.Bot) func (c *tb.Callback){
	return func(c *tb.Callback) {
		if countTv >0 {
			countTv--
			res := fmt.Sprintf("Place: #%v\nTitle: %v\nPopularity: %v\nRelease date: %v", countTv+1,tv.Results[countTv].Name,tv.Results[countTv].Popularity,tv.Results[countTv].FirstAirDate)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL+tv.Results[countTv].BackdropPath),Caption:res}
			b.Send(c.Sender,photo,&tb.ReplyMarkup{InlineKeyboard:tvInline})
		}
	}
}


func (ef *endpointsFactory) GetPopularMovies(b *tb.Bot) func(m *tb.Message){
	return func(m *tb.Message) {

		req,_ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/movie/popular?api_key=%v&language=en-US&page=%v",API_KEY,page))
		reqData,_ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqData,&movies)

		res := fmt.Sprintf("Place: #1\nTitle: %v\nPopularity: %v\nRelease date: %v",movies.Results[0].Title,movies.Results[0].Popularity,movies.Results[0].ReleaseDate)
		photo := &tb.Photo{File: tb.FromURL(IMG_URL+movies.Results[0].BackdropPath),Caption:res}

		b.Send(m.Sender,photo,&tb.ReplyMarkup{InlineKeyboard:moviesInline})

	}
}


func (ef *endpointsFactory) NextPopularMovie(b *tb.Bot) func (c *tb.Callback){
	return func(c *tb.Callback) {
		if countMovie <9 {
			countMovie++
			res := fmt.Sprintf("Place: #%v\nTitle: %v\nPopularity: %v\nRelease date: %v", countMovie+1,movies.Results[countMovie].Title,movies.Results[countMovie].Popularity,movies.Results[countMovie].ReleaseDate)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL+movies.Results[countMovie].BackdropPath),Caption:res}
			b.Send(c.Sender,photo,&tb.ReplyMarkup{InlineKeyboard:moviesInline})

		}
	}
}
func (ef *endpointsFactory) PrevPopularMovie(b *tb.Bot) func (c *tb.Callback){
	return func(c *tb.Callback) {
		if countMovie >0 {
			countMovie--
			res := fmt.Sprintf("Place: #%v\nTitle: %v\nPopularity: %v\nRelease date: %v", countTv+1,movies.Results[countMovie].Title,movies.Results[countMovie].Popularity,movies.Results[countMovie].ReleaseDate)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL+movies.Results[countMovie].BackdropPath),Caption:res}
			b.Send(c.Sender,photo,&tb.ReplyMarkup{InlineKeyboard:moviesInline})
		}
	}
}

func (ef *endpointsFactory) SaveMovie(b *tb.Bot) func(c *tb.Callback){
	return func(c *tb.Callback) {
		req, _ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/movie/%v?api_key=%v&language=en-US",movies.Results[countMovie].ID,API_KEY))
		reqData,_ := ioutil.ReadAll(req.Body)
		movie := Movie{}
		json.Unmarshal(reqData,&movie)

		_,er := ef.movies.CreateMovie(&movie)

		if er != nil {
			b.Respond(c, &tb.CallbackResponse{Text: er.Error(),ShowAlert: true})
			return
		}
		list,_ := ef.movies.GetUser(c.Sender.ID)

		fmt.Print(c.Sender.ID)
		updList := ""
		if len(list.MoviesList)==0{
			updList=strconv.Itoa(movies.Results[countMovie].ID)

		}else{
			lis := strings.Split(list.MoviesList," ")
			if !contains(lis,strconv.Itoa(movies.Results[countMovie].ID)){
				updList=list.MoviesList+" "+strconv.Itoa(movies.Results[countMovie].ID)
			}else{
				updList = list.MoviesList
			}
		}
		updUser := &UserMovies{
			ID:     list.ID,
			MoviesList: updList,
		}
		ef.movies.UpdateUser(list.ID,updUser)
		b.Respond(c,&tb.CallbackResponse{Text: "Movie saved!"})
	}
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}



/*

func GetListPopularMovies() gin.HandlerFunc{
	return func(c *gin.Context) {
		page := c.Param("page")
		req, _ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/movie/popular?api_key=%v&language=en-US&page=%v",API_KEY,page))
		reqData, _ := ioutil.ReadAll(req.Body)
		movies := MovieRate{}
		json.Unmarshal(reqData,&movies)
		c.JSON(http.StatusOK, movies)
	}
}

func GetListPopularTV() gin.HandlerFunc{
	return func(c *gin.Context) {
		page := c.Param("page")
		req, _ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/tv/popular?api_key=%v&language=en-US&page=%v",API_KEY,page))
		reqData, _ := ioutil.ReadAll(req.Body)
		tv := TVRate{}
		json.Unmarshal(reqData,&tv)
		c.JSON(http.StatusOK, tv)
	}
}

func GetMovie() gin.HandlerFunc{
	return func(c *gin.Context) {
		id := c.Param("id")
		req, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/movie/%v?api_key=%v&language=en-US",id,API_KEY))
		if err!=nil{
			c.JSON(http.StatusBadRequest,nil)
			return
		}
		reqData,_ := ioutil.ReadAll(req.Body)
		movie := Movie{}
		json.Unmarshal(reqData,&movie)
		c.JSON(http.StatusOK,movie)
	}
}

func GetTV() gin.HandlerFunc{
	return func(c *gin.Context) {
		id := c.Param("id")
		req, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/tv/%v?api_key=%v&language=en-US",id,API_KEY))
		if err!=nil{
			c.JSON(http.StatusBadRequest,nil)
			return
		}
		reqData,_ := ioutil.ReadAll(req.Body)
		tv := TV{}
		json.Unmarshal(reqData,&tv)
		c.JSON(http.StatusOK,tv)
	}
}

func GetMovieGenres() gin.HandlerFunc{
	return func(c *gin.Context) {
		req,_ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/genre/movie/list?api_key=%v&language=en-US",API_KEY))
		reqData,_ := ioutil.ReadAll(req.Body)
		genres := ListGenres{}
		json.Unmarshal(reqData,&genres)
		c.JSON(http.StatusOK,genres)

	}
}


func GetPersonTrendingList() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("op")
		req,_ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/trending/person/%v?api_key=%v",param,API_KEY))
		reqData,_ := ioutil.ReadAll(req.Body)
		persons := TrendingPersons{}
		json.Unmarshal(reqData,&persons)
		c.JSON(http.StatusOK,persons)
	}
}
func GetPerson() gin.HandlerFunc{
	return func(c *gin.Context) {
		id := c.Param("id")
		req,_ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/person/%v?api_key=%v&language=en-US",id,API_KEY))
		reqData,_ := ioutil.ReadAll(req.Body)
		person := Person{}
		json.Unmarshal(reqData,&person)
		c.JSON(http.StatusOK,person)

	}
}

*/
