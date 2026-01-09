# Contributing to GoSensitive

Thank you for your interest in contributing to GoSensitive! This document provides guidelines for contributing to the project.

## Code of Conduct

Be respectful and inclusive. We welcome contributors from all backgrounds.

## How to Contribute

### Reporting Bugs

1. Check if the bug has already been reported in Issues
2. If not, create a new issue with:
   - Clear title and description
   - Steps to reproduce
   - Expected vs actual behavior
   - Go version and OS

### Suggesting Features

1. Check if the feature has already been suggested
2. Create a new issue describing:
   - The problem you're trying to solve
   - Your proposed solution
   - Any alternatives you've considered

### Pull Requests

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes
4. Add tests for new functionality
5. Run tests: `make test`
6. Run linter: `make lint`
7. Commit with clear messages
8. Push to your fork
9. Create a Pull Request

## Development Setup

```bash
# Clone your fork
git clone https://github.com/yourusername/gosensitive.git
cd gosensitive

# Install dependencies
go mod download

# Run tests
make test

# Run benchmarks
make bench
```

## Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting
- Add comments for exported functions
- Write meaningful commit messages

## Testing

- Write unit tests for new features
- Maintain or improve test coverage
- Include benchmark tests for performance-critical code

## Documentation

- Update README.md if needed
- Add godoc comments for exported APIs
- Include code examples for new features

## License

By contributing, you agree that your contributions will be licensed under the MIT License.


