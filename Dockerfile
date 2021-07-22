FROM golang:1.16

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY src/*.go ./
COPY src/Makefile ./

RUN go build -o /api_stone

EXPOSE 16453

CMD [ "/api_stone" ]