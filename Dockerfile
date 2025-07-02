# 构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的系统依赖
RUN apk add --no-cache git ca-certificates tzdata

# 复制 go mod 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
# 使用静态链接，确保在 Cloud Run 环境中正常运行
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a -installsuffix cgo \
    -ldflags="-w -s" \
    -o main .

# 运行阶段 - 使用 distroless 镜像以获得最小的攻击面
FROM gcr.io/distroless/static-debian11:nonroot

# 设置时区
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .

# 暴露端口 (注意：这里使用 8080，与代码中的端口一致)
EXPOSE 8080

# 设置环境变量
ENV PORT=8080

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD ["/app/main", "-health-check"] || exit 1

# 启动应用
CMD ["/app/main"] 