FROM golang:1.22-alpine AS build
ENV CGO_ENABLED=0
WORKDIR /src
COPY . .
RUN go generate ./... && go build -v -ldflags "-s -w" -o app ./cmd

FROM alpine:3
COPY --from=build /src/app /app
CMD ["/app"]
