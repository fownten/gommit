# gommit

An interactive CLI tool for creating standardized, stylized git commit messages.

Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss).

## Features

- **Interactive Selection**: Choose between `feat`, `fix`, and `chore`.
- **Stylized UI**: Modern, inline terminal interface that doesn't disrupt your scrollback buffer.
- **Smart Detection**: Automatically detects staged files and includes them in the commit message body for better traceability.
- **Direct Execution**: Safely executes `git commit` after your confirmation.

## Installation

### From Source
Ensure you have Go installed, then clone and build:

```bash
git clone https://github.com/fownten/gommit.git
cd gommit
go build -o gommit main.go
```

## Usage

1. Stage your changes as usual: `git add .`
2. Run the tool: `./gommit`
3. Follow the interactive prompts to select your commit type and enter a description.
4. Confirm the preview to execute the commit.

## Requirements

- [Go](https://go.dev/) 1.24+
- [Git](https://git-scm.com/) installed and available in your PATH.

## License
MIT (or your preferred license)
