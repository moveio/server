FROM golang

RUN go get github.com/flowup/owl/cmd/owl
RUN go get github.com/smartystreets/goconvey
