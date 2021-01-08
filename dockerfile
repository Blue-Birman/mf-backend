FROM golang

WORKDIR /build

COPY . /build

RUN go build

EXPOSE 8080

CMD ["./mf-backend"]