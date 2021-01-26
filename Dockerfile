FROM golang

WORKDIR /src

COPY ./src /src

RUN go build .

EXPOSE 8080

CMD ["./mf-backend"]