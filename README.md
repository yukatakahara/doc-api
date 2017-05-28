# DOC API

## CLI

```
cd cmd/cli
go build
```

## Web

```
cd cmd/web
go run server.go
```

## Setup
```
mkdir -p $HOME/.config/doc-api
cp cmd/cli/config.json $HOME/.config/doc-api/
sudo mkdir /var/lib/doc-api
sudo chown $(whoami) /var/lib/doc-api
```
