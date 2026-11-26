FROM golang:1.27.4 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/glowl ./cmd/Glowl

FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/glowl /app/glowl
EXPOSE 8000
CMD ["/app/glowl"]