ARG API_PORT=8000

FROM golang:1.15 AS builder
WORKDIR /src/
ADD . .
RUN make go-build-prod

FROM alpine:latest AS production
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait
WORKDIR /root/
COPY --from=builder /src/build/api .
COPY --from=builder /src/scripts ./scripts
RUN mkdir -p storage/logs
EXPOSE $API_PORT
CMD ["sh", "-c", "/wait && ./api"]
