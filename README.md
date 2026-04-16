
# Logdy (rk-version)

A personal fork of Logdy that incorporates various improvements based on community feedback and my own development experience.

## Table of Contents
- [Acknowledgments](#acknowledgments)
- [Added Improvements](#added-improvements)
- [Quick Start](#quick-start)
- [Monitoring Multiple Applications](#monitoring-multiple-applications)
- [Installation and PATH Setup](#installation-and-path-setup)
    - [Windows](#windows)
    - [Ubuntu](#ubuntu)
- [Building and Releasing](#building-and-releasing)
- [License](#license)

## Acknowledgments
All credit for the core engine and design goes to the original **[Logdy Team](https://github.com/logdyhq)**. This fork simply builds upon their work to add a few specific tools I needed.

## Added Improvements

These modifications address specific workflow friction and UI layout issues:


*   **Dynamic Application Title**: Set a custom application name in the settings to update the browser tab title and display a dedicated, centered title bar below the highlights.


*   **Custom Highlighting**: Allows you to highlight specific words or phrases with custom colors, making it much easier to visually scan and identify important log entries.

https://github.com/user-attachments/assets/6d2f0b13-dbb9-4f7e-b4ad-cb08619dfe6a

*   **Live Search and Filtering**: Results filter instantly as you type.
    *   **Instant Text and RegEx Search**: Search through raw logs using simple text or powerful Regular Expressions.
        - *Example:* `error`

https://github.com/user-attachments/assets/e37d61c3-4e98-4783-a8dd-56fc9b0bc53f

*   **Structured JSON Queries**: Deeply filter logs by querying specific JSON fields.
        - *Example:* `data.level == "trace"`, `data.method == "PATCH"`, `data.duration == "441"`, `data.domain == "nationalbest-of-breed.com"`

https://github.com/user-attachments/assets/31c716f4-04a7-48ac-a95d-e912c2ecd1b0

*   **Show in Context**: If you find an entry through search but want to see the surrounding logs, you can jump directly into its original context with a single click.

https://github.com/user-attachments/assets/c2bb0d7d-f969-4a06-9aac-2cc11d9b146b

## Quick Start

You can use `logdy-rk` just like the original tool:

### Linux / macOS
```bash
# Pipe any log output
$ tail -f app.log | logdy-rk

# Follow specific files
$ logdy-rk follow system.log
```

### Windows
If you downloaded the Windows release, you can run it from PowerShell or Command Prompt. 

**Note on -t (Fallthrough):** Use the `-t` flag if you want to see the logs in your terminal and the Web UI simultaneously.

## Monitoring Multiple Applications

Logdy helps you monitor multiple services at once using different port configurations.

### 1. Multiple Input Ports (Aggregated View)
You can start one Logdy instance that listens on several ports. Logs from all sources will appear in a single dashboard.

```bash
# Listen on ports 8123, 8124, and 8125
logdy-rk socket 8123 8124 8125
```

### 2. Multiple Web UI Instances (Isolated View)
If you want separate dashboards for different environments (e.g., App A and App B), run multiple processes with different UI ports.

```bash
# Instance 1: UI on port 8080
logdy-rk socket 8123 --port 8080

# Instance 2: UI on port 8081
logdy-rk socket 8124 --port 8081
```

## Installation and PATH Setup

### Windows

To run `logdy-rk` from any folder without typing the full path (`.\`), you must add it to your Windows PATH.

#### Step-by-Step: Adding to Windows PATH
1. **Prepare the folder:**
   - Create a folder like `C:\tools\logdy`.
   - Move the downloaded `logdy-rk.exe` into that folder.
2. **Open Environment Variables:**
   - Press the `Win` key and search for **"Edit the system environment variables"**.
   - Click **Environment Variables** (bottom right).
3. **Edit User Path:**
   - Under **"User variables for [YourUser]"**, find the row named **Path** and select it.
   - Click **Edit...**, then click **New**.
   - Paste the path to your folder: `C:\tools\logdy`.
4. **Apply and Restart:**
   - Click **OK** on all windows.
   - **IMPORTANT:** Restart any open terminals (PowerShell/VS Code) for the changes to take effect.
5. **Verify and Use:**
   - Type `logdy-rk --version` from any new terminal. It should work!

### Ubuntu

There are two primary ways to install Logdy on Ubuntu.

#### Prerequisite: Install Go (Required for Method 1)
If you don't have Go installed, you can install it using Snap:
```bash
sudo snap install go --classic
```

#### Method 1: Using Go Install
If you have Go installed on your system, this is the most direct method:

```bash
# 1. Install the binary from the repository
go install github.com/rithakith/logdy-core-rk@latest

# 2. Rename it for easier use
mv ~/go/bin/logdy-core-rk ~/go/bin/logdy-rk

# 3. Ensure your Go bin directory is in your PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

#### Method 2: Manual Download (Binary)
If you prefer to download the pre-compiled binary:

```bash
# 1. Download the latest release binary (replace [version] with the actual tag)
curl -L -O https://github.com/rithakith/logdy-core-rk/releases/download/[version]/logdy-rk-linux-amd64

# 2. Grant executable permissions
chmod +x logdy-rk-linux-amd64

# 3. Move it to a directory in your PATH
sudo mv logdy-rk-linux-amd64 /usr/local/bin/logdy-rk
```

#### Verification
Type the following to confirm the installation:
```bash
logdy-rk --version
```

## License
This fork follows the same licensing as the original [Logdy](https://github.com/logdyhq/logdy-core) project.
