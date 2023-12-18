<!-- <p align="center">
<img src="https://pepy.tech/badge/rss-aggregator" alt="https://pepy.tech/project/rss-aggregator">
<img src="https://pepy.tech/badge/rss-aggregator/month" alt="https://pepy.tech/project/rss-aggregator">
<img src="https://img.shields.io/github/license/Jhnvlglmlbrt/rss-aggregator.svg" alt="https://github.com/Jhnvlglmlbrt/rss-aggregator/blob/master/LICENSE"> -->

# ⚙️ RESTful CRUD API 


## ❗ Requirements

Поменять под себя данные в config.yaml

## 💿 Installation

```
go get 
```

<!-- ## 💻 Example -->

## 🪛 How to use?

```
make run
```

## Queries

### /api/tags

- **GET:** `http://localhost:8080/api/tags` 

- **GET:** `http://localhost:8080/api/tags/tagID`

- **POST:** `http://localhost:8080/api/tags`
    ```json example
    {
	"name":"C*"
    }
    ```

- **PATCH:** `http://localhost:8080/api/tags/tagID`
   ```json example
    {
	"name":"C* New"
    }
    ```

- **DEL:** `http://localhost:8080/api/tags/tagID`


<!-- - **request_method** -  is a callable (like app.get, app.post, foo_router.patch and so on.).  
- **service_url** - the path to the endpoint on another service (like "https://microservice1.example.com").  
- **service_path** - the path to the method in microservice (like "/v1/microservice/users").  
- **gateway_path** - is the path to bind gateway.  
For example, your gateway api is located here - *https://gateway.example.com* and the path to endpoint (**gateway_path**) - "/users" then the full way to this method will be - *https://gateway.example.com/users*
- **override_headers** - Boolean value allows you to return all the headlines that were created by microservice in gateway.  
- **query_params** - used to extract query parameters from endpoint and transmission to microservice
- **form_params** -  used to extract form model parameters from endpoint and transmission to microservice
- **param body_params** - used to extract body model from endpoint and transmission to microservice

⚠️ - **Be sure to transfer the name of the argument to the router, which is in the endpoint func!**  -->
