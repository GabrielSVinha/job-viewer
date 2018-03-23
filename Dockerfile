FROM golang

WORKDIR /

ADD viewer /

CMD ["./viewer"]
