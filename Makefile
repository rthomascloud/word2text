all:
	go build -o bin/word2text

all-linux32:
	GOOS=linux GOARCH=386 go build -o bin/word2text-linux32

all-linux64:
	GOOS=linux GOARCH=amd64 go build -o bin/word2text-linux64

all-linux: all-linux64
