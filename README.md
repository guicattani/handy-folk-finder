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
Tests need to be run individually, go test ./... breaks because of dotenv for now. You can manually test the other features with some mock data for now:
```
INSERT INTO partners (first_name,last_name,email,address_lat, address_lon, operating_radius, experience, created_at, updated_at)  VALUES('Jackie', 'Brown', 'jackie@brown.email', 1.52, 0.3, 2, '{"wood"}', NOW(),NOW() ) /*This partner will be 2 km away*/
INSERT INTO partners (first_name,last_name,email,address_lat, address_lon, operating_radius, experience, created_at, updated_at)  VALUES('Billie', 'Jean', 'billie@jean.email', 1.501, 0.3, 2, '{"wood"}', NOW(),NOW() ) /*This partner will be very close to the customer*/
INSERT INTO partners (first_name,last_name,email,address_lat, address_lon, operating_radius, experience, created_at, updated_at)  VALUES('Jibbie', 'Hibbie', 'jibbie@hibbie.email', 1.51, 0.3, 2, '{"wood" , "stone"}', NOW(),NOW() ) /*This partner close to the customer and he works with stone*/

/*You can use this customer to 'login'*/
INSERT INTO customers (first_name,last_name,email, "password", address_lat, address_lon, created_at, updated_at)  VALUES('ABBY', 'ABBEY', 'a@a.com','123456',  1.5, 0.3, NOW(),NOW() )

/*You can change the rating like so*/
UPDATE partners
SET rating = 10
WHERE first_name = 'Jibbie'

/*
Queries will then be something like this
http://localhost:8080/closest_partner?email=a@a.com&password=123456&needed_experience="wood"
http://localhost:8080/closest_partner?email=a@a.com&password=123456&needed_experience="wood","stone"
*/

```


Endpoints
```
http://localhost:8080/ # List all partners
http://localhost:8080/partner?id=PARTNER_ID # List specific partner
http://localhost:8080/closest_partner?email=USER_EMAIL&password=PASSWORD&needed_experience=NEEDED_EXPERTISE # List closest partner based on distance, NEEDED_EXPERTISE is separated by commas ans needs to be between double quotes e.g.: "wood","stone"
```