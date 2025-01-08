# Build Stage
FROM golang:1.22.1-alpine AS BuildStage
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
EXPOSE 8080
ENTRYPOINT ["./app"]