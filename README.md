# deer-api
API Demo based on Revel framework with JWT Authentication.

### About JWT Authetication

Using jwt-go lib. Setting user's email as `email` in cliams.

It behaviors like knock ruby gem. https://github.com/nsarno/knock

### Authenticating from a web or mobile application

1. register user
```
curl -v -X POST -d 'email="test@test.com"&password="your_pass"' http://localhost:9001/register
```

the response   
```
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5ldHF5cUAxNjMuY29tIiwibmJmIjoxNDQ0NDc4NDAwfQ.bZo1DzrzZBetB9IP7fVip5XA_GiFBb_z8zDNTalReuU"
```


2. request to get a token from your API:
```
POST /login
{"auth": {"email": "foo@bar.com", "password": "secret"}}
```

Example response from the API:
```js
{
    "result": "login success",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5ldHF5cUAxNjMuY29tIiwibmJmIjoxNDQ0NDc4NDAwfQ.bZo1DzrzZBetB9IP7fVip5XA_GiFBb_z8zDNTalReuU"
}
```

To make an authenticated request to your API, you need to pass the token via the request header:
```
Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9...
GET /my_resources
POST /my_resources
...
```


### Basic CRUD Example
`resources :products`

create product
```
curl -v -X POST -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5ldHF5cUAxNjMuY29tIiwibmJmIjoxNDQ0NDc4NDAwfQ.bZo1DzrzZBetB9IP7fVip5XA_GiFBb_z8zDNTalReuU"  -d 'name="tomato"&price="12.0"&code="T0001"}' "http://localhost:9001/products"
```

show product
```
curl -v -X GET -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5ldHF5cUAxNjMuY29tIiwibmJmIjoxNDQ0NDc4NDAwfQ.bZo1DzrzZBetB9IP7fVip5XA_GiFBb_z8zDNTalReuU"   "http://localhost:9001/products/1"
```

index product 
```
curl -v -X GET -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5ldHF5cUAxNjMuY29tIiwibmJmIjoxNDQ0NDc4NDAwfQ.bZo1DzrzZBetB9IP7fVip5XA_GiFBb_z8zDNTalReuU"   "http://localhost:9001/products"
```

update product
```
curl -v -X PUT -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5ldHF5cUAxNjMuY29tIiwibmJmIjoxNDQ0NDc4NDAwfQ.bZo1DzrzZBetB9IP7fVip5XA_GiFBb_z8zDNTalReuU"  -d 'name="tomato"&price="12.0"&code="T0001"}' "http://localhost:9001/products/1"
```

delete product
```
curl -v -X DELETE -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5ldHF5cUAxNjMuY29tIiwibmJmIjoxNDQ0NDc4NDAwfQ.bZo1DzrzZBetB9IP7fVip5XA_GiFBb_z8zDNTalReuU"  "http://localhost:9001/products/1"
```



### Press Testing
see docs/press_testing.md


### JWT Secret Key 
change `hmacSecret` var to your own.


### Start the web server:

   revel run deer-api



### Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites




### Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).
