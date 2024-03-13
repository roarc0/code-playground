FROM golang:1.22-alpine AS build
ENV CGO_ENABLED=0
COPY . /src
WORKDIR /src
RUN go generate ./... && go build -v -o app

FROM alpine:3
COPY --from=build /src/app /app
CMD ["/app"]
