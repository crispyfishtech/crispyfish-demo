FROM golang:1.24-alpine AS app
RUN apk add -U build-base git
COPY app /go/src/app
WORKDIR /go/src/app
ENV GO111MODULE=on
RUN go build -mod=vendor -a -v -tags 'netgo' -ldflags '-w -extldflags -static' -o crispyfish-demo .

FROM alpine:latest
RUN apk add -U --no-cache curl
COPY app/static /static
COPY --from=app /go/src/app/crispyfish-demo /bin/crispyfish-demo
COPY app/templates /templates
EXPOSE 8080
ENTRYPOINT ["/bin/crispyfish-demo"]
