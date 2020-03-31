FROM golang:1.13-buster as build
WORKDIR /app
ADD ./app /app
RUN make 

FROM alpine
WORKDIR /app
COPY --from=build /app /
ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
