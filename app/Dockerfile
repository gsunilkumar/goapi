FROM golang:1.18.0-stretch

WORKDIR /app
COPY . ./

RUN go mod download
RUN go build -o /app/main

EXPOSE 7004

CMD [ "/app/main" ]