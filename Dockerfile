FROM golang:bookworm

EXPOSE 8000 4430

COPY *.go go.mod /go/src/
WORKDIR /go/src
RUN go mod tidy && go build -o ../bin/ . && useradd -r user -p user
WORKDIR /go
USER user

CMD [ "bin/microusers" ]
