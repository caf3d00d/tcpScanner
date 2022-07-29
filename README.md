# tcpScanner

**A simple network scanning written in go**

### Why?
    * well, educational purpose maybe.

### Usage:
Build the package:
```shell
go build
```

Scanning opened ports:
```shell
./tcpScanner hostname from to
```

Example:
```shell
./tcpScanner hostname <range of ports>
./tcpScanner scanme.nmap.org 21 81
```

### Attention

This script is under active maintain, although it is not perfect.
There might be unexpected behaviour as it would grow up, so feel free to contribute,
I'm open to learn.

### TODO
* Enhance the performance using goroutine technique.
* Add more scanning technique.
* Fix packet loosing issue, maybe?
* Refactor the code.
* Write more TODO.

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

[![forthebadge](https://forthebadge.com/images/badges/works-on-my-machine.svg)](https://forthebadge.com)
