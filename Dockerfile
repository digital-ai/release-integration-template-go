# Step 1
FROM golang:1.19 AS build

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -ldflags "-s -w" -o /release-integration-template-go

#Step 2 - UPX Compression
FROM alpine:3.17 AS compress

RUN apk add upx

COPY --from=build /release-integration-template-go /

RUN upx /release-integration-template-go

#Step 3
FROM gcr.io/distroless/static-debian11

ENV INPUT_LOCATION=/input
ENV OUTPUT_LOCATION=/output

COPY --from=compress release-integration-template-go /

ENTRYPOINT ["/release-integration-template-go"]
