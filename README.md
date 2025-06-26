# zenta

> Mindfulness for terminal users

**zenta** is a lightweight command-line tool that brings mindfulness into your developer workflow. Take guided breathing moments, track your awareness, and cultivate presence—all from your terminal.

Inspired by Zen philosophy and the Unix principle of "do one thing, and do it well."

## Why zenta?

**For developers who want to:**
- 🧘 **Take mindful breaks** during intense coding sessions
- 📊 **Track distractions** and build self-awareness
- 🌿 **Reduce context switching** with terminal-native mindfulness
- ⚡ **Stay present** without leaving their development environment

**The result:** More focused work, better awareness, and a calmer mind.

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

## Quick Start

```bash
# Take a mindful moment with guided breathing
zenta now

# Log when you notice distraction
zenta log "Scrolled social media instead of coding"

# Track different types of awareness
zenta log -t reflection "Noticed I was feeling anxious"
zenta log -t insight "Morning hours are my most focused time"

# View your mindfulness patterns
zenta stats week
```

## Commands

### 🧘 `zenta now`
**Take a mindful moment with guided breathing exercise + wisdom quote**

```bash
$ zenta now
🧘 Let's take a moment to breathe together...
   Follow the rhythm: Inhale → Hold → Exhale → Hold

   Cycle 1/3:
   🌬️  Inhale ●●●●
   ⏸️  Hold   ○○○○
   💨 Exhale ●●●●
   ⏸️  Hold   ○○○○
   
   [... 2 more cycles ...]

✨ Beautiful. Now, here's a moment of wisdom:

🌊 You have power over your mind—not outside events. 
   Realize this, and you will find strength. - Marcus Aurelius
```

### 📝 `zenta log`
**Track your awareness moments**

```bash
# Log distraction (default)
zenta log "Got distracted by notifications"

# Log reflection on your mental state
zenta log -t reflection "Noticed I was feeling overwhelmed"

# Log insights about your productivity
zenta log -t insight "I focus better with music"
```

**Types:** `distraction` (default), `reflection`, `insight`

### 📊 `zenta stats`
**Understand your mindfulness patterns**

```bash
$ zenta stats
📊 Zenta Statistics (week)
────────────────────────────────────────

📅 Time Range: Jun 24 to Jun 26, 2025
📝 Total Entries: 8

Entry Types:
  🔴 Distractions: 5
  🤔 Reflections:  2
  💡 Insights:     1
  ⏰ Sessions:     0

Hourly Activity:
  09:00 ██████████ 2
  14:00 ████████████████████ 4
  16:00 █████ 2
```

**Periods:** `today`, `week` (default), `month`, `all`

## Key Features

✅ **Guided breathing exercises** - 4-4-4-4 box breathing with visual cues  
✅ **Mindfulness quotes** - 25+ wisdom quotes from Zen, Stoic, and mindfulness traditions  
✅ **Awareness tracking** - Log distractions, reflections, and insights  
✅ **Pattern analysis** - ASCII charts showing your mindfulness trends  
✅ **Privacy-first** - All data stays local, no internet required  
✅ **Unix philosophy** - Simple, composable, does one thing well  

## Unix Philosophy in Action

zenta works great with other terminal tools:

```bash
# Email yourself weekly mindfulness reports
zenta stats week | mail -s "Weekly Mindfulness" you@example.com

# Count your distractions
zenta stats | grep "Distractions:" | awk '{print $2}'

# Archive your logs
zenta stats all > mindfulness-report-2025.txt
```

## Privacy & Data

- **Local only** - All data stored in `~/.zenta/` on your machine
- **No tracking** - No analytics, telemetry, or data collection
- **No internet** - Works completely offline
- **Your data** - Export, modify, or delete anytime

## What's Next?

See our [**Roadmap**](ROADMAP.md) for upcoming features like:
- Focus timers (`zenta start`)
- Mindful breaks (`zenta break`) 
- Advanced analytics and insights

## Get Involved

- 🐛 **Found a bug?** [Report it](https://github.com/e6a5/zenta/issues)
- 💡 **Have an idea?** [Share it](https://github.com/e6a5/zenta/discussions)
- 🔧 **Want to contribute?** See [CONTRIBUTING.md](CONTRIBUTING.md)
- 📖 **Need help?** Check out the [documentation](https://github.com/e6a5/zenta/wiki)

## License

MIT License - see [LICENSE](LICENSE) file for details.

---

> *"The best way to take care of the future is to take care of the present moment."*

**zenta** - mindfulness for terminal users 🧘‍♂️ 