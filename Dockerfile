FROM golang:1.14-alpine AS golang-dist

WORKDIR /go/src/github.com/jvrmaia/circleci-playground/
EXPOSE 8080
COPY . .

RUN go get -d -v ./... \
    && go build -o golang-dist .

FROM alpine

RUN apk add --no-cache pcre ca-certificates tzdata && \
    apk --no-cache update && \
    mkdir ./up-to-date-zoneinfo && cd ./up-to-date-zoneinfo && \
    wget http://data.iana.org/time-zones/releases/tzdata2019c.tar.gz && \
    tar -xzvf tzdata2019c.tar.gz southamerica && \
    zic -d zoneinfo southamerica && \
    cd zoneinfo && \
    cp -r * /usr/share/zoneinfo/

COPY --from=golang-dist /go/src/github.com/jvrmaia/circleci-playground/golang-dist /usr/bin/golang-dist

CMD ["golang-dist"]