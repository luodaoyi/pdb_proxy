FROM golang:1.22-alpine as builder
MAINTAINER szdyg "szdyg@outlook.com"

WORKDIR /usr/src/app

ARG GOPROXY=https://goproxy.cn,direct

RUN apk add --no-cache ca-certificates tzdata
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o server


FROM scratch as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/src/app/server /opt/app/
COPY --from=builder /usr/src/app/pdb_proxy.ini /opt/app/
CMD ["/opt/app/server"]
