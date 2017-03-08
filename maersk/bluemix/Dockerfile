FROM golang:1.8-alpine
ADD . /go/src/maersk/bluemix
RUN go install maersk/bluemix
ENTRYPOINT /go/bin/bluemix
EXPOSE 8080