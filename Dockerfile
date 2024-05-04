FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ENV TOKEN 5230264384:AAFJYWJ0EW4aM4GcPiBlU0Pp_OJK1-vUO7Y

RUN go build -o cmd/telemonitor
CMD [ "./cmd/telemonitor" ]