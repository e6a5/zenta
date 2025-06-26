# Product Requirements Document: zenta
## 1. Overview
Zenta is a lightweight command-line tool designed for developers and terminal users who want to cultivate mindfulness, reduce distractions, and maintain awareness throughout their workday. Inspired by Zen philosophy and Stoic practice, zenta brings calm, clarity, and presence into the world of deep work.

## 2. Development Philosophy & Rules

**Core Principle: "Do one thing, and do it well"**

Zenta exists solely to help terminal users cultivate mindfulness and awareness. Every feature must serve this single purpose.

### **Unix Philosophy Guidelines:**

**1. Single Purpose Focus**
- Zenta is a mindfulness tool, not a task manager, note-taker, or productivity suite
- Each command should have one clear, well-defined function
- Resist feature creep that dilutes the core purpose

**2. Composability**
- Commands should work well with other Unix tools via pipes and redirection
- `zenta export` should produce clean output suitable for `awk`, `grep`, `sort`
- Logs should be human-readable and processable by standard Unix utilities

**3. Simplicity Over Features**
- Prefer simple, obvious solutions over complex, clever ones
- When choosing between two approaches, pick the one with fewer dependencies
- Configuration should be minimal and have sensible defaults

**4. Text-Based Interface**
- All output should be plain text, suitable for terminal environments
- No GUI components, web interfaces, or rich media
- ASCII-safe output that works in any terminal

**5. Fail Gracefully**
- Handle errors quietly and provide useful error messages
- Degrade gracefully when optional features (sound, notifications) aren't available
- Never crash or produce cryptic error messages

**6. Respect User Environment**
- Don't pollute the user's environment with unnecessary files or processes
- Use standard Unix conventions for configuration and data storage
- Work well with existing terminal workflows and tools

### **Development Rules:**

**Feature Addition Criteria:**
- Does this feature directly support mindfulness/awareness?
- Can this be accomplished better by composing with existing Unix tools?
- Will this add complexity that outweighs the benefit?
- Does this maintain the "lightweight" nature of the tool?

**Code Quality Standards:**
- Functions should do one thing well
- Minimal external dependencies
- Clear, self-documenting code over complex optimizations
- Comprehensive error handling without verbosity

**Anti-Patterns to Avoid:**
- ❌ Adding task management features ("I'll just add a simple todo list...")
- ❌ Complex configuration systems with dozens of options
- ❌ Web dashboards or GUI components
- ❌ Cloud sync, accounts, or external service dependencies
- ❌ AI features that require large models or internet connectivity
- ❌ Notification systems that are intrusive or attention-grabbing

**Decision Framework:**
When considering any new feature or change, ask:
1. Does this align with helping users be more mindful?
2. Can users achieve this by combining zenta with other Unix tools?
3. Will this make zenta harder to understand or use?
4. Does this respect the "lightweight" promise?

If any answer is "no," reconsider the change.

## 3. Goals
Help users return to the present moment while working in terminal environments.

Reduce unconscious distractions and multitasking habits.

Track self-awareness by logging "moments of drifting."

Provide a simple system for mindful focus cycles (e.g., deep work sessions).

Offer gentle prompts and reflections without intruding or judging.

## 4. Target Users
Developers, sysadmins, or knowledge workers who spend most of their day in the terminal.

Users interested in mindfulness, Zen, Stoicism, or productivity.

People who struggle with digital distractions and want more clarity and self-awareness.

## 5. Core Features
✅ **zenta now**
Display a short mindfulness message or quote.

Designed to be triggered manually (e.g., when taking a breath or resetting focus).

Pull quotes randomly from a built-in list or custom user-defined list.

✅ **zenta log "<reason>"**
Log a moment of distraction or reflection.

Automatically timestamps each entry with session correlation.

Stores logs in a local JSON file (~/.zenta/logs.json).

Supports different log types: distraction, reflection, insight.

✅ **zenta stats [period]**
Show analytics from logs:

Number of distractions per hour/day/week.

ASCII-based charts (text histogram) showing patterns over time.

Session completion rates and focus duration trends.

Optional period filters: today, week, month, all.

Reflection patterns (e.g., time blocks with most drift).

✅ **zenta start [duration]**
Start a "deep work" session (default 45 minutes, configurable).

Visual timer in terminal with progress indicator.

Handles interruptions gracefully (pause/resume capability).

Optional bell sound or message at the end (configurable).

Records session start/end for analytics.

Compatible with terminal multiplexers (tmux, screen).

🔄 **zenta break [duration]**
Initiate a short mindful break (default 10 minutes, configurable).

