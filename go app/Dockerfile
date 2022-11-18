FROM golang:latest

EXPOSE 8080

RUN mkdir /app
## We copy main.go in the root directory
## into our /app directory
ADD main.go /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go mod init main.go
RUN go build -o main .
RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install unzip -y
RUN wget -c https://github.com/GauravDawra/Beast/releases/download/v1.1.0/beast-linux-64bit.zip
RUN unzip beast-linux-64bit.zip
RUN chmod +x beast
RUN mv beast /usr/local/bin/beast
RUN rm beast-linux-64bit.zip

## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]


