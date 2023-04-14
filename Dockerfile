# Build environment
# ----------------------
FROM golang:1.20-alpine as build-env
WORKDIR /myapp

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd

# Deployment environment
# ----------------------
FROM alpine

COPY --from=build-env /myapp/bin/api /myapp/
COPY --from=build-env /myapp/.env .

EXPOSE 8081
CMD ["/myapp/api"]