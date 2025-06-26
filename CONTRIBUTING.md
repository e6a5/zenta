# Contributing to zenta

Thank you for your interest in contributing to zenta! This document provides guidelines and information for contributors.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
- [Development Setup](#development-setup)
- [Coding Guidelines](#coding-guidelines)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)
- [Release Process](#release-process)

## Code of Conduct

This project adheres to a code of conduct that fosters an open and welcoming environment. Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md).

## How to Contribute

### Reporting Issues

- **Security vulnerabilities**: Please report security issues privately by emailing the maintainers
- **Bug reports**: Use the GitHub issue tracker with the bug report template
- **Feature requests**: Use the GitHub issue tracker with the feature request template
- **Documentation improvements**: Create an issue or submit a pull request directly

### Development Areas

We welcome contributions in:

- **Core functionality**: Mindfulness features, quotes, logging
- **User experience**: CLI interface improvements, help text
- **Performance**: Optimization, memory usage improvements
- **Testing**: Unit tests, integration tests, benchmarks
- **Documentation**: README, code comments, examples
- **Platform support**: Cross-platform compatibility
- **Accessibility**: Terminal compatibility, screen readers

## Development Setup

### Prerequisites

- Go 1.20 or later
- Git
- Make (optional, but recommended)

### Local Development

1. **Fork and clone the repository**
   ```bash
   git clone https://github.com/your-username/zenta.git
   cd zenta
   ```

2. **Install dependencies**
   ```bash
   make deps
   # or manually:
   go mod download
   ```

3. **Build and test**
   ```bash
   make build
   make test
   ```

4. **Run the application**
   ```bash
   ./zenta help
   ```

### Available Make Targets

```bash
make help          # Show all available targets
make build         # Build the binary
make test          # Run tests
make test-coverage # Generate coverage report
make lint          # Run linting
make fmt           # Format code
make check         # Run all checks (test, lint, vet)
make clean         # Clean build artifacts
```

## Coding Guidelines

### Go Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use `gofmt` and `goimports` for formatting
- Write clear, self-documenting code
- Add comments for exported functions and complex logic

### Project Principles

zenta follows the **Unix Philosophy**:

1. **"Do one thing, and do it well"** - Focus on mindfulness functionality only
2. **Composability** - Work well with other Unix tools via pipes and redirection
3. **Simplicity** - Prefer simple, obvious solutions over complex ones
4. **Text-based interface** - ASCII-safe output for any terminal
5. **Fail gracefully** - Clear error messages, degrade gracefully
6. **Respect user environment** - Standard Unix conventions

### Decision Framework

Before adding any feature, ask:

1. Does this align with helping users be more mindful?
2. Can users achieve this by combining zenta with other Unix tools?
3. Will this make zenta harder to understand or use?
4. Does this respect the "lightweight" promise?

If any answer is "no," reconsider the change.

### Code Organization

```
├── main.go                 # CLI entry point
├── internal/
│   ├── models/            # Data structures
│   ├── storage/           # File I/O operations
│   ├── quotes/            # Quote management
│   ├── stats/             # Analytics and charts
│   └── version/           # Version information
├── .github/workflows/     # CI/CD pipelines
└── docs/                  # Documentation (if needed)
```

## Testing

### Test Requirements

- All new features must include tests
- Maintain or improve test coverage
- Tests should be clear and focused
- Use table-driven tests where appropriate

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench

# Run specific package tests
go test ./internal/quotes/
```

### Test Categories

1. **Unit tests**: Test individual functions and methods
2. **Integration tests**: Test component interactions
3. **CLI tests**: Test command-line interface behavior

## Pull Request Process

### Before Submitting

1. **Create an issue** for substantial changes to discuss approach
2. **Fork the repository** and create a feature branch
3. **Write tests** for your changes
4. **Run the full test suite** and ensure all tests pass
5. **Run linting** and fix any issues
6. **Update documentation** if necessary

### Pull Request Guidelines

1. **Clear title and description**
   - Use present tense ("Add feature" not "Added feature")
   - Reference related issues (#123)
   - Explain what and why, not just how

2. **Small, focused changes**
   - One logical change per PR
   - Keep PRs small and reviewable
   - Split large changes into multiple PRs

3. **Quality checklist**
   - [ ] Tests pass locally
   - [ ] Linting passes
   - [ ] Documentation updated
   - [ ] No breaking changes (or clearly marked)
   - [ ] Follows project conventions

### Review Process

1. Automated checks (CI) must pass
2. At least one maintainer review required
3. Address review feedback promptly
4. Maintainer will merge when ready

## Release Process

Releases follow [Semantic Versioning](https://semver.org/):

- **MAJOR**: Incompatible API changes
- **MINOR**: New functionality (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Release Steps

1. Update `CHANGELOG.md` with release notes
2. Create and push a version tag: `git tag v1.2.3`
3. GitHub Actions automatically builds and publishes the release
4. Update package managers (Homebrew, etc.) as needed

## Getting Help

- **Questions**: Open a GitHub issue with the question label
- **Real-time chat**: Join our community discussions
- **Email**: Contact maintainers for sensitive issues

## Recognition

Contributors are recognized in:
- `CHANGELOG.md` for each release
- GitHub contributors list
- Release notes for significant contributions

Thank you for contributing to zenta! 🧘‍♂️ 