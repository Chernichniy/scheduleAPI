FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/senya/schedule-bot 
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/schedule-bot /go/src/gitlab.com/senya/schedule-bot 


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/schedule-bot /usr/local/bin/schedule-bot
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["schedule-bot"]
