# Build
FROM golang:1.20-alpine AS build-stage

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download
COPY . ./

COPY entrypoint.sh /tmp/entrypoint.sh
RUN chmod +x /tmp/entrypoint.sh

RUN CGO_ENABLED=0 go build -o /tmp/servus-api

# Final
FROM alpine:3.18.2

WORKDIR /

COPY --from=build-stage /tmp/servus-api /servus-api
COPY --from=build-stage /tmp/entrypoint.sh /entrypoint.sh

ENV PORT=8080
ENV MODE=""

RUN ls /

ENTRYPOINT ["/entrypoint.sh"]