Gentle prompts for breathing exercises or stretching.

## 6. Additional Commands

**zenta help**
Show available commands, usage examples, and tips.

Context-sensitive help for specific commands.

**zenta config [key] [value]**
Configure timer durations, custom quotes, sound preferences.

Show current configuration when called without parameters.

**zenta history [filter]**
View recent logs with optional filtering by type or date range.

Support for grep-like pattern matching in log messages.

**zenta export [format]**
Export logs in JSON, CSV, or plain text format for external analysis.

Support for date range filtering in exports.

## 7. Non-Goals
Not a full task management system.

Not an AI therapist or chat-based mindfulness app.

No internet dependency — runs completely offline.

No cloud storage or user accounts required.

## 8. CLI Design
```bash
zenta now                    → show a mindfulness quote
zenta log "reason"          → log a moment of drifting or reflection  
zenta log -t insight "text" → log with specific type
zenta stats                 → view logs and distraction analytics
zenta stats week            → view weekly analytics
zenta start                 → start a 45-min focus timer
zenta start 60              → start a 60-min focus timer
zenta break                 → initiate a 10-min mindful break
zenta help                  → show available commands
zenta config                → show current configuration
zenta config timer 50       → set default timer to 50 minutes
zenta history               → view recent logs
zenta export json           → export logs as JSON
```

## 9. Data Storage & Configuration

**Directory Structure:**
```
~/.zenta/
├── logs.json       # Activity and session logs
├── config.json     # User configuration
└── quotes.txt      # Custom user quotes (optional)
```

**Enhanced Log Format:**
```json
[
  {
    "id": "uuid-v4",
    "timestamp": "2025-06-26T14:12:00+07:00",
    "type": "distraction|reflection|insight|session_start|session_end",
    "note": "Scrolled Reddit instead of debugging",
    "session_id": "uuid-v4",
    "duration": 1800
  }
]
```

**Configuration Format:**
```json
{
  "timer_duration": 45,
  "break_duration": 10,
  "sound_enabled": true,
  "notification_method": "bell|message|silent",
  "timezone": "auto",
  "custom_quotes_enabled": false,
  "stats_default_period": "week"
}
```

## 10. Technical Requirements
Written in Go (cross-platform and fast).

Output is terminal-friendly (minimal colors, ASCII-safe).

Single binary install (no dependencies).

Compatible with Unix shells (bash, zsh, fish).

Graceful handling of system sleep/wake during timers.

Support for terminal multiplexers (tmux, screen).

Cross-platform sound notification support.

Robust error handling for file permissions and storage issues.

## 11. User Experience

**First-Time Setup:**
- Automatic creation of ~/.zenta/ directory and default config
- Welcome message with basic usage tips
- Optional interactive setup for preferences

**Timer Behavior:**
- Visual progress indicator with time remaining
- Pause/resume capability (Ctrl+Z friendly)
- Graceful handling of terminal closure/reopening
- Background process management

**Error Handling:**
- Clear error messages for common issues
- Fallback behaviors when ~/.zenta/ isn't writable
- Validation for command arguments and configuration values

## 12. Stretch Goals / Future Features
🔔 Desktop notifications integration (notify-send, osascript).

⛩️ Custom themes or "modes" (Zen, Stoic, Tao) with different quote sets.

☁️ Optional sync across machines via dotfile repo or Git.

🧠 Local-only reflection pattern analysis and suggestions.

📈 Weekly summary generation with insights and trends.

🔌 Plugin system for custom quote sources or integrations.

## 13. Success Metrics

**Quantitative:**
- User retention: >50% of users active after 30 days
- Daily usage: Average of 3+ commands per active user per day
- Session completion: >70% of started focus sessions completed
- Log engagement: Users log distractions/reflections at least 2x/day

**Qualitative Indicators:**
- Measurable reduction in distraction logs over 4-week periods
- Increased session completion rates over time
- User reports of improved focus awareness (via optional feedback)

**Measurement Methods:**
- Anonymous usage analytics (local only, no telemetry)
- Optional user surveys via CLI prompts (quarterly)
- Log pattern analysis showing behavioral improvements

## 14. Development Phases

**Phase 1 (MVP):**
- Core commands: now, log, stats, help
- Basic configuration system
- Local JSON storage

**Phase 2 (Enhanced):**
- Timer functionality: start, break
- Advanced stats with ASCII charts
- History and export commands

**Phase 3 (Polish):**
- Sound notifications
- Terminal multiplexer optimization
- Custom themes and quote management