# Kumaoni Backend

For now this is just a POC

## SSL

```bash
docker-compose run --rm --entrypoint "certbot certonly --dns-cloudflare --dns-cloudflare-credentials /path/to/cloudflare.ini -d kumaoni-api.dailycommit.dev --dry-run" certbot
```
