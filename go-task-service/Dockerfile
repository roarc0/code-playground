FROM golang:1.22-alpine AS build
ENV CGO_ENABLED=0
COPY . /app
WORKDIR /app
RUN go generate ./... && go build -v -o gtask

FROM alpine:3
COPY --from=build /app/gtask /gtask
CMD ["/gtask"]
