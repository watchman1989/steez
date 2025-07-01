#!/bin/bash

# 设置构建参数
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

echo "开始构建 steez 应用..."

# 清理旧的构建产物
if [ -f "main" ]; then
    echo "清理旧的构建产物..."
    rm main
fi

# 构建应用
echo "编译 Go 应用..."
go build -a -installsuffix cgo -o main .

# 检查构建结果
if [ $? -eq 0 ]; then
    echo "✅ 构建成功！"
    echo "二进制文件: $(pwd)/main"
    echo "文件大小: $(ls -lh main | awk '{print $5}')"
    
    # 显示文件信息
    echo "文件信息:"
    file main
    echo "依赖库:"
    ldd main 2>/dev/null || echo "静态链接，无外部依赖"
else
    echo "❌ 构建失败！"
    exit 1
fi

echo ""
echo "现在可以运行以下命令部署到 Google Cloud:"
echo "gcloud builds submit --config cloudbuild.yaml" 