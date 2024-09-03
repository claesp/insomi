FROM golang:latest AS build
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY *.go .
RUN CGO_ENABLED=0 go build -ldflags "-s -w -X main.Version=0.0.0-ci" -o /insomi

FROM scratch AS run
WORKDIR /
COPY lib ./lib
COPY views ./views
COPY --from=build /insomi /insomi
EXPOSE 8080
ENTRYPOINT ["/insomi"]
