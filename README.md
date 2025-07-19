# 📋 Tedo - Terminal Todo

A beautiful, interactive command-line todo application built with Go and Bubble Tea. Manage your daily tasks, upcoming todos, and general notes with vim-style keybindings and an intuitive calendar interface.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macOS%20%7C%20windows-lightgrey.svg)

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

### 🎨 **Beautiful Interface**
- Clean, modern terminal UI
- Color-coded todo states
- Pagination for large todo lists
- Real-time input validation

## 🚀 Quick Start

### Installation

#### Option 1: Install from source
```bash
# Clone the repository
git clone https://github.com/WasathTheekshana/Tedo.git
cd Tedo

# Build and install
go build -o tedo cmd/tedo/main.go
sudo mv tedo /usr/local/bin/

# Or just run directly
go run cmd/tedo/main.go
```

#### Option 2: Go install (if published)
```bash
go install github.com/WasathTheekshana/Tedo/cmd/tedo@latest
```

### First Run
```bash
tedo
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
Tedo/
├── cmd/tedo/           # Application entry point
│   └── main.go
├── internal/           # Private application code
│   ├── models/         # Data structures
│   ├── storage/        # JSON persistence layer
│   └── ui/             # Terminal user interface
│       ├── app.go      # Main application logic
│       ├── calendar.go # Calendar component
│       ├── keys.go     # Keyboard handling
│       ├── render.go   # UI rendering
│       ├── styles.go   # Visual styling
│       ├── input.go    # Input handling
│       ├── validation.go # Input validation
│       └── errors.go   # Error management
├── data/               # JSON data files (auto-created)
├── go.mod
├── go.sum
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
