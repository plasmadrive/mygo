FROM golang
ADD . /go/src/mygo/echosvc
WORKDIR /go/src/mygo/echosvc
RUN go install mygo/echosvc
ENTRYPOINT /go/bin/echosvc
EXPOSE 8080