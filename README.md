# openport

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
