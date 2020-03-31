FROM golang:1.13-buster as build
WORKDIR /app
ADD ./app /app
RUN make 
RUN ls /app

FROM alpine
WORKDIR /app
COPY --from=build /app /
RUN ls /app
ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
