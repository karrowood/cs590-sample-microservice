FROM golang:latest 
LABEL maintainer="Kyle Arrowood karrowo@purdue.edu"

ENV TERM=linux

WORKDIR /app
COPY main.go ./
RUN go mod init main.go 
RUN go get -u github.com/gin-gonic/gin
RUN go build -o main .
EXPOSE 8082
ENTRYPOINT ["./main"]

