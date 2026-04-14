
# Logdy (rk-version)

A personal fork of Logdy that incorporates various improvements based on community feedback and my own development experience.


## Acknowledgments
All credit for the core engine and design goes to the original **[Logdy Team](https://github.com/logdyhq)**. This fork simply builds upon their work to add a few specific tools I needed.



## Added Improvements

These modifications address specific workflow friction and UI layout issues:

*   **Custom Highlighting**: Allows you to highlight specific words or phrases with custom colors, making it much easier to visually scan and identify important log entries.

  
https://github.com/user-attachments/assets/6d2f0b13-dbb9-4f7e-b4ad-cb08619dfe6a


*   **Live Search & Filtering**: Results filter instantly as you type.
    *   **Instant Text & RegEx Search**: Search through raw logs using simple text or powerful Regular Expressions.
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

**Note on `-t` (Fallthrough):** Use the `-t` flag if you want to see the logs in your terminal *and* the Web UI simultaneously.

```

## Installation & PATH Setup

To run `logdy-rk` from any folder without typing the full path (`.\`), you must add it to your Windows PATH.

### Step-by-Step: Adding to Windows PATH
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
4. **Apply & Restart:**
   - Click **OK** on all windows.
   - **IMPORTANT:** Restart any open terminals (PowerShell/VS Code) for the changes to take effect.
5. **Verify & Use:**
   - Type `logdy-rk --version` from any new terminal. It should work!
   - Now you can run any process from any folder. For example, to start a web project and see logs in Logdy:
     ```powershell
     cd C:\path\to\your\project
     logdy-rk stdin "npm run dev" -t
     ```

### Alternative: Go Install
If you have Go installed, this handles the PATH setup automatically (to your GOBIN):
```bash
go install github.com/rithakith/logdy-core-rk@latest
```



## Building/Releasing

To build cross-platform binaries:
```bash
# Requires GoReleaser
goreleaser build --clean --snapshot
```

To create a new release on GitHub, simply push a new tag:
```bash
git tag v0.17.2-rk
git push origin v0.17.2-rk
```


## License
This fork follows the same licensing as the original [Logdy](https://github.com/logdyhq/logdy-core) project.
