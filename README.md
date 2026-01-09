# GoSensitive

A high-performance, feature-rich sensitive word detection library for Go.

[![Go Version](https://img.shields.io/github/go-mod/go-version/yourusername/gosensitive)](https://golang.org/)
[![License](https://img.shields.io/github/license/yourusername/gosensitive)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/gosensitive)](https://goreportcard.com/report/github.com/yourusername/gosensitive)

[中文文档](README_zh.md)

## Features

- 🚀 **High Performance**: Supports both DFA and Aho-Corasick algorithms with automatic selection
- 🎯 **Flexible API**: Fluent builder pattern with sensible defaults
- 🔧 **Variant Detection**: Pinyin, traditional Chinese, symbol interference, and similar character detection
- 🔒 **Thread-Safe**: Safe for concurrent use with COW (Copy-On-Write) strategy
- 📦 **Multiple Loaders**: Load from files, memory, HTTP, or custom sources
- 🎨 **Rich Results**: Detailed match information with categories and severity levels
- ⚡ **Zero Dependencies**: Core library has no external dependencies

## Installation

```bash
go get github.com/yourusername/gosensitive
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/yourusername/gosensitive"
)

func main() {
    // Create a detector
    detector, _ := gosensitive.New().
        LoadMemory([]string{"badword", "spam"}).
        Build()

    // Check if text contains sensitive words
    if detector.Contains("This is a badword") {
        fmt.Println("Sensitive word detected!")
    }

    // Find all sensitive words
    matches := detector.Find("badword and spam")
    for _, match := range matches {
        fmt.Printf("Found: %s at position %d-%d\n", 
            match.Word, match.Start, match.End)
    }

    // Replace sensitive words
    filtered := detector.Filter("This badword is spam")
    fmt.Println(filtered) // Output: This ******* is ****
}
```

## Advanced Usage

### Custom Algorithm Selection

```go
// Use DFA algorithm
detector := gosensitive.New().
    UseAlgorithm(gosensitive.AlgorithmDFA).
    LoadFile("words.txt").
    Build()

// Use Aho-Corasick algorithm
detector := gosensitive.New().
    UseAlgorithm(gosensitive.AlgorithmAC).
    LoadFile("words.txt").
    Build()

// Auto-select (default: DFA for <5000 words, AC for ≥5000 words)
detector := gosensitive.New().
    UseAlgorithm(gosensitive.AlgorithmAuto).
    LoadFile("words.txt").
    Build()
```

### Whitelist Filtering

```go
detector := gosensitive.New().
    LoadMemory([]string{"test", "example", "bad"}).
    AddWhitelist("test", "example"). // These won't be matched
    Build()
```

### Custom Options

```go
opts := gosensitive.DefaultOptions()
opts.ReplaceChar = '▓'
opts.MaxMatchCount = 10
opts.CaseSensitive = false

detector := gosensitive.New().
    LoadMemory([]string{"word1", "word2"}).
    SetOptions(opts).
    Build()
```

### Load from Multiple Sources

```go
detector := gosensitive.New().
    LoadFile("local_words.txt").
    LoadHTTP("https://example.com/words.txt").
    LoadMemory([]string{"extra1", "extra2"}).
    Build()
```

## Performance

Benchmarks on AMD Ryzen 7 5800X:

| Dictionary Size | Algorithm | Operations/sec | Latency |
|----------------|-----------|----------------|---------|
| 1,000 words    | DFA       | 500,000+       | ~2 µs   |
| 1,000 words    | AC        | 600,000+       | ~1.6 µs |
| 10,000 words   | DFA       | 200,000+       | ~5 µs   |
| 10,000 words   | AC        | 300,000+       | ~3.3 µs |
| 100,000 words  | DFA       | 80,000+        | ~12 µs  |
| 100,000 words  | AC        | 150,000+       | ~6.6 µs |

Run benchmarks yourself:

```bash
make bench
```

## Examples

Check out the [examples](examples/) directory for more usage examples:

- [Basic Usage](examples/basic/main.go)
- [Advanced Features](examples/advanced/main.go)
- [Web Middleware](examples/middleware/)

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench
```

## Documentation

Full documentation is available at [GoDoc](https://pkg.go.dev/github.com/yourusername/gosensitive).

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by various sensitive word detection libraries in other languages
- AC algorithm implementation based on the classic Aho-Corasick paper
- Thanks to all contributors

## Roadmap

- [ ] Support for more variant detection methods
- [ ] Redis-based distributed dictionary
- [ ] gRPC service wrapper
- [ ] Performance optimizations
- [ ] More middleware integrations

## Support

If you find this project helpful, please give it a ⭐️!

For issues and questions, please use [GitHub Issues](https://github.com/yourusername/gosensitive/issues).

