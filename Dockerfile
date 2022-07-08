FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go run main.go

EXPOSE 8080

CMD [ "go", "run" ]