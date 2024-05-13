FROM golang:alpine3.18 as build
WORKDIR /app/
COPY . .
RUN apk add --update make
RUN make build



FROM busybox:stable
WORKDIR /app/
COPY --from=build /app/bin/main ./
EXPOSE 8000
CMD [ "./main" ]
