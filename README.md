# 🧘 zenta

> **Mindfulness that fits your workflow**

**zenta** brings calm to your terminal. Take guided breathing breaks, track awareness moments, and cultivate presence—all without leaving your development flow.

*Built for developers who code mindfully.* ✨

---

## ⚡ Try it now

```bash
# Take a quick 1-minute breathing break
zenta now --quick

# Standard 3-minute session with wisdom quote  
zenta now

# Extended 5-minute deep session
zenta now --extended
```

**🎯 Perfect for:** Code reviews, debugging sessions, before standups, or anytime you need to reset.

---

## 🌸 What makes zenta special?

### **Instant Calm**
```bash
$ zenta now --quick

       Let's breathe 🌸

    🌬️ Breathe in gently, let your body expand...
                                                
                      ○                         
                    ○ · ○                       
                     ···                        
              ○ · · ·   · · · ○                
                     ···                        
                    ○ · ○                       
                      ·                         
                      ○                         


    🌊 T h e   q u i e t e r   y o u   b e c o m e, 
       t h e   m o r e   y o u   a r e   a b l e   t o   h e a r.
```

### **Organic & Natural**
- **`--quick`** → 1-minute breathing break for busy moments
- **`--extended`** → 5-minute deep session when you have time  
- **`--silent`** → Pure breathing, no quotes for focused sessions
- **Breathing circle** → Organic visualization that expands and contracts like real breathing
- **Immediate flow** → Starts instantly, no interruptions, pure focus
- **Typing wisdom** → Quotes reveal word-by-word after breathing for contemplative moments

### **Built for Developers**
```bash
# Track patterns that matter
zenta log "Got distracted by Slack notifications"
zenta log -t insight "Morning standup made me anxious"

# Understand your mind
zenta stats week
📊 Zenta Statistics (week)
📅 Time Range: Jun 24 to Jun 26, 2025
Entry Types:
  🔴 Distractions: 5  ← Most happen at 2pm
  🤔 Reflections:  2
  💡 Insights:     1
```

---

## 🚀 Installation

**macOS/Linux (quick):**
```bash
git clone https://github.com/e6a5/zenta.git && cd zenta && make install-system
```

**Or download pre-built binaries:** [GitHub Releases](https://github.com/e6a5/zenta/releases)

*Supports: Linux, macOS, Windows, FreeBSD (all architectures)*

---

## 💡 Core Commands

| Command | What it does | When to use |
|---------|-------------|-------------|
| `zenta now` | Organic breathing circle with 3 continuous breaths | Regular mindful breaks |
| `zenta now --quick` | Quick 1-breath session with typing wisdom | Between meetings, quick reset |
| `zenta now --extended` | Deep 5-breath session for longer calm | Start of day, end of sprint |
| `zenta now --silent` | Pure breathing circle, no quote | Pure focus, no distractions |
| `zenta log "reason"` | Track awareness moments | Notice patterns, build habits |
| `zenta stats` | View your mindfulness patterns | Weekly reflection, insights |

---

## 🎯 Why developers love zenta

✅ **Terminal-native** - No context switching, fits your workflow  
✅ **Lightweight** - Single binary, no dependencies, <2MB  
✅ **Private** - All data local, no internet, no tracking  
✅ **Unix-friendly** - Pipes, scripts, automation-ready  
✅ **Instant** - From stress to calm in under 60 seconds  
✅ **Progressive** - Build awareness over time  

---

## 🧠 Smart Awareness Tracking

Track the moments that matter and build better habits:

```bash
# Common developer distractions
zenta log "Opened Twitter instead of documentation"
zenta log "Started debugging without reading error message"
zenta log "Procrastinated on code review"

# Reflections and insights  
zenta log -t reflection "Feeling overwhelmed by technical debt"
zenta log -t insight "Pair programming reduces my anxiety"

# See patterns emerge
zenta stats month
```

**Result:** Better self-awareness, fewer distractions, calmer coding.

---

## 🛠 Unix Philosophy in Action

zenta plays well with your existing tools:

```bash
# Daily mindfulness report
zenta stats today | mail -s "Daily Mindfulness" you@example.com

# Mindful git workflow
alias mindful-commit="zenta now --quick && git commit"

# Weekly team insights
zenta stats week > weekly-mindfulness.md

# Count distractions
zenta stats | grep "Distractions:" | awk '{print $2}'
```

---

## 🌿 What's Next?

**Immediate roadmap:**
- [ ] Focus timers (`zenta start 25` for Pomodoro)
- [ ] Mindful breaks (`zenta break` between deep work)
- [ ] Custom breathing patterns (4-7-8, triangle breathing)
- [ ] Integration hooks for IDEs and editors

**Vision:** The most thoughtful developer productivity tool ever built.

---

## 🤝 Join the Community

- 🐛 **Found a bug?** → [Report it](https://github.com/e6a5/zenta/issues)
- 💡 **Have an idea?** → [Share it](https://github.com/e6a5/zenta/discussions)  
- 🔧 **Want to contribute?** → See [CONTRIBUTING.md](CONTRIBUTING.md)
- 📖 **Need help?** → Check the [documentation](https://github.com/e6a5/zenta/wiki)

---

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

---

> *"The best way to take care of the future is to take care of the present moment."*

**Start your mindful coding journey today.** 🧘‍♂️

```bash
zenta now --quick  # Just try it
``` 