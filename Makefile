# this is for gvm users using golang version 1.6.2
# make sure to run "make setup" before running this!
SERVICE := ghost

setup:
	rm -f ./$(SERVICE) && rm -rf vendor # preliminary clean up
	glide install
	make build

build:
	go build -ldflags "-X main.version=`date -u +%Y-%m-%d.%H:%M:%S`" -o $(SERVICE)

run:
	GOPRO_ENV=development ./$(SERVICE)

test:
	GOPRO_ENV=test go test -v `glide novendor`

install_golint:
	go get github.com/golang/lint/golint
	go install github.com/golang/lint/golint

run_golint:
	golint

