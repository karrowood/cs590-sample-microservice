FROM golang:latest 
LABEL maintainer="Daniel Krolopp dkrolopp@purdue.edu"

ENV TERM=linux

WORKDIR /app
COPY main.go ./
RUN go get -u github.com/gin-gonic/gin
RUN go build -o main .
EXPOSE 8080
ENTRYPOINT ["./main"]

