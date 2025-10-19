> Kabeer, I am the Lord's dog; Moti is my name. – [Guru Granth Sahib](https://www.hinduwebsite.com/sacredscripts/sikhscripts/guru940.asp)

# Moti [![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

a cross-platform desktop app for searching in text files

## Technology stack

- Backend: Go, [bleve](https://github.com/blevesearch/bleve)
- Frontend: React, Bootstrap
- Infrastructure: [Wails](https://wails.io/)

## Development

```sh
wails dev
```

## Maintenance

```sh
wails build

# Create macOS binary
wails build -platform darwin/arm64
wails build -platform darwin/amd64

# Create Windows binary
wails build -platform windows/amd64
```

**Also see**

- https://wails.io/docs/reference/cli#platforms

## Alternatives

- [DocFetcher](https://docfetcher.sourceforge.io/)
- [Copernic Desktop Search](https://copernic.com)

## License

[MIT](./LICENSE)
