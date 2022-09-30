FROM golang:alpine
WORKDIR $GOPATH/src/p_manager
ADD . ./
ENV GO111MODULE=on
RUN go build -o p_manager .
EXPOSE 80
ENTRYPOINT  ["./p_manager"]
