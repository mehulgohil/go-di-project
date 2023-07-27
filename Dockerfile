# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage

WORKDIR /go/src/app
ADD . /go/src/app

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o /di-project

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /di-project /di-project

USER nonroot:nonroot

CMD ["/di-project"]