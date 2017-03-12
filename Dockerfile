FROM golang

ENV GOPHER_PORT 70
ENV GOPHER_ADDRESS localhost
ENV GOPHER_DIR /public

EXPOSE 70
VOLUME /public

COPY . /go/src/github.com/prodhe/gopher
RUN go install github.com/prodhe/gopher
COPY README.md /public

CMD /go/bin/gopher -d ${GOPHER_DIR} -p ${GOPHER_PORT} -a ${GOPHER_ADDRESS}
