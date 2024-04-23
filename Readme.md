# Distributed Calculator

Проект Distributed Calculator - веб-приложение для создания и распределенного подсчета арифметических выражений. Творческое задание на курсе [Лицея Академии Яндекса](https://lyceum.yandex.ru/go)

По всем вопросам можно написать автору проекта - [Kuzmin Dmitry](https://t.me/kuzmindev)

**Стэк:** Golang + React JS

## Content
- [How Run Project](#how-run-project)
- [API Documentation](#api-documentation)
- [Use Cases](#use-cases)

# How Run Project

Make file contains such commands to start a project:
- `make docker-start` - starting a project (http and grpc servers)
- `make test-project` - running all test for project (for some packages there are no test)


# API Documentation

Server HTTP requests:
- POST /new_expression
- POST /new_user
- GET  /get_expression
- GET  /list_of_expressions

# Use Cases    
When you start a project, you can type these commands to create new user, create new expression, get expression.

**Example** `/new_user`:
```
curl -X POST -H "Content-Type: application/json" --data-binary "{\"login\":\"developer\", \"password\":\"password\"}" http://localhost:8080/new_user
```

**Example** `/new_expression`:
```
curl -X POST  -H "Content-Type: application/json" --data-binary "{\"value\":\"1 * 10 + 2 * 3\", \"user_id\":5}" http://localhost:8080/new_expression
```

**Example** `/get_expression`
```
curl  -X GET "http://localhost:8080/get_expression?expression_id=35"
```


## Полная схема проекта

<!--![Project Schema](./docs/Distributed%20Calculator%20Schema.png) -->
<img src="./docs/Scheme Project.png" alt="schema" style="margin: 0 auto; width:600px;"/>
