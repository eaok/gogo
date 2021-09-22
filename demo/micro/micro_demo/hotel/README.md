# Booking.com Example

# Booking Service 聚合微服务
```
api/hotel：聚合API--通过登⼊入和登出时间查询符合的酒店以及酒店的房价信息
srv/auth：登陆认证服务--VerifyToken服务验证是否合规
srv/geo：地理理信息API--Nearby找到附近的酒店
srv/profile：酒店信息API--通过酒店的Ids查找到所有酒店的信息
srv/rate：酒店登⼊入信息API--酒店的房价信息查询
```

This is [@harlow](https://github.com/harlow)'s [go-micro-services](https://github.com/harlow/go-micro-services) example converted to use Micro.

His README (with required changes):

The API Endpoint accepts HTTP requests at `localhost:8080` and then spawns a number of RPC requests to the backend services.

_Note:_ Data for each of the services is stored in JSON flat files under the `/data/` directory. In reality each of the services could choose their own specialty datastore. The Geo service for example could use PostGis or any other database specializing in geospacial queries.

### Setup

Docker is required for running the services https://docs.docker.com/engine/installation.

Protobuf v3 are required:

    $ brew install protobuf

Install the protoc-gen libraries and other dependencies:

    $ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    $ go get -u github.com/micro/protoc-gen-micro
    $ go get -u github.com/micro/go-micro
    $ go get -u github.com/hailocab/go-geoindex

Clone the repository:

    $ git clone git@github.com:micro/micro.git

Change to examples dir

    $ cd micro/examples/booking

### Protobufs

If changes are made to the Protocol Buffer files use the Makefile to regenerate:

    $ make proto

### Run

To make the demo as straigforward as possible; [Docker Compose](https://docs.docker.com/compose/) is used to run all the services at once (In a production environment each of the services would be run (and scaled) independently).

    $ make build
    $ make run

Curl the endpoint with an invalid auth token:

    $ curl -H 'Content-Type: application/json' \
           -H "Authorization: Bearer INVALID_TOKEN" \
           -d '{"inDate": "2015-04-09"}' \
            http://localhost:8080/hotel/rates

    {"id":"api.hotel.rates","code":401,"detail":"Unauthorized","status":"Unauthorized"}

Curl the endpoint without checkin or checkout dates:

    $ curl -H 'Content-Type: application/json' \
           -H "Authorization: Bearer VALID_TOKEN" \
           -d '{"inDate": "2015-04-09"}' \
            http://localhost:8080/hotel/rates

    {"id":"api.hotel.rates","code":400,"detail":"Please specify inDate/outDate params","status":"Bad Request"}

Curl the API endpoint with a valid auth token:

    $ curl -H 'Content-Type: application/json' \
           -H "Authorization: Bearer VALID_TOKEN" \
           -d '{"inDate": "2015-04-09", "outDate": "2015-04-10"}' \
            http://localhost:8080/hotel/rates

The JSON response:

```json
{
    "hotels": [
        {
            "id": 1,
            "name": "Clift Hotel",
            "phoneNumber": "(415) 775-4700",
            "description": "A 6-minute walk from Union Square and 4 minutes from a Muni Metro station, this luxury hotel designed by Philippe Starck features an artsy furniture collection in the lobby, including work by Salvador Dali.",
            "address": {
                "streetNumber": "495",
                "streetName": "Geary St",
                "city": "San Francisco",
                "state": "CA",
                "country": "United States",
                "postalCode": "94102"
            }
        }
    ],
    "ratePlans": [
        {
            "hotelId": 1,
            "code": "RACK",
            "inDate": "2015-04-09",
            "outDate": "2015-04-10",
            "roomType": {
                "bookableRate": 109,
                "totalRate": 109,
                "totalRateInclusive": 123.17,
                "code": "KNG"
            }
        }
    ]
}
```
