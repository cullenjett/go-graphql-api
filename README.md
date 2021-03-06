# Go GraphQL API

## Description

This is a demo GraphQL API written in Go using the [`graphql-go/graphql`](https://github.com/graphql-go/graphql) package. During local development it connects to a PostgreSQL database via `docker-compose`.

## Commands

- `make`

  - Run [`graphql-playground`](https://github.com/prisma/graphql-playground) on `http://localhost:3000` and the API on `/api`

- `make stop`

  - Clean up any running docker images that were mishandled by sending SIGINT (ctrl-c) to `make`
