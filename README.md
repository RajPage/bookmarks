# Bookmarks / Read Later

[![Go](https://github.com/RajPage/bookmarks/actions/workflows/go-build.yml/badge.svg)](https://github.com/RajPage/bookmarks/actions/workflows/go-build.yml)

Work in Progress. 

## Installation

### Server Setup
1. Install dependencies
    ```bash
    go mod download
    ```
2. Setup the DB.
    ```sql
    CREATE USER bookmarks WITH PASSWORD bookmarks; 
    CREATE DATABASE bookmarks; 
    GRANT ALL PRIVILEGES ON DATABASE bookmarks TO bookmarks; 
    ```
3. Set env variables
    ```ini
    BOOKMARKS_SERVER_ENVIRONMENT="development"
    EnvDbDsn="host=localhost user=bookmarks password=bookmarks dbname=bookmarks port=5432 sslmode=disable TimeZone=Asia/Kolkata"
    ```
4. Run the server
    ```bash
    go run main.go
    ```