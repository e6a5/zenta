# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Version command (`zenta version`, `zenta --version`, `zenta -v`)
- Comprehensive test suite with unit tests
- GitHub Actions CI/CD pipeline
- Cross-platform release automation
- Makefile for build automation
- Linting configuration with golangci-lint
- Contributing guidelines
- Production-ready build system

### Changed
- Improved error handling and user feedback
- Enhanced help system with detailed examples

## [0.1.0] - 2025-06-26

### Added
- Core mindfulness CLI tool
- `zenta now` command with breathing exercise and mindfulness quotes
  - 4-4-4-4 box breathing pattern (3 cycles)
  - 25 built-in mindfulness quotes from Zen, Stoic, and mindfulness traditions
- `zenta log` command for tracking distractions, reflections, and insights
  - Support for different log types: distraction (default), reflection, insight
  - Automatic timestamping with UUID generation
  - Local JSON storage in `~/.zenta/logs.json`
- `zenta stats` command for analytics
  - Period-based filtering (today, week, month, all)
  - ASCII bar charts for hourly activity distribution
  - Entry type breakdown and counts
- `zenta help` command with comprehensive usage information
- Local data storage system using JSON files
- Configuration system with sensible defaults
- Unix philosophy compliance
  - Single-purpose focus on mindfulness
  - Composable with other Unix tools
  - Text-based terminal-friendly output
  - Graceful error handling

### Technical Features
- Go-based implementation for cross-platform compatibility
- Modular architecture with clean separation of concerns
- No external dependencies except UUID generation
- Offline-first design with no internet requirements
- Standard Unix conventions for configuration and data storage

[Unreleased]: https://github.com/e6a5/zenta/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/e6a5/zenta/releases/tag/v0.1.0 