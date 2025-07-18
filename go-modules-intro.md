# Go 模块介绍

## 什么是 Go 模块

Go 模块是 Go 语言的依赖管理系统，从 Go 1.11 版本开始引入，在 Go 1.13 版本中成为默认方式。模块是相关 Go 包的集合，存储在文件树中，根目录包含 `go.mod` 文件。

## 核心概念

### go.mod 文件
- 定义模块路径和依赖需求
- 包含模块名称、Go 版本、依赖包及其版本信息

### go.sum 文件
- 包含依赖模块的校验和
- 确保依赖的完整性和安全性

## 基本命令

### 初始化模块
```bash
go mod init [模块名]
```

### 添加依赖
```bash
go get [包名]
```

### 移除未使用的依赖
```bash
go mod tidy
```

### 下载依赖到本地
```bash
go mod download
```

### 查看依赖图
```bash
go mod graph
```

### 验证依赖
```bash
go mod verify
```

## 版本管理

### 语义化版本
- 主版本.次版本.修订版本 (例如: v1.2.3)
- 主版本变更表示不兼容的 API 变更

### 版本选择
- `go get package@version` - 获取特定版本
- `go get package@latest` - 获取最新版本
- `go get package@master` - 获取指定分支

## 工作区模式

Go 1.18 引入了工作区模式，允许在多个模块间进行开发：

```bash
go work init [模块目录...]
go work use [模块目录]
```

## 最佳实践

1. **模块路径命名**：使用域名/组织/项目的形式
2. **版本标记**：使用 Git 标签进行版本管理
3. **依赖管理**：定期运行 `go mod tidy` 清理依赖
4. **安全性**：检查 `go.sum` 文件的完整性

## 示例

### 创建新模块
```bash
mkdir myproject
cd myproject
go mod init github.com/username/myproject
```

### go.mod 文件示例
```go
module github.com/username/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/stretchr/testify v1.8.4
)

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
    // 其他间接依赖...
)
```

## 迁移指南

### 从 GOPATH 迁移
1. 在项目根目录运行 `go mod init`
2. 运行 `go mod tidy` 自动发现依赖
3. 删除 vendor 目录（如果使用了 vendor）
4. 测试构建：`go build ./...`

Go 模块极大地简化了 Go 项目的依赖管理，提供了版本控制、安全性验证和构建可重现性。