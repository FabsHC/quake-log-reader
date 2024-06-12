# quake-log-reader
![quake-log-reader](https://img.shields.io/badge/quake--log--reader-gray?logo=go)
![technology Go 1.22.0](https://img.shields.io/badge/technology-go%201.22.0-blue.svg)

## Technology
The technology used is the Go language in version 1.22.0.

## Project Organization
```
.
├── cmd..................: Contains the main file to run the application.
├── docs.................: Contains the documentation from some examples.
├── internal.............: all core implementation its here.
│   ├── model............: All application structures (DTO).
│   ├── usecase..........: Application core validations.
│   └── util.............: General stuff.

```

## How to run
The application expects data in STDIN format and will return json in STDOUT format. So you can provide a JSON or a file containing several lines with each line having a JSON.
Examples below:

### Using Go Run
You can run the commands below to run the application
```shell
go run cmd/main.go
```
```shell
go run cmd/main.go < resources/logfile 
```

### Using Docker
Just run the docker commands below to create the docker image and run the container.
``` shell
docker run -it --rm -v $PWD:$PWD -w $PWD -v /var/run/docker.sock:/var/run/docker.sock golang:1.22.0 go run cmd/main.go
```
``` shell
docker run -i --rm -v $PWD:$PWD -w $PWD -v /var/run/docker.sock:/var/run/docker.sock golang:1.22.0 go run cmd/main.go < resources/logfile
```

## How to run the tests
Run the command below in the terminal to run the application tests.
### Using docker:
```shell
make -f Makefile test-docker
```
### Using go:
```shell
make -f Makefile test-go
```

