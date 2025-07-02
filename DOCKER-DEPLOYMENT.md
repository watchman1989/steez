# Steez 项目 Docker 部署指南

## 项目概述

Steez 是一个基于 Gin 框架的 Go Web 服务，使用 Go 1.23.4 版本。

## Dockerfile 说明

### 构建阶段优化

```dockerfile
# 使用多阶段构建
FROM golang:1.23-alpine AS builder
```

**优化点：**
- ✅ 使用 Alpine Linux 减小基础镜像大小
- ✅ 多阶段构建分离构建和运行环境
- ✅ 静态链接确保在 Cloud Run 中正常运行

### 运行阶段优化

```dockerfile
# 使用 distroless 镜像
FROM gcr.io/distroless/static-debian11:nonroot
```

**优化点：**
- ✅ 最小化攻击面
- ✅ 非 root 用户运行
- ✅ 不包含 shell 和调试工具

### 构建参数

```dockerfile
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a -installsuffix cgo \
    -ldflags="-w -s" \
    -o main .
```

**参数说明：**
- `CGO_ENABLED=0`: 禁用 CGO，确保静态链接
- `GOOS=linux`: 目标操作系统
- `GOARCH=amd64`: 目标架构
- `-ldflags="-w -s"`: 去除调试信息，减小二进制大小

## 部署方式

### 方式一：使用 Dockerfile（推荐）

```bash
# 1. 构建镜像
docker build -t gcr.io/YOUR_PROJECT_ID/steez:latest .

# 2. 推送到 Container Registry
docker push gcr.io/YOUR_PROJECT_ID/steez:latest

# 3. 部署到 Cloud Run
gcloud run deploy steez \
  --image gcr.io/YOUR_PROJECT_ID/steez:latest \
  --region asia-east1 \
  --platform managed \
  --allow-unauthenticated \
  --port 8080
```

### 方式二：使用 Cloud Buildpack

```bash
# 直接部署，无需 Dockerfile
gcloud run deploy steez \
  --source . \
  --region asia-east1 \
  --platform managed \
  --allow-unauthenticated \
  --port 8080
```

### 方式三：使用 Cloud Build

```bash
# 使用 cloudbuild.yaml
gcloud builds submit --config cloudbuild.yaml

# 或使用 Cloud Buildpack 配置
gcloud builds submit --config cloudbuild-buildpack.yaml
```

## 配置说明

### 端口配置

**重要：** 项目代码中使用的是 8080 端口，与 Cloud Run 的默认端口一致。

```go
// srv/srv.go
err := r.Run(fmt.Sprintf(":%s", "8080"))
```

### 环境变量

```dockerfile
ENV PORT=8080
ENV TZ=Asia/Shanghai
```

### 健康检查

```dockerfile
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD ["/app/main", "-health-check"] || exit 1
```

**注意：** 需要在实际应用中实现健康检查端点。

## 性能优化

### 1. 镜像大小优化

- 使用多阶段构建
- 使用 distroless 基础镜像
- 去除调试信息
- 使用 .dockerignore 排除不必要文件

### 2. 构建速度优化

- 利用 Docker 层缓存
- 先复制 go.mod 和 go.sum
- 使用 .dockerignore 减少构建上下文

### 3. 运行时优化

- 静态链接二进制文件
- 非 root 用户运行
- 最小化攻击面

## 安全最佳实践

### 1. 基础镜像安全

```dockerfile
# 使用官方 distroless 镜像
FROM gcr.io/distroless/static-debian11:nonroot
```

### 2. 非 root 用户

```dockerfile
# distroless 镜像默认使用非 root 用户
# 无需额外配置
```

### 3. 最小化攻击面

- 不包含 shell
- 不包含调试工具
- 只包含运行时必需的文件

## 故障排除

### 1. 构建失败

```bash
# 检查 Dockerfile 语法
docker build --dry-run .

# 查看详细构建日志
docker build --progress=plain -t steez .
```

### 2. 运行时错误

```bash
# 查看容器日志
docker logs <container_id>

# 进入容器调试（如果使用 Alpine 镜像）
docker run -it --rm steez sh
```

### 3. 端口问题

确保代码中的端口与 Dockerfile 和 Cloud Run 配置一致：

```go
// 代码中使用 8080 端口
r.Run(":8080")
```

```dockerfile
# Dockerfile 中暴露 8080 端口
EXPOSE 8080
ENV PORT=8080
```

```yaml
# cloudbuild.yaml 中配置 8080 端口
- '--port'
- '8080'
```

## 监控和日志

### 1. 查看 Cloud Run 日志

```bash
# 实时日志
gcloud run services logs tail steez --region asia-east1

# 历史日志
gcloud logging read "resource.type=cloud_run_revision AND resource.labels.service_name=steez"
```

### 2. 监控指标

- 请求量
- 错误率
- 响应时间
- 实例数
- 内存使用率

## 扩展配置

### 1. 自定义环境变量

```bash
gcloud run services update steez \
  --set-env-vars KEY1=VALUE1,KEY2=VALUE2 \
  --region asia-east1
```

### 2. 资源配置

```bash
gcloud run services update steez \
  --memory 1Gi \
  --cpu 2 \
  --max-instances 20 \
  --region asia-east1
```

### 3. 自动扩缩容

```bash
gcloud run services update steez \
  --min-instances 1 \
  --max-instances 50 \
  --region asia-east1
```

## 总结

这个 Dockerfile 配置提供了：

✅ **优化的构建过程**
✅ **安全的运行环境**
✅ **最小的镜像大小**
✅ **快速的启动时间**
✅ **Cloud Run 兼容性**

通过使用多阶段构建和 distroless 镜像，确保了应用在 Google Cloud Run 上的最佳性能和安全性。 