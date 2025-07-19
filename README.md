# 📋 Tedo - Terminal Todo

A beautiful, interactive command-line todo application built with Go and Bubble Tea. Manage your daily tasks, upcoming todos, and general notes with vim-style keybindings and an intuitive calendar interface.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macOS%20%7C%20windows-lightgrey.svg)

---

<img width="1609" height="472" alt="image" src="https://github.com/user-attachments/assets/6f15cbf3-a85a-4a6c-95a8-b19a90cbee08" />
<img width="1697" height="720" alt="image" src="https://github.com/user-attachments/assets/d688f4ae-9933-46c5-a507-b17351adc668" />


## ✨ Features

### 🎯 **Smart Todo Organization**
- **Today View**: Focus on today's tasks only
- **Upcoming View**: See all future-dated todos
- **Calendar View**: Monthly calendar with todo counts
- **General View**: Non-dated todos and notes

### ⚡ **Vim-Style Navigation**
- `hjkl` for content navigation
- Arrow keys for menu switching
- Familiar vim operations (`i`, `e`, `d`, `x`)
- Fast keyboard-driven workflow

### 📅 **Interactive Calendar**
- Monthly view with todo indicators
- Jump to any date to view/add todos
- Navigate months with `n`/`p`
- Quick return to today with `t`

### 💾 **Reliable Data Storage**
- JSON file-based persistence
- Automatic data organization by date
- No external database required
- Human-readable data format

### Enhanced Features ✨
- **Smart Input Validation**: Character limits and real-time feedback
- **Auto-clearing Errors**: Error messages disappear after 5 seconds
- **Enhanced Keyboard Shortcuts**: `Ctrl+S` to save, `Ctrl+A` select all
- **Performance Optimized**: Handles 1000+ todos efficiently
- **Character Counters**: Live character count in input forms
- **Version Information**: `tedo -version` for version details
- **Help System**: `tedo -help` for usage information
- Clean, modern terminal UI
- Color-coded todo states
- Pagination for large todo lists
- Real-time input validation

## 🚀 Quick Start

### Installation

## 🚀 Installation

### Method 1: Go Install (Recommended)
```bash
go install github.com/WasathTheekshana/tedo/cmd/tedo@latest
```

### Method 2: Using Installation Script
```bash
curl -fsSL https://raw.githubusercontent.com/WasathTheekshana/tedo/main/install.sh | bash
```

### Method 3: From Source
```bash
git clone https://github.com/WasathTheekshana/tedo.git
cd tedo
go build -o tedo cmd/tedo/main.go
sudo mv tedo /usr/local/bin/
```

### Verify Installation
```bash
tedo -version
```

**Note:** Make sure `$GOPATH/bin` is in your `$PATH`. Add this to your shell profile if needed:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Command Line Options
```bash
tedo                # Start the application
tedo -version       # Show version information  
tedo -help          # Show help message
```

The app will create a `data/` directory in the current folder to store your todos.

## 📖 Usage Guide

### 🔤 **Navigation**
| Key | Action |
|-----|--------|
| `←` `→` | Switch between tabs |
| `j` `k` | Navigate up/down in lists |
| `h` `j` `k` `l` | Navigate calendar dates |
| `1` `2` `3` `4` | Jump to specific views |
| `c` | Quick jump to calendar |
| `q` / `Ctrl+C` | Quit |

### ✏️ **Todo Operations**
| Key | Action |
|-----|--------|
| `i` | Add new todo |
| `e` | Edit selected todo |
| `d` | Delete selected todo |
| `x` | Toggle completion |
| `Enter` | View date (from calendar) |

### 📝 **Input Mode**
| Key | Action |
|-----|--------|
| `Tab` | Switch between title/description |
| `Enter` / `Ctrl+S` | Save todo |
| `Esc` | Cancel |
| `Ctrl+A` | Select all text |
| `Ctrl+C` | Quit application |

**Input Validation:**
- Title: Required, max 100 characters
- Description: Optional, max 500 characters  
- Real-time character counting
- Auto-clearing error messages

### 📅 **Calendar Navigation**
| Key | Action |
|-----|--------|
| `h` `j` `k` `l` | Move between dates |
| `n` `p` | Next/previous month |
| `t` | Jump to today |
| `Enter` | View todos for selected date |
| `i` | Add todo for selected date |

### 📄 **Pagination**
- Automatically enabled for 10+ todos
- `Ctrl+F` / `Ctrl+B` for page navigation
- Seamless navigation between pages

## 🏗️ Project Structure

