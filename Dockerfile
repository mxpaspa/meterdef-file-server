FROM golang:1.16

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
# COPY go.sum .
RUN go mod download
# RUN go mod download
# RUN go mod verify

COPY . .

RUN go build -o ./build/meterdef-file-server

EXPOSE 8100
CMD ["./build/meterdef-file-server"]
