#FROM golang:alpine
#
#RUN apk add --update bash
#COPY app /
#
#EXPOSE 8080
#
#CMD ["/app"]



FROM golang:alpine as builder
RUN mkdir /build
ADD app/ /build/
WORKDIR /build
RUN apk add --update git
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
