FROM golang AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . /build/
RUN ls -la
RUN cd static && go run generatestatic.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
RUN mkdir public
COPY --from=builder /build/main .
COPY --from=builder /build/static/public/ ./public/
CMD ./main
