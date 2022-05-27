FROM golang:1.18.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY ./services/story ./services/story
COPY ./packages ./packages

EXPOSE 8081

WORKDIR /app/services/story

RUN go build -o story-service

ENTRYPOINT ./story-service
