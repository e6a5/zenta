# zenta

> Mindfulness for terminal users

**zenta** is a lightweight command-line tool designed for developers and terminal users who want to cultivate mindfulness, reduce distractions, and maintain awareness throughout their workday. Inspired by Zen philosophy and Stoic practice, zenta brings calm, clarity, and presence into the world of deep work.

## Philosophy

Following the Unix philosophy: **"Do one thing, and do it well"**

zenta exists solely to help terminal users cultivate mindfulness and awareness. Every feature serves this single purpose.

## Features

### Phase 1 (MVP) - Available Now ✅

- **`zenta now`** - Display mindfulness quotes to help you return to the present moment
- **`zenta log "<reason>"`** - Log moments of distraction or reflection with timestamps
- **`zenta stats [period]`** - View analytics from your logs with ASCII charts
- **`zenta help`** - Context-sensitive help and usage examples

## Installation

### Quick Install

```bash
# Using make (recommended)
git clone https://github.com/e6a5/zenta.git
cd zenta
make install-system
```

### From Source

```bash
git clone https://github.com/e6a5/zenta.git
cd zenta
make build
sudo mv zenta /usr/local/bin/  # or any directory in your PATH
```

### Pre-built Binaries

Download the latest release for your platform from [GitHub Releases](https://github.com/e6a5/zenta/releases).

Available platforms:
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)
- FreeBSD (amd64)

### Requirements

- Go 1.20 or later (for building from source)

## Usage

### Quick Start

```bash
# Show a mindfulness quote
zenta now

# Log a moment of distraction
zenta log "Scrolled social media instead of coding"

# View your statistics
zenta stats

# See all available commands
zenta help
```

### Commands

#### `zenta now`
Get a random mindfulness quote to help center yourself.

```bash
$ zenta now
🌊 You have power over your mind—not outside events. Realize this, and you will find strength. - Marcus Aurelius
```

#### `zenta log "<reason>"`
Log a moment of distraction, reflection, or insight.

```bash
$ zenta log "Checked email instead of finishing the code review"
✓ Logged: Checked email instead of finishing the code review
  Take a moment to breathe and return to the present.
```

#### `zenta stats [period]`
View analytics about your logged entries.

```bash
$ zenta stats week
📊 Zenta Statistics (week)
────────────────────────────────────────

📅 Time Range: Jun 24 to Jun 26, 2025
📝 Total Entries: 5

Entry Types:
  🔴 Distractions: 5
  🤔 Reflections:  0
  💡 Insights:     0
  ⏰ Sessions:     0

Hourly Activity:
  09:00 ██████████ 2
  14:00 ████████████████████ 4
  16:00 █████ 1
```

**Available periods:**
- `today` - Today's entries
- `week` - This week's entries (default)
- `month` - This month's entries  
- `all` - All entries

## Data Storage

zenta stores all data locally in `~/.zenta/`:

```
~/.zenta/
├── logs.json       # Your logged entries
├── config.json     # Configuration (created when needed)
└── quotes.txt      # Custom quotes (future feature)
```

### Log Format

```json
[
  {
    "id": "uuid-v4",
    "timestamp": "2025-06-26T14:12:00+07:00",
    "type": "distraction",
    "note": "Scrolled Reddit instead of debugging"
  }
]
```

## Philosophy & Design

### Unix Philosophy Guidelines

1. **Single Purpose Focus** - zenta is a mindfulness tool, not a task manager
2. **Composability** - Works well with other Unix tools via pipes and redirection
3. **Simplicity Over Features** - Minimal dependencies, sensible defaults
4. **Text-Based Interface** - ASCII-safe output that works in any terminal
5. **Fail Gracefully** - Clear error messages, degrade gracefully
6. **Respect User Environment** - Standard Unix conventions, minimal footprint

### Composability Examples

```bash
# Export logs for external analysis
zenta stats | mail -s "Weekly Focus Report" user@example.com

# Count distractions
zenta stats | grep "Distractions:" | awk '{print $2}'

# View logs in your preferred pager
zenta stats all | less
```

## Roadmap

### Phase 2 (Enhanced)
- **Timer functionality** - `zenta start`, `zenta break` 
- **Advanced stats** - More detailed analytics and patterns
- **History and export** - `zenta history`, `zenta export`

### Phase 3 (Polish)
- **Sound notifications** - Optional audio cues
- **Terminal multiplexer optimization** - Better tmux/screen support
- **Custom themes** - Different quote collections (Zen, Stoic, etc.)

## Contributing

zenta follows strict development principles:

**Before adding any feature, ask:**
1. Does this align with helping users be more mindful?
2. Can users achieve this by combining zenta with other Unix tools?
3. Will this make zenta harder to understand or use?
4. Does this respect the "lightweight" promise?

If any answer is "no," reconsider the change.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Inspiration

> "The best way to take care of the future is to take care of the present moment."

zenta draws inspiration from:
- Zen Buddhism and mindfulness practices
- Stoic philosophy  
- The Unix philosophy of simple, composable tools
- The growing need for digital wellness in developer workflows 