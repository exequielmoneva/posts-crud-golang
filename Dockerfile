#get a base image
FROM golang:1.17-alpine


WORKDIR /app
COPY . .

RUN go get -d -v
RUN go build -v

EXPOSE 3000

CMD ["./posts-crud-golang"]


