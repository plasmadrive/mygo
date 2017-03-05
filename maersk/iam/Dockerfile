FROM golang:1.8-alpine
ADD .  /go/src/maersk/iam
RUN go install maersk/iam
ENTRYPOINT /go/bin/iam
EXPOSE 8080