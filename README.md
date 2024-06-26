# **Telemonitor** - simple *Telegram* bot for system monitoring

Deploy `Telemonitor` instance on your server and monitor metrics via your phone
+ *(only works with `Ubuntu` and `Debian` based OS')*


## Deployment:
+ Deployment is much easier than in `Grafana` or other monitoring utilities

1) Install `Docker` and `Docker compose` with [this](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04) guide

2) Clone `GitHub` repo and move to created dir
```sh
git clone https://github.com/LCcodder/Telemonitor && cd Telemonitor
```
3) Paste your `Telegram` bot token in `TOKEN` variable inside `Dockerfile`
4) Configure whitelist in `config.json` file *(paste usernames, not display names)*
5) Build and run `Docker` container
```sh
docker build . -t telemonitor && docker run telemonitor
```
---
Created with `go-telegram`, `docker` proprietary lib and `sysinfo`
