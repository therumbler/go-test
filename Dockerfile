FROM golang AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/main .
CMD ./main
