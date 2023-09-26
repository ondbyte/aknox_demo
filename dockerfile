
FROM golang:latest AS build

WORKDIR /app

COPY ./ ./

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o ./main

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main ./

RUN chmod +x main

EXPOSE 8000

CMD ["./main"]
