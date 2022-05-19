FROM golang:latest 

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /root/go/src/coffe_go

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Export necessary port
EXPOSE 4000

CMD ["./main"]