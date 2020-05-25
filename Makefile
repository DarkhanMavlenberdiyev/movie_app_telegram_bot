

build:
	go build main.go
run:
	./main start
depends:
	go get "github.com/urfave/cli"
	go get "gopkg.in/tucnak/telebot.v2"
	go get "github.com/go-pg/pg"