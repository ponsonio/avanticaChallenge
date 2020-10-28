# Avantica Challenge

##Design Notes
* This is designed using interfaces, as an attempt to make easy unit test injecting mock, although 
per time constraint not all test has been development neither all code present the facilities to easily inject mocks.

*Test has been neglected, coverage should be ~85%, only one test has been included to demonstrate test skills

*Security hasn't been addressed as a concern at all, because requirements are empty about that and time constraint

* Implementation is somehow _relational_ between collections, it should enforce valid ID references, this is not implemented.

* Dockerize the app should be simple to be deployed as a microservice 

* Check for comments in the code for detailed notes

##Local dev
In order to use the app locally you'll need :

* Go 1.15.12
* Local mongo db on port 27017 with localhost access exception

You can change the connection parameters on code on _app.go_ (yeah it needs to be ab os env var!) 

mongodb://localhost:27017

##Usage
Once the db it's setup, you can use 

````
go run main.go (inside the app directory)
````

the following request can be used to test the application (replace the id with the adequate value):

Spot:
````

curl --location --request POST '127.0.0.1:8080/api/maze/spot/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "My Spot 2",
    "gold": 10.11,
    "quadrant" : "5f9853b91987d26ce300c532",
    "point": {
        "x": 1,
        "y": 2
    }
}'

curl --location --request DELETE '127.0.0.1:8080/api/maze/spot/5f997db2dc5fbe2aa00de7fd'

curl --location --request GET '127.0.0.1:8080/api/maze/spot/5f997c6cdc5fbe2aa00de7fb'
````

Path:

````
curl --location --request POST '127.0.0.1:8080/api/maze/path/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "a": "5f997c6cdc5fbe2aa00de7fb",
    "b": "5f997c6cdc5fbe2aa00de7fc0",
    "distance": 100.10
}'

curl --location --request GET '127.0.0.1:8080/api/maze/path/5f99913028fcd620a86f9276'

curl --location --request PUT '127.0.0.1:8080/api/maze/path/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "5f99913028fcd620a86f9276",
    "distance": 1023.1,
    "a": "5f997c6cdc5fbe2aa00de7fb",
    "b": "5f997c6cdc5fbe2aa00de7fc0"
}'

````

Quadrant:

A cuandrant is defined by a center (relative to 0,0) and a number (1,2,3,4), following the 
standard geometric naming for them

2 | 1
-----
3 | 4

https://en.wikipedia.org/wiki/Quadrant_(plane_geometry)#:~:text=The%20axes%20of%20a%20two,%2C%20and%20IV%20(%2B%3B%20%E2%88%92).
````
curl --location --request POST '127.0.0.1:8080/api/maze/quadrant/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "number": 1,
    "center": {
        "x": 1,
        "y": 2
    }
}'

curl --location --request GET '127.0.0.1:8080/api/maze/quadrant/5f99948ad67eb9b786de987e'

curl --location --request PUT '127.0.0.1:8080/api/maze/quadrant/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "5f99948ad67eb9b786de987e",
    "number": 2,
    "center": {
        "x": 1,
        "y": 1
    }
}'

curl --location --request DELETE '127.0.0.1:8080/api/maze/quadrant/5f99948ad67eb9b786de987e'

````