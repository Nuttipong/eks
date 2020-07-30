FROM        golang:alpine as builder
RUN         mkdir /build
WORKDIR     /build
COPY        . .
RUN         ls -la 
RUN         CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM        scratch
COPY        --from=builder /build/ /app/
WORKDIR     /app
#EXPOSE      9001
CMD         ["./main"]