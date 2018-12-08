# build stage
# multi-stage build to reduce the size of final image.
FROM golang:1.11.2 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch
COPY --from=builder /app/fetch /app/
EXPOSE 8080
ENTRYPOINT ["/app/fetch"]