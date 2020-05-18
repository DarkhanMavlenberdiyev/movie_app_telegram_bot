package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
	tb "gopkg.in/tucnak/telebot.v2"
	"./endpoints"
	"./endpoints/movie"
	"./endpoints/tv"

)

func main() {

	app := cli.NewApp()
	app.Commands = cli.Commands{
		&cli.Command{
			Name:   "start",
			Usage:  "start the bot",
			Action: StartBot,
		},
	}
	app.Run(os.Args)

}

func StartBot(d *cli.Context) error {
	b, err := tb.NewBot(tb.Settings{
		Token:  "1065088890:AAHsp6mSFeTC0mf3sZ5WEi8ODL4ZfxHi1cg",
		URL:    "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	userMovie := movie.PostgreConfig{
		User:     "postgres",
		Password: "daha2000",
		Port:     "5432", //5432
		Host:     "127.0.0.1",
	}
	userTv := tv.PostgreConfig{
		User:     "postgres",
		Password: "daha2000",
		Port:     "5432", //5432
		Host:     "127.0.0.1",
	}
	dbMovie := movie.PostgreMovies(userMovie)
	dbTv := tv.PostgreTv(userTv)
	endpointMovie := movie.NewEndpointsFactory(dbMovie)
	endpointTv := tv.NewEndpointsFactoryTv(dbTv)

	b.Handle("/start", endpoints.Start(b))
	// Tv endpoints
	b.Handle(&endpoints.PopularTvKey, endpointTv.GetPopularTv(b))
	b.Handle(&tv.NextTV, endpointTv.NextPopularTv(b))
	b.Handle(&tv.PrevTV, endpointTv.PrevPopularTv(b))
	b.Handle(&tv.SaveTV,endpointTv.SaveTv(b))
	// My TV endpoints
	b.Handle(&endpoints.MyTvKey,endpointTv.GetMyTv(b))
	b.Handle(&tv.NextMyTv,endpointTv.NextMyTv(b))
	b.Handle(&tv.PrevMyTv,endpointTv.PrevMyTv(b))
	b.Handle(&tv.DeleteMyTv,endpointTv.DeleteMyTv(b))

	// Movies endpoints
	b.Handle(&endpoints.PopularMovieKey, endpointMovie.GetPopularMovies(b))
	b.Handle(&movie.NextMovie, endpointMovie.NextPopularMovie(b))
	b.Handle(&movie.PrevMovie, endpointMovie.PrevPopularMovie(b))
	b.Handle(&movie.SaveMovie, endpointMovie.SaveMovie(b))
	// My movies endpoints
	b.Handle(&endpoints.MyMoviesKey, endpointMovie.GetMyMovies(b))
	b.Handle(&movie.NextMyMovie, endpointMovie.NextMyMovie(b))
	b.Handle(&movie.PrevMyMovie, endpointMovie.PrevMyMovie(b))
	b.Handle(&movie.DeleteMyMovie, endpointMovie.DeleteMyMovie(b))
	b.Start()
	return nil
}
