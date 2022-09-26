# openport

[![CodeQL Analysis](https://github.com/ahmad4fifz/openport/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/ahmad4fifz/openport/actions/workflows/codeql-analysis.yml)
[![Go Build](https://github.com/ahmad4fifz/openport/actions/workflows/go-build.yml/badge.svg)](https://github.com/ahmad4fifz/openport/actions/workflows/go-build.yml)

Currently only tcp supported. udp to-do

## Installing

```bash
go install github.com/ahmad4fifz/openport@latest
```

## Usage

### Remote server

```bash
openport -p <1-65535>
```
### Local server

```bash
nc -zv <remote_ip> <port>
```
