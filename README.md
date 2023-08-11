# 🍿 Greenlight

_Greenlight_ is a movies REST API written in Go, with a PostgreSQL database. It features pagination, user authentication and authorization. The API supports the following endpoints and actions:

| Method | URL Pattern               | Action                                          |
| :----- | :------------------------ | :---------------------------------------------- |
| GET    | /v1/healthcheck           | Show application health and version information |
| GET    | /v1/movies                | Show details of all the movies                  |
| POST   | /v1/movies                | Create a new movie                              |
| GET    | /v1/movies/:id            | Show the details of a specific movie            |
| PATCH  | /v1/movies/:id            | Update the details of a specific movie          |
| DELETE | /v1/movies/:id            | Delete a specific movie                         |
| POST   | /v1/users                 | Register a new user                             |
| PUT    | /v1/users/activated       | Activate a specific user                        |
| PUT    | /v1/users/password        | Update the password for a specific user         |
| POST   | /v1/tokens/authentication | Generate a new authentication token             |
| POST   | /v1/tokens/password-reset | Generate a new password-reset token             |
| GET    | /debug/vars               | Display application metrics                     |

## ⚙️ Setup

You'll need to set up the PostgreSQL database using the `greenlight.sql` file and then running

```bash
$ make db/migrations/up
```

Also, you should put the database Data Source name in a `.envrc` file, it should look something like this:

```bash
export GREENLIGHT_DB_DSN=postgres://greenlight:yourpassword@localhost/greenlight?sslmode=disable
```

Finally, you'll need your own SMTP server. For testing, you can use [Mailtrap](https://mailtrap.io/) and replace the host, port, username and password using the command line flags (more on this later).

Now, you can build the application using

```bash
$ make build/api
```

and run it using:

```bash
./bin/api -db-dsn=postgres://greenlight:yourpassword@localhost/greenlight?sslmode=disable -smtp-host=yourhost -smtp-port=yourport -smtp-username=yoursmtpusername -smtp-password=yoursmtppassword
```

(Yes, it's a bit of a pain to set up)

## 🎬 Examples

### User registration

Request

```
$ BODY='{"name": "Robert Deniro", "email": "robertdeniro@example.com", "password": "password"}'
$ curl -d "$BODY" localhost:4000/v1/users
```

Response

```json
{
  "user": {
    "id": 22,
    "created_at": "2023-07-18T17:59:37-03:00",
    "name": "Robert Deniro",
    "email": "robertdeniro@example.com",
    "activated": false
  }
}
```

### User activation

Request

```
$ curl -X PUT -d '{"token": "P4B3URJZJ2NW5UPZC2OHN4H2NM"}' localhost:4000/v1/users/activated
```

Response

```json
{
  "user": {
    "id": 22,
    "created_at": "2023-07-18T17:59:37-03:00",
    "name": "Robert Deniro",
    "email": "robertdeniro@example.com",
    "activated": true
  }
}
```

### User authentication

Request

```
$ BODY='{"email": "robertdeniro@example.com", "password": "password"}'
$ curl -i -d "$BODY" localhost:4000/v1/tokens/authentication
```

Response

```json
{
  "authentication_token": {
    "token": "QBR545EJGTD3NOBU4BP7SRYLMM",
    "expiry": "2023-07-19T18:06:01.4111448-03:00"
  }
}
```

### Getting all movies

Request

```
curl -H "Authorization: Bearer QBR545EJGTD3NOBU4BP7SRYLMM" localhost:4000/v1/movies
```

Response

```json
{
  "metadata": {
    "current_page": 1,
    "page_size": 20,
    "first_page": 1,
    "last_page": 1,
    "total_records": 3
  },
  "movies": [
    {
      "id": 2,
      "title": "Blade Runner 2049",
      "year": 2017,
      "runtime": "164 mins",
      "genres": ["action", "drama", "mystery", "sci-fi", "thriller"],
      "version": 1
    },
    {
      "id": 3,
      "title": "Drive",
      "year": 2011,
      "runtime": "100 mins",
      "genres": ["action", "drama"],
      "version": 1
    },
    {
      "id": 4,
      "title": "Her",
      "year": 2013,
      "runtime": "126 mins",
      "genres": ["drama", "romance", "sci-fi"],
      "version": 1
    }
  ]
}
```

### Get movies filtered by title and/or genre

Request

```
curl -H "Authorization: Bearer QBR545EJGTD3NOBU4BP7SRYLMM" localhost:4000/v1/movies?title=runner
```

Response

```json
{
  "metadata": {
    "current_page": 1,
    "page_size": 20,
    "first_page": 1,
    "last_page": 1,
    "total_records": 1
  },
  "movies": [
    {
      "id": 2,
      "title": "Blade Runner 2049",
      "year": 2017,
      "runtime": "164 mins",
      "genres": ["action", "drama", "mystery", "sci-fi", "thriller"],
      "version": 1
    }
  ]
}
```

### Get a specific movie by id

Request

```
curl -H "Authorization: Bearer QBR545EJGTD3NOBU4BP7SRYLMM" localhost:4000/v1/movies/4
```

Response

```json
{
  "movie": {
    "id": 4,
    "title": "Her",
    "year": 2013,
    "runtime": "126 mins",
    "genres": ["drama", "romance", "sci-fi"],
    "version": 1
  }
}
```
