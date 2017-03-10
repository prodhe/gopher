FROM golang

EXPOSE 70
VOLUME /public

COPY . /go/src/github.com/prodhe/gopher
RUN go install github.com/prodhe/gopher
COPY README.md /public

ENTRYPOINT ["/go/bin/gopher", "-d", "/public"]
