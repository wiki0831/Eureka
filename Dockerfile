# syntax=docker/dockerfile:1

FROM golang:1.17.3-alpine

WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -o /eureka

EXPOSE 8080

CMD [ "/eureka" ]