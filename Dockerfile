FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/go-template

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/go-template .
EXPOSE 3000
CMD ["./go-template"]