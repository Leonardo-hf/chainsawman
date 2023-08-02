FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY graph graph
COPY common common
COPY graph/cmd/api/etc /app/etc
ADD go.mod .
ADD go.sum .
RUN go mod tidy && go build -ldflags="-s -w" -o /app/chs-graph graph/cmd/api/graph.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/chs-graph /app/chs-graph
COPY --from=builder /app/etc /app/graph/cmd/api/etc

CMD ["./chs-graph"]