```
tedo/
├── cmd/tedo/           # Application entry point
│   └── main.go
├── internal/           # Private application code
│   ├── models/         # Data structures
│   ├── storage/        # JSON persistence layer
│   ├── version/        # Version information
│   └── ui/             # Terminal user interface
│       ├── app.go      # Main application logic
│       ├── calendar.go # Calendar component
│       ├── keys.go     # Keyboard handling
│       ├── render.go   # UI rendering
│       ├── styles.go   # Visual styling
│       ├── input.go    # Input handling
│       ├── validation.go # Input validation
│       ├── errors.go   # Error management
│       ├── performance.go # Performance monitoring
│       └── help.go     # Help system
├── install.sh          # Installation script
├── uninstall.sh        # Uninstallation script
├── data/               # JSON data files (auto-created)
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## 🎨 Screenshots

### Today View
```
📋 Tedo Today Upcoming Calendar General

📅 2025-07-19

> ☐ 1. Morning Exercise
    30 minutes of jogging
  ✓ 2. Code Review
    Review PR #123
  ☐ 3. Team Meeting
    Daily standup at 10 AM

j/k: navigate • ←/→: switch tabs • x: toggle • d: delete • e: edit • i: add • q: quit
```

### Calendar View
```
📋 Tedo Today Upcoming [Calendar] General

📅 July 2025

  Su  Mo  Tu  We  Th  Fr  Sa
       1   2   3   4   5
   6   7   8   9  10  11  12
  13  14  15  16  17  18 >19<
  20  21  22• 23  24  25  26
  27  28  29  30  31

📝 2025-07-19 (3 todos)
  ☐ Morning Exercise
  ✓ Code Review
  ☐ Team Meeting

h/j/k/l: navigate dates • n/p: month • t: today • enter: view date • i: add • q: quit
```

## 🗑️ Uninstallation

### Quick Uninstall
```bash
# Remove Tedo binary from Go installation
rm $(go env GOPATH)/bin/tedo

# Or remove from system location
sudo rm /usr/local/bin/tedo
```

### Complete Uninstall Script
```bash
curl -fsSL https://raw.githubusercontent.com/WasathTheekshana/tedo/main/uninstall.sh | bash
```

### Manual Cleanup
```bash
# Remove binary from all possible locations
sudo rm -f /usr/local/bin/tedo /usr/bin/tedo ~/.local/bin/tedo ~/bin/tedo $(go env GOPATH)/bin/tedo

# Optionally remove your todo data
rm -rf ./data
```

## ⚙️ Configuration

### Data Location
By default, todos are stored in `./data/` relative to where you run the command. Files include:
- `general.json` - General todos
- `YYYY-MM-DD.json` - Date-specific todos

### Customization
The app uses a clean, minimal design. Colors and styles can be customized by modifying `internal/ui/styles.go`.

## 🤝 Contributing

We welcome contributions! Here's how to get started:

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes**
4. **Add tests if applicable**
5. **Commit your changes**
   ```bash
   git commit -m 'Add amazing feature'
   ```
6. **Push to the branch**
   ```bash
   git push origin feature/amazing-feature
   ```
7. **Open a Pull Request**

### Development Setup
```bash
# Clone and enter directory
git clone https://github.com/WasathTheekshana/Tedo.git
cd Tedo

# Install dependencies
go mod tidy

# Run tests
go test ./...

# Run the application
go run cmd/tedo/main.go
```

### Code Style
- Follow standard Go formatting (`go fmt`)
- Add comments for exported functions
- Keep functions focused and testable
- Use meaningful variable names

## 🐛 Troubleshooting

### Common Issues

**Q: App doesn't start or shows garbled text**
- Ensure your terminal supports ANSI colors
- Try running with `TERM=xterm-256color tedo`

**Q: Data not persisting**
- Check write permissions in the current directory
- Ensure the `data/` directory is not read-only

**Q: Performance issues with many todos**
- The app uses pagination (10 todos per page) automatically
- Consider archiving completed todos periodically

**Q: Keyboard shortcuts not working**
- Verify your terminal emulator supports the key combinations
- Some terminals may intercept certain key combinations

### Getting Help
- 📋 [Open an issue](https://github.com/WasathTheekshana/Tedo/issues)
- 💬 [Start a discussion](https://github.com/WasathTheekshana/Tedo/discussions)
- 📧 Email: wasaththeekshana@gmail.com

## 📚 Technical Details

### Dependencies
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling and layout

### Minimum Requirements
- Go 1.19 or later
- Terminal with ANSI color support
- 50MB disk space

### Performance
- Handles 1000+ todos efficiently
- Lazy loading for large datasets
- Memory usage typically under 10MB

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Charm](https://charm.sh/) for the amazing Bubble Tea framework
- The Go community for excellent tooling and libraries
- All contributors who help improve this project

## 🔮 Roadmap

- [ ] **Import/Export**: CSV and JSON import/export
- [ ] **Search**: Full-text search across all todos
- [ ] **Categories**: Tag-based organization
- [ ] **Reminders**: Due date notifications
- [ ] **Sync**: Cloud synchronization options
- [ ] **Themes**: Customizable color schemes
- [ ] **Stats**: Productivity analytics

---

**⭐ Star this repo if you find it helpful!**

Made with ❤️ and Go
