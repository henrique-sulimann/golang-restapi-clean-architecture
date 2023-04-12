FROM golang:alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN cd cmd && go build -o sulimannapp
EXPOSE 3000
CMD ["./cmd/sulimannapp"]