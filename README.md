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

Make sure that you have [Soda CLI](https://gobuffalo.io/documentation/database/soda/) in your path to run migrations, then run:
```
soda create
```

Start server
```
go run ./cmd/web
```

Run tests
```
Tests need to be run individually, go test ./... breaks because of dotenv for now
```


Endpoints
```
http://localhost:8080/ # List all partners
http://localhost:8080/partner?id=PARTNER_ID # List specific partner
http://localhost:8080/closest_partner?email=USER_EMAIL&password=USER_EMAIL&needed_experience=NEEDED_EXPERTISE # List closest partner based on distance, NEEDED_EXPERTISE is separated by commas ans needs to be between double quotes e.g.: "wood","stone"
```