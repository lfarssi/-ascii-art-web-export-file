FROM golang:latest

LABEL maintainer="medlfarssi10@gmail.com" \
      version="1.0" \
      description="Dockirizing my  Go application"

WORKDIR "/app"

COPY . .

RUN go build -o myapp . 

EXPOSE 3000

CMD ["./myapp"]