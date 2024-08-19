FROM golang:alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o finance_manager ./cmd/app/main.go

EXPOSE 7778

CMD ["./finance_manager"]