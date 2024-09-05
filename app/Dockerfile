FROM golang:1.23.0-alpine as build

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./
RUN go build -o build

FROM scratch

EXPOSE 8080

COPY --from=build /opt/app/build /

ENTRYPOINT ["/build"]
