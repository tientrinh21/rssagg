# RSS Aggregatar

A HTTP Server which continuously retreive new feeds to provide an all-in-one place for user to update news from users' own favorite source.
Developer can create a front-end UI to display data fetched from this as posts, check out this example [RSS Reader](https://github.com/tientrinh21/rssreader)

## Tech Used

-   **Languages:** Go
-   **Library:** [google/uuid](https://github.com/google/uuid), [go-chi](https://go-chi.io/), [godotenv](https://github.com/joho/godotenv)
-   **Database:** PostgreSQL
-   **Tools:**
    -   [goose](https://github.com/pressly/goose) - database migration
    -   [sqlc](https://sqlc.dev/) - SQL compiler

## Quick Start

Create `.env` file and assign `PORT` value and `DB_URL` to your local PostgreSQL database.

```sh
cd rssagg
go build && ./rssagg
```

## How to use

- **Healthcheck - Ready / Error**

```http
GET http://localhost:PORT/v1/ready
GET http://localhost:PORT/v1/error
```

- **User - Create / Get**

  - Create user:
```http
POST http://localhost:PORT/v1/users

{
  name: "[USER_NAME]"
}
```
The received response will contain the `apiKey`, this key will be use to get user, add feed, etc.

  - Get user:
```http
GET http://localhost:PORT/v1/users
Authorization: ApiKey [API_KEY]
```
- **Feed - Add feed / Get feed lists**

  - Add feed:
```http
POST http://localhost:PORT/v1/feeds
Authorization: ApiKey [API_KEY]

{
  name: "[FEED_NAME]"
  url: "[RSS_URL]"
}
```

  - Get feed lists:
```http
GET http://localhost:PORT/v1/feeds
```


