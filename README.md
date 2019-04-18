# RocketChat CLI

[![CircleCI](https://circleci.com/gh/mriedmann/rocketchat-cli.svg?style=svg)](https://circleci.com/gh/mriedmann/rocketchat-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/mriedmann/rocketchat-cli)](https://goreportcard.com/report/github.com/mriedmann/rocketchat-cli)

DO NOT USE IN PRODUCTION!

Please wait with testing and contribution till the basic API is stable.

## Docker

[![](https://images.microbadger.com/badges/image/mriedmann/rocketchat-cli.svg)](https://microbadger.com/images/mriedmann/rocketchat-cli "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/mriedmann/rocketchat-cli.svg)](https://microbadger.com/images/mriedmann/rocketchat-cli "Get your own version badge on microbadger.com")

```
docker run --rm -it \
           -e RCCLI_ROCKETCHAT_URL=http://localhost:3000 \
           -e RCCLI_USER_EMAIL=admin@localhost \
           -e RCCLI_USER_PASSWORD=admin \
           mriedmann/rocketchat-cli ping
```
