# newsbot.collector.api

This is the collection service of newsbot to pull articles from the web.

## Deployment

1. Create a copy of the docker compose file and make it local
2. Update the `docker-compose.yaml` with your secrets
3. Run migrations
   2. `docker compose run api /app/goose -dir "/app/migrations" up`
4. Run app
   1. `docker compose up -d`
5. Once the app is running go to the swagger page and validate that you see the seeded sources.
   1. `http://localhost:8081/swagger/index.html#/Source/get_config_sources`
   2. `curl -X 'GET' 'http://localhost:8081/api/config/sources' -H 'accept: application/json'`
6. Add any new sources
7. Add a Discord Web Hook
   1. `curl -X 'POST' 'http://localhost:8081/api/discord/webhooks/new?url=WEBHOOKURL&server=SERVERNAME&channel=CHANNELNAME' -H 'accept: application/json' -d ''`
8. Create your subscription links
   1. This is a link between a source and a discord web hook.  Without this, the app will not send a notification about new posts.

### Errors

- pq: permission denied to create extension "uuid-ossp"
  - Might need to grant your account `ALTER USER root WITH SUPERUSER;` to create the 'uuid-ossp' for uuid creations
