# Logdy (rk-version)

A personal fork of Logdy that incorporates various improvements based on community feedback and my own development experience.

---

## Acknowledgments
All credit for the core engine and design goes to the original **[Logdy Team](https://github.com/logdyhq)**. This fork simply builds upon their work to add a few specific tools I needed.

---

## ✨ Added Improvements

These modifications address specific workflow friction and UI layout issues:

*   **Custom Highlighting**: Allows you to highlight specific words or phrases with custom colors, making it much easier to visually scan and identify important log entries.
    [INSERT_SNIPPET_HERE]
*   **Live Search & Filtering**: Supports instant filtering as you type. It also handles Regular Expressions and Structured JSON Queries (e.g., `data.user == "admin"`) for more detailed data exploration.
    [INSERT_SNIPPET_HERE]
*   **Show in Context**: If you find an entry through search but want to see the surrounding logs, you can jump directly into its original context with a single click.
    [INSERT_SNIPPET_HERE]
*   **Rendering Performance**: UI optimizations to prevent freezes and ensure smooth log navigation during high-volume bursts.
    [INSERT_SNIPPET_HERE]


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
