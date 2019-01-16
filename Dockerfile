# go build stage
FROM golang:1.11 as gobuilder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go version
RUN go mod download

# I copy the front end needlessly here... need to fix this
COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/stopandsearch/stopandsearch.go



# vue build stage
FROM node as vuebuilder
WORKDIR /app
COPY . .
RUN make frontend

# get ca-certificate to sign the police api
FROM alpine:3.6 as alpine

RUN apk add -U --no-cache ca-certificates


# final stage
FROM scratch
COPY --from=gobuilder /app/stopandsearch /app/
COPY --from=gobuilder /app/conf/config.json /app/conf/
COPY --from=vuebuilder /app/web/dist/ /app/
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/