FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/senya/scheduleAPI 
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/scheduleAPI /go/src/gitlab.com/senya/scheduleAPI


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/scheduleAPI /usr/local/bin/scheduleAPI
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["scheduleAPI"]
