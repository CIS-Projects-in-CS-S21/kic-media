FROM golang:1.15.6

ENV MONGO_URI=mongodb://mongo:27017
ENV PORT=50051

WORKDIR /usr/code

COPY go.mod /usr/code

RUN go mod download

COPY . /usr/code

ENTRYPOINT ["go", "test", "-v", "./..."]