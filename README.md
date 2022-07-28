# Handy Folk finder

This project uses
- Pop (from the Buffalo package) -> creation of migrations
- NoSurf -> CSRF Middleware
- Chi -> Router
- PGX -> Postgres Driver
- Scany -> Better PGX scan

Server run instructions:
- Rename .env.example to .env
- Rename database.yml.example to database.yml

Make sure that you have [Soda CLI](https://gobuffalo.io/documentation/database/soda/) in your path, then run:
```
soda create
```