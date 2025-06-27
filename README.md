# 🧘 zenta

> **Mindfulness that fits your coding flow**

When your mind wanders while coding, simply return to breath. No tracking, no metrics, no disruption—just pure awareness in your terminal.

**The noticing itself is the practice.** ✨

---

## ⚡ Quick Start

```bash
# Install (one-liner)
curl -fsSL https://raw.githubusercontent.com/e6a5/zenta/main/install.sh | bash

# Set up zen aliases
echo "alias breath='zenta now --quick'" >> ~/.zshrc
echo "alias breathe='zenta now'" >> ~/.zshrc
echo "alias reflect='zenta reflect'" >> ~/.zshrc
source ~/.zshrc

# Use instantly
breath    # When mind wanders → return to breath (1 breath cycle)
breathe   # Need deeper centering → longer session (3 breath cycles)  
reflect   # Evening → gentle day review
```

---

## 🌸 Why zenta?

### **Fits Your Real Workflow**
1. **Notice**: "I'm stuck in planning thoughts"
2. **Return**: `breath` 
3. **Continue**: Back to coding, more present

### **True Zen Approach**
- ✅ **No tracking** → Awareness isn't data to optimize
- ✅ **No analytics** → The practice is the goal
- ✅ **No disruption** → Stays in your terminal
- ✅ **Just breathing** → Pure mindfulness

### **Built for Developers**
- Terminal-native (fits your workflow)
- Single binary (no dependencies)
- Private (no data collection)
- Instant (`breath` is one keystroke away)

---

## 🌊 Beautiful Breathing

```bash
$ breath

       Let's breathe 🌸

    🌬️ Breathe in gently, let your body expand...
                                                
                      ○                         
                    ○ · ○                       
                     ···                        
              ○ · · ·   · · · ○                
                     ···                        
                    ○ · ○                       
                      ·                         

       Carry this calm with you throughout your day 🙏
```

*Pure visual breathing guidance with gentle animations*

---

## 💡 Commands

### **Essential Commands (with aliases)**
| Command | Cycles | What It Does | Perfect For |
|---------|--------|--------------|-------------|
| `breath` | 1 cycle | Quick breathing + wisdom quote | Mind wandering moments |
| `breathe` | 3 cycles | Standard breathing + wisdom quote | Before difficult tasks |
| `reflect` | - | Gentle evening review | End of day contemplation |

### **Full Commands**
| Command | Cycles | What It Does |
|---------|--------|--------------|
| `zenta now` | 3 cycles | Standard breathing + wisdom quote |
| `zenta now --quick` | 1 cycle | Quick breathing + wisdom quote |
| `zenta now --extended` | 5 cycles | Extended breathing + wisdom quote |
| `zenta now --silent` | 3 cycles | Breathing only, no quote |
| `zenta now --simple` | 3 cycles | Simple line animation (terminal compatibility) |

**Mix options:** `zenta now --quick --silent` (1 cycle, no quote)

---

## 🔧 Terminal Compatibility

**Beautiful circles vs simple lines:**
- **Most terminals**: Gorgeous expanding/contracting breathing circles
- **macOS Terminal.app**: Auto-detects and uses simple line animation  
- **tmux/screen**: Works great with complex animations

**Force simple mode:** Add `--simple` to any command
```bash
breath --simple     # Force simple animation
zenta now --simple  # Works with any options
```

**Why?** Only macOS Terminal.app has ANSI escape sequence quirks. zenta auto-detects and adapts for the best experience.

---

## 🎯 Real vs Fake Mindfulness

**✅ Real mindfulness (zenta's way):**
- Notice when mind wanders
- Return to breath instantly  
- Continue work with awareness
- No measurement needed

**❌ Fake mindfulness:**
- Tracking meditation streaks
- Optimizing awareness metrics
- Quantifying inner peace
- Making mindfulness productive

---

## 🌿 Philosophy

> *"The quieter you become, the more you are able to hear."*

**zenta believes:**
- Mindfulness tools should disappear into practice
- The noticing itself is enlightenment
- Developers need presence, not productivity hacks
- True zen has no metrics

---

## 🚀 Installation

### **Option 1: Pre-built Binaries (Recommended)**

**macOS/Linux:**
```bash
# Download and install the latest binary for your platform
curl -s https://api.github.com/repos/e6a5/zenta/releases/latest \
| grep "browser_download_url.*$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m | sed 's/x86_64/amd64/')" \
| cut -d '"' -f 4 \
| xargs curl -L -o zenta.tar.gz \
&& tar -xzf zenta.tar.gz \
&& sudo mv zenta-* /usr/local/bin/zenta \
&& rm zenta.tar.gz

# Or download manually from GitHub Releases
```

**Windows:**
Download the latest `.zip` file from [GitHub Releases](https://github.com/e6a5/zenta/releases), extract it, and add the executable to your PATH.

**Manual Download:** [GitHub Releases](https://github.com/e6a5/zenta/releases) - Choose your platform

### **Option 2: Build from Source**

*Requires Go 1.23+ installed*

```bash
git clone https://github.com/e6a5/zenta.git && cd zenta && make install-system
```

**Install Go first:** [https://golang.org/dl/](https://golang.org/dl/)

*Supports: Linux, macOS, Windows, FreeBSD (all architectures)*

---

## 🧘 Join the Practice

- 🐛 **Bug reports** → [Issues](https://github.com/e6a5/zenta/issues)
- 💭 **Mindful discussions** → [Discussions](https://github.com/e6a5/zenta/discussions)  
- 🔧 **Contributions** → [CONTRIBUTING.md](CONTRIBUTING.md)

---

## 📄 License

MIT License - [LICENSE](LICENSE)

---

> *"The best way to take care of the future is to take care of the present moment."* — Thich Nhat Hanh

**Start your mindful coding practice today:**

```bash
breath  # Just try it 🙏
```
