# candystore-exercise

A service that scrapes data from [Zimpler's Candystore Website](https://candystore.zimpler.net), processes it, and returns a list of its top customers and their favorite snacks.

## Running the service locally in Docker

* in root directory: `make up`

## Stopping the service running locally in Docker

* in root directory: `make down`

# Endpoint:

## Get a list of top customers and their favorite snacks:
`GET :8080/top_customer_favorite_snacks`

### An example 200 ok response:

```json 
[
  {
    "name":"Jonas",
    "favoriteSnack":"Geisha",
    "totalSnacks":1800
  },
  {
    "name":"Annika",
    "favoriteSnack":"Geisha",
    "totalSnacks":200
  },
  {
    "name":"Jane",
    "favoriteSnack":"Nötchoklad",
    "totalSnacks":22
  },
  {
    "name":"Aadya",
    "favoriteSnack":"Center",
    "totalSnacks":9
  }
]
```