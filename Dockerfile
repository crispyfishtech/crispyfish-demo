FROM golang:1.24-alpine AS app
RUN apk add -U build-base git
COPY app /go/src/app
WORKDIR /go/src/app
RUN go build -mod=vendor -a -v -tags 'netgo' -ldflags '-w -extldflags -static' -o crispyfish-demo .

FROM alpine:3.22
RUN apk add -U --no-cache curl
COPY app/static /static
COPY --from=app /go/src/app/crispyfish-demo /bin/crispyfish-demo
EXPOSE 8080
ENTRYPOINT ["/bin/crispyfish-demo"]
