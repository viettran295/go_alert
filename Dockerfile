FROM --platform=$BUILDPLATFORM golang:1.21 AS build

WORKDIR /app
COPY . /app
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN go mod download && go mod verify
RUN go build -o /app/go-alert

ENTRYPOINT [ "/app/go-alert"]
