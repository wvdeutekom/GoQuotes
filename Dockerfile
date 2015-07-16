FROM google/golang

WORKDIR /gopath/src/github.com/wvdeutekom/webhookproject
ADD . /gopath/src/github.com/wvdeutekom/webhookproject/

# go get all of the dependencies
RUN go get code.google.com/p/gcfg
RUN go get gopkg.in/dancannon/gorethink.v1
RUN go get github.com/gorilla/schema
RUN go get github.com/labstack/echo

RUN go get github.com/wvdeutekom/webhookproject

EXPOSE 8000
CMD []
ENTRYPOINT ["/gopath/bin/webhookproject"]
