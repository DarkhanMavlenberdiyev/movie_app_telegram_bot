package tv

import (
	"encoding/json"
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"net/http"

	"fmt"
)
const (
	API_KEY = "398aa9b4671e95dd4cc9f81eb2dcea03"
	IMG_URL = "https://image.tmdb.org/t/p/w500"
)

var (
	countTv    = 0
	page   = 1
	tv     = TVRate{}

	NextTV   = tb.InlineButton{Text: "Next", Unique: "t1"}
	PrevTV   = tb.InlineButton{Text: "Previous", Unique: "t2"}
	SaveTV   = tb.InlineButton{Text: "Save", Unique: "t3"}
	tvInline = [][]tb.InlineButton{[]tb.InlineButton{PrevTV, NextTV},
		[]tb.InlineButton{SaveTV}}
)

func NewEndpointsFactoryTv(tv TvDb) *endpointsFactory {
	return &endpointsFactory{tvs: tv}
}

type endpointsFactory struct {
	tvs  TvDb
}

func (ef *endpointsFactory) GetPopularTv(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {

		req, _ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/tv/popular?api_key=%v&language=en-US&page=%v", API_KEY, page))
		reqData, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(reqData, &tv)

		res := fmt.Sprintf("Place: #1\nTitle: %v\nPopularity: %v\nRelease date: %v", tv.Results[0].Name, tv.Results[0].Popularity, tv.Results[0].FirstAirDate)
		photo := &tb.Photo{File: tb.FromURL(IMG_URL + tv.Results[0].BackdropPath), Caption: res}

		b.Send(m.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: tvInline})

	}
}


func (ef *endpointsFactory) NextPopularTv(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		if countTv < 9 {
			countTv++
			res := fmt.Sprintf("Place: #%v\nTitle: %v\nPopularity: %v\nRelease date: %v", countTv+1, tv.Results[countTv].Name, tv.Results[countTv].Popularity, tv.Results[countTv].FirstAirDate)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL + tv.Results[countTv].BackdropPath), Caption: res}
			b.Delete(c.Message)
			b.Send(c.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: tvInline})

		}
	}
}
func (ef *endpointsFactory) PrevPopularTv(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		if countTv > 0 {
			countTv--
			res := fmt.Sprintf("Place: #%v\nTitle: %v\nPopularity: %v\nRelease date: %v", countTv+1, tv.Results[countTv].Name, tv.Results[countTv].Popularity, tv.Results[countTv].FirstAirDate)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL + tv.Results[countTv].BackdropPath), Caption: res}
			b.Delete(c.Message)
			b.Send(c.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: tvInline})
		}
	}
}