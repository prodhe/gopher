FROM golang

ENV GOPHER_ADDRESS localhost

EXPOSE 70
VOLUME /public

COPY . /go/src/gopher
RUN go install gopher
COPY README.md /public

CMD /go/bin/gopher -d /public/ -p 70 -a ${GOPHER_ADDRESS}
