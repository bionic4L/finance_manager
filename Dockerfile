FROM golang:alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o finance_manager ./cmd/app/main.go

CMD ["./finance_manager"]