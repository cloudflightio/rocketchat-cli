# Rocket.Chat CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/cloudflightio/rocketchat-cli)](https://goreportcard.com/report/github.com/cloudflightio/rocketchat-cli)

Extremely simple [Rocket.Chat](https://rocket.chat/) CLI Client.

Especially useful for kubernetes/openshift init-containers or other simple automatization tasks.

## Docker

[![](https://images.microbadger.com/badges/image/cloudflight/rocketchat-cli.svg)](https://microbadger.com/images/cloudflightio/rocketchat-cli "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/cloudflight/rocketchat-cli.svg)](https://microbadger.com/images/cloudflightio/rocketchat-cli "Get your own version badge on microbadger.com")

```
docker run --rm -it \
           -e RCCLI_ROCKETCHAT_URL=http://localhost:3000 \
           -e RCCLI_USER_EMAIL=admin@localhost \
           -e RCCLI_USER_PASSWORD=admin \
           cloudflightio/rocketchat-cli ping
```
