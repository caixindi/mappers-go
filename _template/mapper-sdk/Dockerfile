ARG BASE=golang:1.17-alpine3.15
FROM ${BASE} AS builder
WORKDIR /mapper-go-sdk
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --update make gcc
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN make build
FROM alpine:3.14
WORKDIR /
COPY --from=builder /mapper-go-sdk/example/Template/bin/ /bin
COPY --from=builder /mapper-go-sdk/example/Template/res /res
ENTRYPOINT ["/bin/main"]
CMD ["--v=4"]
