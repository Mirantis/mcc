FROM golang:1.13

RUN apt-get update && apt-get install -y golint

