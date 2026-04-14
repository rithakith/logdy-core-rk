
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

```bash
# Pipe any log output
$ tail -f app.log | logdy-rk

# Follow specific files
$ logdy-rk follow system.log
```



## Installation

You can find the latest precompiled binaries for all platforms on the [Releases](https://github.com/rithakith/logdy-core-rk/releases) page. Just download the one for your system and add it to your PATH.



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
