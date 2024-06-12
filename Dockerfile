FROM golang:1.22.0
WORKDIR /app
COPY . .
RUN go test -v ./...
WORKDIR /app/cmd
RUN go build -o quake-log-reader .
ENTRYPOINT [ "/app/cmd/quake-log-reader"]