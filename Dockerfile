FROM --platform=$BUILDPLATFORM golang:1.21-alpine3.19 AS build
WORKDIR /app
COPY . /app
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN go mod download && go mod verify
RUN go build -o /app/go-alert

FROM alpine:latest
WORKDIR /app
COPY app.env /app
COPY --from=build /app/go-alert /app
ENTRYPOINT [ "/app/go-alert"]
