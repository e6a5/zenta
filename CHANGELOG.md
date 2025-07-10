# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2025-07-15

### Added

- **`anchor` command**: A new interactive breathing mode where you guide the session with keyboard input (`SPACE` to switch phases, `q` to quit). This provides a user-led alternative to the fixed-cycle `now` command.

### Changed

- **Command Naming**: The interactive mode was renamed from `manual` to `breathe`, then `pacer`, and finally `anchor` to better align with mindfulness principles and avoid ambiguity.
- **UI Polish**: Improved the exit sequence for the `anchor` command to provide a smoother, more polished user experience.
- **Documentation**: Updated `README.md` to include the new `anchor` command, its alias, and clear usage instructions.

### Fixed

- Resolved a layout issue where the closing quote would appear too abruptly after quitting the `anchor` session.

## [1.0.1] - 2025-07-08

### Changed

- Improved the exhale animation for a smoother visual experience.

## [1.0.0] - 2025-07-03

### Added

- `reflect` command was transformed into a paced, guided mindfulness session.
- Nix flake support for reproducible development environments.

### Changed

- The installation script now creates `/usr/local/bin` if it doesn't exist.

## [0.3.5] - 2025-06-28

### Changed

- Removed outdated animation restrictions for `tmux` and `screen` sessions.

## [0.3.4] - 2025-06-28

### Changed

- The terminal cursor is now hidden during breathing animations for a cleaner look.

## [0.3.3] - 2025-06-28

### Added

- `--complex` flag to force the complex circle animation on any terminal.
- `CODE_OF_CONDUCT.md` and other project documentation.

## [0.3.2] - 2025-06-27

### Changed

- Improved the installation script with better error handling.
- Help messages now correctly use the program name.
- Fixed `gofmt` formatting issues.

## [0.3.1] - 2025-06-27

### Fixed

- Improved terminal compatibility for the breathing animations, especially in simple mode.

## [0.3.0] - 2025-06-27

### Changed

- Simplified the quote display to better focus on the breathing cycles.
- Updated README with clearer instructions.

## [0.2.0] - 2025-06-27

### Changed

- Refactored the codebase into a clean, package-based architecture.

## [0.0.1] - 2025-06-27

### Added

- Initial release of `zenta`.
- Core breathing functionality with `zenta now`.
- A collection of mindfulness quotes.

[1.1.0]: https://github.com/e6a5/zenta/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/e6a5/zenta/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/e6a5/zenta/compare/v0.3.5...v1.0.0
[0.3.5]: https://github.com/e6a5/zenta/compare/v0.3.4...v0.3.5
[0.3.4]: https://github.com/e6a5/zenta/compare/v0.3.3...v0.3.4
[0.3.3]: https://github.com/e6a5/zenta/compare/v0.3.2...v0.3.3
[0.3.2]: https://github.com/e6a5/zenta/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/e6a5/zenta/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/e6a5/zenta/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/e6a5/zenta/compare/v0.0.1...v0.2.0
[0.0.1]: https://github.com/e6a5/zenta/releases/tag/v0.0.1
