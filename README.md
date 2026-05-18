# Blog Aggregator in Golang

## Lightweight RSS Fetcher, scrape posts and follow them also browse offline.

This is a Boot.dev course project.

Blog Aggregator is a compact Go service for consolidating blogs and posts into a single local store. It periodically fetches RSS feeds, scrapes article content when feeds are partial, saves posts in a simple database, and exposes commands and handlers for adding feeds, following/unfollowing, and browsing content. Perfect for developers who want a no‑friction, self‑hosted feed reader.

## Installation & Setup

### Prerequisites

You need to have **Go** and **PostgreSQL** installed on your system to run and build this project.

#### macOS

```bash
brew install go
brew install postgresql@18
brew services start postgresql@18
```

#### Linux

```
sudo apt update
sudo apt install golang-go postgresql-contrib
sudo systemctl start postgresql
```

#### Windows (winget)

To install Go Programming Language with winget, use the following command:

```
winget install -e --id GoLang.Go
```

To install PostgreSQL 18 with winget, use the following command:

```
winget install -e --id PostgreSQL.PostgreSQL.18
```

### Installing gator CLI

If you want to install **gator CLI** you can use go install command like this:

```
go install github.com/emrebozdogan/blog-aggregator-go@latest
```

### Configuration

Gator requires a connection to a PostgreSQL database to store users and feeds.

- Create the Database:
  ```
  CREATE DATABASE gator;
  ```
- Setup Config File:
  ```
  {
    "db_url": "postgres://username:@localhost:5432/gator?sslmode=disable"
  }
  ```
  The **?sslmode=disable** flag is required for most local PostgreSQL installations to avoid >connection errors.

### Usage

| Command     | Description                                                    | Example Usage                    |
| :---------- | :------------------------------------------------------------- | :------------------------------- |
| `register`  | Create a new user and log in automatically                     | `gator register <name>`          |
| `login`     | Switches the current user                                      | `gator login <name>`             |
| `users`     | Lists all of the registered users                              | `gator users`                    |
| `addfeed`   | Adds a new RSS feed to the database.                           | `gator addfeed <name> <url>`     |
| `feeds`     | View all available feeds                                       | `gator feeds`                    |
| `follow`    | Follows an existing feed in the system. (for the current user) | `gator follow <url>`             |
| `unfollow`  | Unfollows a followed feed. (for the current user)              | `gator unfollow <url>`           |
| `following` | List all followed feeds. (for the current user)                | `gator following`                |
| `agg`       | Fetch posts from followed feed (e.g., 1m, 30s, 1h)             | `gator agg <time>`               |
| `browse`    | Displays the latest fetched posts.                             | `gator browse <limit by number>` |
| `reset`     | Deletes all existing users and data                            | `gator reset`                    |

### Quick Start

After you installed the gator you can start with these commands:

- Register:
  ```
  gator register name
  ```
- Login:
  ```
  gator login name
  ```
- Add and follow a feed:
  ```
  gator addfeed "Boot.dev Blog" https://www.boot.dev/blog/index.html
  ```
- Aggregate posts:

  Fetch new posts every minute:

  ```
  gator agg 1m
  ```
