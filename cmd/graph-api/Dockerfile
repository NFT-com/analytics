# syntax=docker/dockerfile:1

# NOTE: Must be run in the context of the repo's root directory

FROM golang:1.16-buster AS build
WORKDIR /app
# Copy source files and download Go modules.
COPY . .
RUN go mod download
# Build the application.
RUN go build -o /api ./cmd/graph-api

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /api /api
EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/api"]
