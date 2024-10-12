# clip

`clip` is a simple command-line utility written in Go that copies the content of a text file to the clipboard. It works across different platforms like macOS, Linux, and Windows by using system commands for clipboard access.

## Features

- Copy content of a specified text file to the system clipboard.
- Cross-platform support (macOS, Linux, Windows).

## Installation

### 1. Clone the repository:
```bash
git clone https://github.com/kuroko1t/clip.git
```

### 2. Navigate to the project directory:
```bash
cd clip
```

### 3. Build the binary:
```bash
go build
```

### 4. Install clipboard dependencies (Linux only):
If you are using Linux, you need to install `xclip` for clipboard functionality. Use the following commands based on your distribution:

#### Ubuntu/Debian:
```bash
sudo apt update
sudo apt install xclip
```

#### Fedora:
```bash
sudo dnf install xclip
```

#### Arch Linux:
```bash
sudo pacman -S xclip
```

### For macOS and Windows:
- **macOS**: `pbcopy` is used, which is pre-installed on macOS, so no extra installation is needed.
- **Windows**: The built-in `clip` command is used, which is also pre-installed on Windows.

## Usage

To use the utility, simply run the following command in your terminal:

```bash
./clip example.txt
```

This will copy the content of `example.txt` to your clipboard.

### Supported Platforms

- **macOS**: Uses `pbcopy` command.
- **Linux**: Uses `xclip` command (requires installation, as explained above).
- **Windows**: Uses the built-in `clip` command.

## License

This project is licensed under the MIT License.
