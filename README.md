# DOC API

## CLI

```
cd cmd/cli
go build
./cli
```

## Web

```
cd cmd/web
go build
./server
```

## Setup
```
mkdir -p $HOME/.config/doc-api
cp cmd/cli/config.json $HOME/.config/doc-api/
sudo mkdir /var/lib/doc-api
sudo chown $(whoami) /var/lib/doc-api
```
