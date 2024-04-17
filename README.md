# RSS Aggregator

A HTTP Server which continuously retreive new feeds to provide an all-in-one place for user to update news from users' own favorite source.
Developer can create a front-end UI to display data fetched from this as posts, check out this example [RSS Reader](https://github.com/tientrinh21/rssreader)

## Tech Used

-   **Languages:** Go
-   **Library:** [google/uuid](https://github.com/google/uuid), [go-chi](https://go-chi.io/), [godotenv](https://github.com/joho/godotenv)
-   **Database:** PostgreSQL
-   **Tools:**
    -   [goose](https://github.com/pressly/goose) - database migration
    -   [sqlc](https://sqlc.dev/) - SQL compiler
-   **Host:** [Supabase](https://supabase.com/)

## Quick Start

Create `.env` file and assign `PORT` value and `DB_URL` to your local PostgreSQL database. Then add `SERVICE_URL` as the server URL.

For example:
```python
PORT=8000
DB_URL=postgres://[USER]:root@localhost:5432/rssagg
SERVICE_URL='http://localhost:8000'
```

Then run:

```sh
cd rssagg
go build && ./rssagg
```

## How to use

### Healthcheck - Ready / Error

```http
GET http://localhost:PORT/v1/ready
GET http://localhost:PORT/v1/error
```

### User - Create / Get

- Create user:
```http
POST http://localhost:PORT/v1/users
```
```json
{
  "name": "[USER_NAME]"
}
```

The received response will contain the `apiKey`, this key will be use to get user, add feed, etc.

- Get user:
```http
GET http://localhost:PORT/v1/users
Authorization: ApiKey [API_KEY]
```
### Feed - Add feed / Get feed lists / Get feed with ID

- Add feed:
```http
POST http://localhost:PORT/v1/feeds
Authorization: ApiKey [API_KEY]
```
```json
{
  "name": "[FEED_NAME]"
  "url": "[RSS_URL]"
}
```

- Get feed lists:
```http
GET http://localhost:PORT/v1/feeds
```

### Follows - Get follow lists / Follow a feed / Unfollow a feed

- Get follow lists
```http
GET http://localhost:PORT/v1/feed_follows
Authorization: ApiKey [API_KEY]
```
- Follow a feed:
```http
POST http://localhost:PORT/v1/feed_follows
Authorization: ApiKey [API_KEY]
```
```json
{
  "feed_id": "[FEED_ID]"
}
```
- Unfollow a feed:
```http
DELETE http://localhost:PORT/v1/feed_follows/{feedFollowID}
Authorization: ApiKey [API_KEY]
```

### Posts - Get posts with feeds ID / Get posts according to user's follows

- Get posts with feeds ID:
```http
GET http://localhost:PORT/v1/feeds/{feedID}
```

- Get posts according to user's follows
```http
GET http://localhost:PORT/v1/feeds/{feedID}
Authorization: ApiKey [API_KEY]
```
