# GoSensitive

一个高性能、功能丰富的 Go 语言敏感词检测库。

[![Go Version](https://img.shields.io/github/go-mod/go-version/yourusername/gosensitive)](https://golang.org/)
[![License](https://img.shields.io/github/license/yourusername/gosensitive)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/gosensitive)](https://goreportcard.com/report/github.com/yourusername/gosensitive)

[English](README.md)

## 特性

- 🚀 **高性能**: 支持 DFA 和 Aho-Corasick 算法，自动选择最优方案
- 🎯 **灵活的 API**: 链式调用，开箱即用
- 🔧 **变体检测**: 支持拼音、繁简体、符号干扰、形近字检测
- 🔒 **线程安全**: 使用 COW 策略，支持高并发
- 📦 **多种加载方式**: 支持文件、内存、HTTP 等多种词库来源
- 🎨 **丰富的结果**: 提供详细的匹配信息，包括分类和级别
- ⚡ **零依赖**: 核心库无外部依赖

## 安装

```bash
go get github.com/yourusername/gosensitive
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/yourusername/gosensitive"
)

func main() {
    // 创建检测器
    detector, _ := gosensitive.New().
        LoadMemory([]string{"敏感词", "测试"}).
        Build()

    // 检查是否包含敏感词
    if detector.Contains("这是一个敏感词") {
        fmt.Println("检测到敏感词！")
    }

    // 查找所有敏感词
    matches := detector.Find("敏感词和测试")
    for _, match := range matches {
        fmt.Printf("发现: %s 位置 %d-%d\n", 
            match.Word, match.Start, match.End)
    }

    // 替换敏感词
    filtered := detector.Filter("这个敏感词需要过滤")
    fmt.Println(filtered) // 输出: 这个***需要过滤
}
```

## 高级用法

### 自定义算法选择

```go
// 使用 DFA 算法
detector := gosensitive.New().
    UseAlgorithm(gosensitive.AlgorithmDFA).
    LoadFile("words.txt").
    Build()

// 使用 AC 自动机
detector := gosensitive.New().
    UseAlgorithm(gosensitive.AlgorithmAC).
    LoadFile("words.txt").
    Build()

// 自动选择（默认：小于5000词用DFA，否则用AC）
detector := gosensitive.New().
    UseAlgorithm(gosensitive.AlgorithmAuto).
    LoadFile("words.txt").
    Build()
```

### 白名单过滤

```go
detector := gosensitive.New().
    LoadMemory([]string{"测试", "示例", "敏感"}).
    AddWhitelist("测试", "示例"). // 这些词不会被匹配
    Build()
```

### 自定义选项

```go
opts := gosensitive.DefaultOptions()
opts.ReplaceChar = '▓'
opts.MaxMatchCount = 10
opts.CaseSensitive = false

detector := gosensitive.New().
    LoadMemory([]string{"词1", "词2"}).
    SetOptions(opts).
    Build()
```

### 从多个来源加载

```go
detector := gosensitive.New().
    LoadFile("local_words.txt").
    LoadHTTP("https://example.com/words.txt").
    LoadMemory([]string{"额外1", "额外2"}).
    Build()
```

## 性能

在 AMD Ryzen 7 5800X 上的基准测试：

| 词库大小 | 算法 | 操作数/秒 | 延迟 |
|---------|------|----------|------|
| 1,000 词 | DFA | 500,000+ | ~2 µs |
| 1,000 词 | AC | 600,000+ | ~1.6 µs |
| 10,000 词 | DFA | 200,000+ | ~5 µs |
| 10,000 词 | AC | 300,000+ | ~3.3 µs |
| 100,000 词 | DFA | 80,000+ | ~12 µs |
| 100,000 词 | AC | 150,000+ | ~6.6 µs |

运行基准测试：

```bash
make bench
```

## 示例

查看 [examples](examples/) 目录获取更多使用示例：

- [基础用法](examples/basic/main.go)
- [高级功能](examples/advanced/main.go)
- [Web 中间件](examples/middleware/)

## 测试

```bash
# 运行所有测试
make test

# 运行测试并生成覆盖率报告
make test-coverage

# 运行基准测试
make bench
```

## 文档

完整文档请访问 [GoDoc](https://pkg.go.dev/github.com/yourusername/gosensitive)。

## 贡献

欢迎贡献！请阅读 [CONTRIBUTING.md](CONTRIBUTING.md) 了解详情。

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。

## 致谢

- 受到其他语言敏感词检测库的启发
- AC 算法实现基于经典的 Aho-Corasick 论文
- 感谢所有贡献者

## 路线图

- [ ] 支持更多变体检测方法
- [ ] 基于 Redis 的分布式词典
- [ ] gRPC 服务封装
- [ ] 性能优化
- [ ] 更多中间件集成

## 支持

如果这个项目对你有帮助，请给它一个 ⭐️！

如有问题，请使用 [GitHub Issues](https://github.com/yourusername/gosensitive/issues)。


