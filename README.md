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

-   **Healthcheck - Ready / Error**

```http
GET http://localhost:PORT/v1/ready
GET http://localhost:PORT/v1/error
```
