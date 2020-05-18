package endpoints

import tb "gopkg.in/tucnak/telebot.v2"

var(
	PopularTvKey    = tb.ReplyButton{Text: "10 popular TV"}
	PopularMovieKey = tb.ReplyButton{Text: "10 popular Movies"}
	MyMoviesKey     = tb.ReplyButton{Text: "My movies"}
	MyTvKey         = tb.ReplyButton{Text: "My TV"}

	startReply = [][]tb.ReplyButton{[]tb.ReplyButton{PopularMovieKey, PopularTvKey},
		[]tb.ReplyButton{MyTvKey, MyMoviesKey},
	}
)

func Start(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {
		b.Send(m.Sender, "Welcome to Movie/TV bot", &tb.ReplyMarkup{ReplyKeyboard: startReply})
	}
}