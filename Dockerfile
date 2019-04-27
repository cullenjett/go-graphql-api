FROM golang:1.12

WORKDIR /app

RUN go get github.com/cespare/reflex

COPY reflex.conf /

CMD ["reflex", "-c", "/reflex.conf"]
