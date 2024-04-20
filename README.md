# **Telemonitor** - simple *Telegram* bot for system monitoring

Deploy `Telemonitor` instance on your server and monitor metrics via your phone
+ *(only works with `Ubuntu` and `Debian` based OS')*


## Deployment:
+ Deployment is much easier than in `Graphana` or other monitoring utilities

1) Install `Docker` and `Docker compose` with [this](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04) guide

2) Clone `GitHub` repo and move to created dir
```sh
git clone https://github.com/LCcodder/Telemonitor && cd Telemonitor
```
3) Paste your `Telegram` bot token in `TOKEN` variable inside `docker-compose.yaml`
4) Configure whitelist in `config.json` file *(paste usernames, not display names)*
5) Run `docker-compose.yaml` file
```sh
docker compose up --build && docker compose up
```
---
Created with `go-telegram`, `docker` proprietary lib and `sysinfo`