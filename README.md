# Logdy (rk-version)

A personal fork of Logdy that incorporates various improvements based on community feedback and my own development experience.

---

## Acknowledgments
All credit for the core engine and design goes to the original **[Logdy Team](https://github.com/logdyhq)**. This fork simply builds upon their work to add a few specific tools I needed.

---

## Added Improvements

These modifications address specific workflow friction and UI layout issues:

*   **Custom Highlighting**: Allows you to highlight specific words or phrases with custom colors, making it much easier to visually scan and identify important log entries.
    <video src="assets/highlighting.mp4" width="100%" controls></video>
*   **Live Search & Filtering**: Results filter instantly as you type.
    *   **Instant Text & RegEx Search**: Search through raw logs using simple text or powerful Regular Expressions.
        - *Example:* `error`
        <video src="assets/live_search.mp4" width="100%" controls></video>
    *   **Structured JSON Queries**: Deeply filter logs by querying specific JSON fields.
        - *Example:* `data.level == "trace"`, `data.method == "PATCH"`, `data.duration == "441"`, `data.domain == "nationalbest-of-breed.com"`
        <video src="assets/query_search.mp4" width="100%" controls></video>
*   **Show in Context**: If you find an entry through search but want to see the surrounding logs, you can jump directly into its original context with a single click.
    <video src="assets/show_in_context.mp4" width="100%" controls></video>



---

## Quick Start

You can use `logdy-rk` just like the original tool:

```bash
# Pipe any log output
$ tail -f app.log | logdy-rk

# Follow specific files
$ logdy-rk follow system.log
```

---

## Installation

You can find the latest precompiled binaries for all platforms on the [Releases](https://github.com/rithakith/logdy-core/releases) page. Just download the one for your system and add it to your PATH.

---

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

---

## License
This fork follows the same licensing as the original [Logdy](https://github.com/logdyhq/logdy-core) project.
