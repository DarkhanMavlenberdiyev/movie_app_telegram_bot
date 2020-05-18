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
	myTv = []*TV{}
	myTvCount=0

	NextTV   = tb.InlineButton{Text: "Next", Unique: "t1"}
	PrevTV   = tb.InlineButton{Text: "Previous", Unique: "t2"}
	SaveTV   = tb.InlineButton{Text: "Save", Unique: "t3"}
	tvInline = [][]tb.InlineButton{[]tb.InlineButton{PrevTV, NextTV},
		[]tb.InlineButton{SaveTV}}

	//My tv
	NextMyTv    = tb.InlineButton{Text: "Next", Unique: "mt1"}
	PrevMyTv    = tb.InlineButton{Text: "Previous", Unique: "mt2"}
	DeleteMyTv  = tb.InlineButton{Text: "Delete", Unique: "mt3"}
	LinkHomepage   = tb.InlineButton{Text: "Link", Unique: "mt4"}
	myTvInline = [][]tb.InlineButton{[]tb.InlineButton{PrevMyTv, NextMyTv}, []tb.InlineButton{DeleteMyTv, LinkHomepage}}
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

func (ef *endpointsFactory) SaveTv(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		req, _ := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/tv/%v?api_key=%v&language=en-US", tv.Results[countTv].ID, API_KEY))
		reqData, _ := ioutil.ReadAll(req.Body)
		tv := &TV{}
		tv.UserID = c.Sender.ID
		json.Unmarshal(reqData, &tv)

		_, er := ef.tvs.CreateTv(tv)

		if er != nil {
			b.Respond(c, &tb.CallbackResponse{Text: "There is error:(", ShowAlert: true})
			fmt.Print(er.Error())
			return
		}

		b.Respond(c, &tb.CallbackResponse{Text: "TV saved!"})
	}
}

func (ef *endpointsFactory) GetMyTv(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {
		res, err := ef.tvs.GetMyTv(m.Sender.ID)
		myTv = res
		if err != nil || len(myTv) == 0 {
			b.Send(m.Sender, "There is error or your list is empty")
			return
		}
		response := fmt.Sprintf("Title: %v\nFirst air date: %v\nOverview: %v", myTv[myTvCount].Name, myTv[myTvCount].FirstAirDate, myTv[myTvCount].Overview)
		photo := &tb.Photo{File: tb.FromURL(IMG_URL + myTv[myTvCount].BackdropPath), Caption: response}
		LinkHomepage = tb.InlineButton{Text: "Link", Unique: "mt4", URL: myTv[myTvCount].Homepage}
		myTvInline[1][1] = LinkHomepage
		fmt.Print(len(myTv))
		b.Send(m.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: myTvInline})

	}
}

func (ef *endpointsFactory) NextMyTv(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		if myTvCount < len(myTv)-1 {
			myTvCount++
			response := fmt.Sprintf("Title: %v\nFirst air date: %v\nOverview: %v", myTv[myTvCount].Name, myTv[myTvCount].FirstAirDate, myTv[myTvCount].Overview)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL + myTv[myTvCount].BackdropPath), Caption: response}
			LinkHomepage = tb.InlineButton{Text: "Link", Unique: "mt4", URL: myTv[myTvCount].Homepage}
			myTvInline[1][1] = LinkHomepage
			b.Delete(c.Message)
			b.Send(c.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: myTvInline})

		}
	}
}

func (ef *endpointsFactory) PrevMyTv(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		if myTvCount > 0 {
			myTvCount--
			response := fmt.Sprintf("Title: %v\nFirst air date: %v\nOverview: %v", myTv[myTvCount].Name, myTv[myTvCount].FirstAirDate, myTv[myTvCount].Overview)
			photo := &tb.Photo{File: tb.FromURL(IMG_URL + myTv[myTvCount].BackdropPath), Caption: response}
			LinkHomepage = tb.InlineButton{Text: "Link", Unique: "mt4", URL: myTv[myTvCount].Homepage}
			myTvInline[1][1] = LinkHomepage
			b.Delete(c.Message)
			b.Send(c.Sender, photo, &tb.ReplyMarkup{InlineKeyboard: myTvInline})

		}
	}
}
func (ef *endpointsFactory) DeleteMyTv(b *tb.Bot) func(c *tb.Callback) {
	return func(c *tb.Callback) {
		err := ef.tvs.DeleteMyTv(myTv[myTvCount].ID, c.Sender.ID)
		if err != nil {
			b.Respond(c, &tb.CallbackResponse{Text: "Can't delete TV. Try again!", ShowAlert: true})
			return
		}
		myTv[myTvCount].Name = fmt.Sprintf("Title: %v (Deleted from list)", myTv[myTvCount].Name)
		b.Respond(c, &tb.CallbackResponse{Text: "TV deleted successfully!"})

	}
}
