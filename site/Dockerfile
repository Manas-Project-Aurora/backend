FROM golang:1.22.2 AS builder

WORKDIR /app


COPY . .
WORKDIR /app/site
RUN go mod tidy


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go


FROM alpine:latest

WORKDIR /app

ARG YAML_PATH=site/dbconfig.yaml
ENV PORT=8080
COPY --from=builder /app/site/server /app/server
COPY --from=builder /app/${YAML_PATH} /app/dbconfig.yaml

EXPOSE ${PORT}

RUN apk add --no-cache bash nano

CMD ["/bin/sh"]


