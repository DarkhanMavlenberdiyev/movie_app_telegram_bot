package main

import (
	"github.com/urfave/cli"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"

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

	b.Handle("/start",endpoints.Start(b))
	b.Handle(&endpoints.PopularTvKey,endpoints.GetPopularTv(b))
	b.Handle(&endpoints.NextTV,endpoints.NextPopularTv(b))
	b.Handle(&endpoints.PrevTV,endpoints.PrevPopularTv(b))

	b.Start()
	return nil
}