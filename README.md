# 🛡️ Zsh History Mask Tool

A powerful utility to **scan, detect, and remove sensitive credentials** from your **Zsh** command history. Keep your terminal history clean and your secrets safe!

## 🚀 Features

- **Automatic Detection**: Identifies sensitive data (e.g., passwords, API keys, tokens) in your Zsh history.
- **History Sanitization**: Masks or removes commands with detected credentials.
- **Backup Support**: Automatically backs up your original history before modification.
- **Efficient and Fast**: Scans large history files quickly and updates securely.

## 📦 Installation

Clone the repository and navigate to the directory:

```bash
brew tap danievanzyl/homebrew-zshhistorymasker
brew install zshhistorymasker
```

## 🛠️ Usage

Run the tool to scan and sanitize your Zsh history:

```bash
zshhistorymasker
fc -R ~/.zsh_history
```

### Options:

| Flag              | Description                              |
| ----------------- | ---------------------------------------- |
| `-p`, `--pattern` | Use custom patterns for detection (todo) |
| `-v`, `--version` | Version information                      |
| `-h`, `--help`    | Show help information                    |

Example:

```bash
zsh-history-mask -p ".*password=.*"
```

## 📊 How It Works

1. **Scan**: Reads your Zsh history file (`~/.zsh_history`).
2. **Detect**: Identifies patterns matching sensitive information.
3. **Sanitize**: Replaces or removes the detected commands.
4. **Update**: Safely updates the history file.

## ✅ To-Do

- [ ] Enhance detection with regex patterns
- [ ] Add support for other shell histories (e.g., Bash)
- [ ] Implement dry-run mode for safe previews

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create a feature branch
3. Submit a pull request

## 📄 License

This project is licensed under the **Unlicense**. See the [LICENSE](LICENSE) file for details.

---

⭐️ **If you find this project helpful, give it a star!**
