> Kabeer, I am the Lord's dog; Moti is my name. -- [Kabir](https://www.hinduwebsite.com/sacredscripts/sikhscripts/guru940.asp)

# Moti

Moti is a simple cross-platform desktop app for searching in text files.

## Technology stack

- Elektron
- [Hydraulic Conveyor](https://conveyor.hydraulic.dev/)

## Development

```sh
export GITHUB_TOKEN=xxx
conveyor run
```

## Maintenance

```sh
# Current OS
conveyor make app

# Windows
conveyor make windows-app

# macOS
conveyor -Kapp.machines=mac.amd64 make mac-app
conveyor -Kapp.machines=mac.aarch64 make mac-app
```
