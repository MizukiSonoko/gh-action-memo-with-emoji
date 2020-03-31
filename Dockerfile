FROM golang:1.13-buster as build
WORKDIR /app
ADD ./app /app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/gen-memo

FROM alpine
COPY --from=build /app /
ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
