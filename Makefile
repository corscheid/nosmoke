all: format compile

format: nosmoke.go
	gofmt -w nosmoke.go

compile:
	go build nosmoke.go

install: all
	[ -e ~/.nosmoke.json ] || cp ./.nosmoke.json ~/
	sudo cp ./nosmoke /usr/local/bin/

uninstall:
	sudo rm -f /usr/local/bin/nosmoke
	rm ~/.nosmoke.json

clean:
	rm -f ./nosmoke
