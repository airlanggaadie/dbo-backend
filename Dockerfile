FROM golang:1.22.1-bullseye as builder

WORKDIR /app

COPY . .

RUN go build .

FROM debian:bullseye

WORKDIR /app

RUN apt-get update && apt-get upgrade -y && apt-get install -y curl

COPY --from=builder /app/assignment-test .

ENV PORT=3000

EXPOSE ${PORT}

CMD ["./assignment-test"]