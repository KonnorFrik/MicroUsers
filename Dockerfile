FROM golang:bookworm

EXPOSE 8000 4430

COPY *.go go.mod /go/src/
COPY init_0.sh /go/src/
WORKDIR /go/src
RUN bash init_0.sh
WORKDIR /go
USER user

CMD [ "bin/microusers" ]
