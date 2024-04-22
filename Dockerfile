FROM golang:1.22.2

WORKDIR /app
COPY go.* ./
RUN go mod download

RUN go build -o /app/cmd/telemonitor

ENV TOKEN 5230264384:AAFJYWJ0EW4aM4GcPiBlU0Pp_OJK1-vUO7Y

ENTRYPOINT [ "/app/cmd/telemonitor" ]