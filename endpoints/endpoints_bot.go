package endpoints

import (
	api	"../../api/endpoints"
	"encoding/json"
	"fmt"


	//"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"net/http"
	"strconv"
)
var (
	count = 17
	page = 1
	tv = api.TVRate{}

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
		[]tb.InlineButton{SaveTV},
		}

)
const (
	API_KEY = "398aa9b4671e95dd4cc9f81eb2dcea03"
	IMG_URL = "https://image.tmdb.org/t/p/w500"
)

func Start(b *tb.Bot) func (m *tb.Message){
	return func(m *tb.Message) {
		b.Send(m.Sender,"Welcome to Movie/TV bot",&tb.ReplyMarkup{ReplyKeyboard:startReply})
	}
}

func GetTVGenres(b *tb.Bot) func(m *tb.Message){
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

func GetPopularTv(b *tb.Bot) func(m *tb.Message){
	return func(m *tb.Message) {
		page = 1
		count=0
		req,_ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/tv/popular?api_key=%v&language=en-US&page=%v",API_KEY,page))
		reqData,_ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqData,&tv)


		res := fmt.Sprintf("Title: %v\nPopularity: %v\nOverview: %v\nRelease date: %v",tv.Results[0].Name,tv.Results[0].Popularity,tv.Results[0].Overview,tv.Results[0].FirstAirDate)
		photo := &tb.Photo{File: tb.FromURL(IMG_URL+tv.Results[0].BackdropPath),Caption:res}

		b.Send(m.Sender,photo,&tb.ReplyMarkup{InlineKeyboard:tvInline})

	}
}

func NextPopularTv(b *tb.Bot) func (c *tb.Callback){
	return func(c *tb.Callback) {
		count++
		if count<10 {
			res := fmt.Sprintf("Place: #%v\nTitle: %v\nPopularity: %v\nRelease date: %v",count+1,tv.Results[count].Name,tv.Results[count].Popularity,tv.Results[count].FirstAirDate)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL+tv.Results[count].BackdropPath),Caption:res}
			b.EditMedia(c.Message,photo)
			b.EditReplyMarkup(c.Message,&tb.ReplyMarkup{InlineKeyboard:tvInline})
		}
	}
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
