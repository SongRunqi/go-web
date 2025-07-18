# Go 项目入门指南

## 1. 安装 Go

### Windows 安装

1. 访问 Go 官网：https://golang.org/dl/
2. 下载适合 Windows 的安装包（.msi 文件）
3. 双击安装包，按照向导完成安装
4. 安装完成后，Go 会自动添加到系统 PATH 中

**验证安装：**
```cmd
go version
```

### macOS 安装

**方法一：使用 Homebrew（推荐）**
```bash
# 安装 Homebrew（如果还未安装）
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# 安装 Go
brew install go
```

**方法二：使用安装包**
1. 访问 Go 官网：https://golang.org/dl/
2. 下载适合 macOS 的安装包（.pkg 文件）
3. 双击安装包，按照向导完成安装

**验证安装：**
```bash
go version
```

### Linux 安装

**Ubuntu/Debian：**
```bash
sudo apt update
sudo apt install golang-go
```

**CentOS/RHEL：**
```bash
sudo yum install golang
```

**验证安装：**
```bash
go version
```

## 2. 验证安装

```bash
go version
```

## 3. 设置环境变量

```bash
# 查看 Go 环境配置
go env

# 设置 GOPATH（可选，Go 1.11+ 推荐使用模块）
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## 4. 创建第一个 Go 项目

### 方法一：使用 Go Modules（推荐）

```bash
# 1. 创建项目目录
mkdir my-go-project
cd my-go-project

# 2. 初始化 Go 模块
go mod init github.com/username/my-go-project

# 3. 创建 main.go 文件
touch main.go
```

### 方法二：在 GOPATH 中创建项目

```bash
# 创建项目目录
mkdir -p $GOPATH/src/github.com/username/my-go-project
cd $GOPATH/src/github.com/username/my-go-project
```

## 5. 编写第一个程序

创建 `main.go` 文件：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## 6. 运行程序

```bash
# 直接运行
go run main.go

# 或者先编译再运行
go build
./my-go-project  # Linux/macOS
# my-go-project.exe  # Windows
```

## 7. Go 项目结构示例

```
my-go-project/
├── go.mod           # 模块定义文件
├── go.sum           # 依赖校验和文件
├── main.go          # 主程序入口
├── internal/        # 内部包（不对外暴露）
│   └── config/
├── pkg/             # 可供外部使用的包
│   └── utils/
├── cmd/             # 命令行程序
│   └── server/
├── web/             # Web 相关文件
│   ├── handlers/
│   └── middleware/
└── README.md
```

## 8. 常用 Go 命令

```bash
# 初始化模块
go mod init <module-name>

# 添加依赖
go get <package>

# 移除未使用的依赖
go mod tidy

# 下载依赖到本地缓存
go mod download

# 运行程序
go run main.go

# 编译程序
go build

# 运行测试
go test

# 格式化代码
go fmt

# 检查代码问题
go vet

# 安装程序到 GOBIN
go install
```

## 9. Web 项目示例

创建一个简单的 Web 服务器：

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Web!")
    })

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

运行：
```bash
go run main.go
# 访问 http://localhost:8080
```

## 10. 下一步

- 学习 Go 语法基础
- 了解 Go 的并发模型（goroutines 和 channels）
- 探索流行的 Go 框架（如 Gin、Echo）
- 学习 Go 的测试框架
- 了解 Go 的包管理和模块系统