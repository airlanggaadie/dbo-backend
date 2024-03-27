# DBO ASSIGNMENT TEST

- [DBO ASSIGNMENT TEST](#dbo-assignment-test)
  - [HOW TO USE](#how-to-use)
  - [ERD](#erd)
  - [ADVANCED CONFIGURATION](#advanced-configuration)
  - [API DOCUMENTATION](#api-documentation)

## HOW TO USE

- Run PostgreSQL with docker: `docker compose up -d db`
- Run main program: `go run .` and check out the [API Documentation](#api-documentation)
- Run main program with docker: `docker compose up -d api`

## ERD

![ERD](./docs/erd.png)

Table List:
- user
- user_password
- user_login_history
- order
- order_item

## ADVANCED CONFIGURATION

Set environment variables on your operating system (or docker-compose.yml service api if run main program on docker) with the following options:
  - DATABASE_URL: the database connection, default is `postgres://postgres:postgres@localhost:5432/dbo?sslmode=disable`.
  - ENVIRONMENT: the application environment, default is `development`.
  - PORT: application running port number, default is `3000`.
  - APP_TZ: application time zone, default is `Asia/Jakarta`.

## API DOCUMENTATION

[here...](./docs/api-docs.md)