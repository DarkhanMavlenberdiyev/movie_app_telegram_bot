package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
	tb "gopkg.in/tucnak/telebot.v2"

	"./endpoints"
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
	user := endpoints.PostgreConfig{
		User:     "postgres",
		Password: "daha2000",
		Port:     "5432", //5432
		Host:     "127.0.0.1",
	}
	db := endpoints.NewPostgreBot(user)
	endpoint := endpoints.NewEndpointsFactory(db)

	b.Handle("/start", endpoint.Start(b))
	//tv endpoints
	b.Handle(&endpoints.PopularTvKey, endpoint.GetPopularTv(b))
	b.Handle(&endpoints.NextTV, endpoint.NextPopularTv(b))
	b.Handle(&endpoints.PrevTV, endpoint.PrevPopularTv(b))

	//movies endpoints
	b.Handle(&endpoints.PopularMovieKey, endpoint.GetPopularMovies(b))
	b.Handle(&endpoints.NextMovie, endpoint.NextPopularMovie(b))
	b.Handle(&endpoints.PrevMovie, endpoint.PrevPopularMovie(b))
	b.Handle(&endpoints.SaveMovie, endpoint.SaveMovie(b))
	b.Handle(&endpoints.MyMoviesKey, endpoint.GetMyMovies(b))
	b.Handle(&endpoints.NextMyMovie, endpoint.NextMyMovie(b))
	b.Handle(&endpoints.PrevMyMovie, endpoint.PrevMyMovie(b))
	b.Handle(&endpoints.DeleteMyMovie, endpoint.DeleteMyMovie(b))
	b.Start()
	return nil
}
