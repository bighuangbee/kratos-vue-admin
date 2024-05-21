FROM golang:1.20 AS builder

WORKDIR /src
COPY . .

#在宿主机手动复制
#cp ~/.ssh/id_rsa.pub .
#cp ~/.ssh/id_rsa .
#ADD id_rsa /root/.ssh/id_rsa
#ADD id_rsa.pub /root/.ssh/id_rsa.pub

RUN go env -w GOPROXY=https://goproxy.cn,direct && \
  go mod download

# -s: 省略符号表和调试信息
# -w: 省略DWARF符号表
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o srv-bin ./app/admin/cmd/server/


FROM alpine:3.7

COPY --from=builder /src/srv-bin /app/
COPY --from=builder /src/app/admin/configs/config.yaml /app/conf/config.yaml
COPY --from=builder /src/app/admin/configs/authz/casbin_model.conf /app/conf/authz/casbin_model.conf

WORKDIR /app
RUN mkdir logs

CMD ["./srv-bin", "-conf", "/app/conf/config.yaml"]
