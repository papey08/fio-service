# fio-service

## Описание

Данный проект представляет собой приложение для сбора, обогащения и хранения 
ФИО пользователей.

## Структура проекта

```text
├── cmd
│   └── server
│       └── main.go // точка входа в приложение
│
├── config
│   └── .env // файл с конфигурациями
│
├── internal
│   ├── adapters // слой для работы с внешними ресурсами
│   │   ├── apis // пакет для запросов возраста, пола и национальности
│   │   └── publisher // пакет для отправки некорректных ФИО в FIO_FAILED
│   │
│   ├── app // слой бизнес-логики
│   │   ├── mocks
│   │   ├── valid // пакет для валидации ФИО
│   │   ├── app.go // реализация интерфейса приложения
│   │   ├── app_interface.go // интерфейс приложения
│   │   └── app_test.go
│   │
│   ├── model // слой сущностей (entities)
│   │   ├── errs.go
│   │   ├── filter.go // структура фильтра поиска
│   │   └── fio.go // структура ФИО
│   │
│   ├── ports // сетевой слой (infrastructure)
│   │   ├── consumer // пакет для получения ФИО из FIO
│   │   ├── graphql // graphql-сервер
│   │   └── rest // rest-сервер
│   │
│   └── repo // хранилище ФИО
│
├── migrations
│   └── fio_repo_init.sql // скрипт для конфигурации fio_repo
│
├── pkg
│   └── logger // пакет для логирования
│
├── README.md
├── go.mod
└── go.sum

```

## Бизнес-логика

ФИО читается из FIO (топик кафки), проверяется на валидность (отсутствие спец. 
символов, обязательно наличие фамилии и имени). Затем выполняются три запроса к 
внешнему API, которые по имени определяют возраст, пол и национальность. В 
структуре ФИО обновляются значения, и она отправляется в БД. Если ФИО 
некорректно, структура отправляется в топик FIO_FAILED. При добавлении ФИО 
через REST или GraphQL все поля должны быть заполнены сразу.

Через REST и GraphQL поддерживаются все CRUD-операции, также можно выбрать 
список пользователей с фильтрами по всем полям (фамилия, имя, отчество, 
возраст, пол, национальность) и пагинацией.

Также для ускоренного получения ФИО реализован кеш, в который на время попадают 
все новые или только что обновлённые ФИО.

## Используемые технологии

* go 1.21
* Apache Kafka
* PostgreSQL
* Redis
* Gin Web Framework
* GraphQL

## Запуск приложения

Самостоятельно сконфигурировать БД PostgreSQL, используя [файл миграции](https://github.com/papey08/fio-service/blob/master/migrations/fio_repo_init.sql),
заменить в файле [*.env*](https://github.com/papey08/fio-service/blob/master/config/.env)
конфигурационные данные на свои, после чего выполнить команды:

```shell
$ go mod download
$ go run cmd/server/main.go
```

## Формат REST-запросов

### Добавление пользователя

* Метод: `POST`
* Эндпоинт: `http://localhost:8080/api/v1/fio`
* Формат тела запроса:

```json
{
    "name": "Matvey", 
    "surname": "Popov", 
    "patronymic": "Romanovich",
    "age": 20,
    "gender": "male",
    "nation": "RU"
}
```

* Формат ответа:

```json
{
    "data": {
        "id": 79,
        "name": "Matvey",
        "surname": "Popov",
        "patronymic": "Romanovich",
        "age": 20,
        "gender": "male",
        "nation": "RU"
    },
    "error": null
}
```

### Получение пользователя

* Метод: `GET`
* Эндпоинт: `http://localhost:8080/api/v1/fio/:id`
* Формат ответа:

```json
{
    "data": {
        "id": 79,
        "name": "Matvey",
        "surname": "Popov",
        "patronymic": "Romanovich",
        "age": 20,
        "gender": "male",
        "nation": "RU"
    },
    "error": null
}
```

### Получение списка пользователей

* Метод: `GET`
* Эндпоинт: `http://localhost:8080/api/v1/fio`
* Формат запроса:

```json
{
    "offset": 0,
    "limit": 1,
    "by_name": true,
    "name": "Matvey",
    "by_surname": false,
    "surname": "",
    "by_patronymic": false,
    "patronymic": "",
    "by_age": true,
    "age": 20,
    "by_gender": false,
    "gender": "",
    "by_nation": false,
    "nation": ""
}
```

* Формат ответа:

```json
{
    "data": [
        {
            "id": 79,
            "name": "Matvey",
            "surname": "Popov",
            "patronymic": "Romanovich",
            "age": 20,
            "gender": "male",
            "nation": "RU"
        }
    ],
    "error": null
}
```

### Обновление пользователя

* Метод: `PUT`
* Эндпоинт: `http://localhost:8080/api/v1/fio/:id`
* Формат запроса:

```json
{
    "name": "Matvey",
    "surname": "Popov",
    "patronymic": "Romanovich",
    "age": 21,
    "gender": "male",
    "nation": "RU"
}
```

* Формат ответа:

```json
{
    "data": {
        "id": 79,
        "name": "Matvey",
        "surname": "Popov",
        "patronymic": "Romanovich",
        "age": 21,
        "gender": "male",
        "nation": "RU"
    },
    "error": null
}
```

### Удаление пользователя

* Метод: `DELETE`
* Эндпоинт: `http://localhost:8080/api/v1/fio/:id`
* Формат ответа:

```json
{
    "data": null,
    "error": null
}
```

## Формат GraphQL-запросов

***Все запросы выполняются по адресу `http://localhost:8081`***

### Добавление пользователя

* Формат запроса:

```text
mutation AddFio {
    addFio(
        name: "Matvey"
        surname: "Popov"
        patronymic: "Romanovich"
        age: 20
        gender: "male"
        nation: "RU"
    ) {
        id
        name
        surname
        patronymic
        age
        gender
        nation
    }
}
```

* Формат ответа:

```json
{
    "data": {
        "addFio": {
            "age": 20,
            "gender": "male",
            "id": 86,
            "name": "Matvey",
            "nation": "RU",
            "patronymic": "Romanovich",
            "surname": "Popov"
        }
    }
}
```

### Получение пользователя

* Формат запроса:

```text
query GetFioById {
    getFioById(id: 86) {
        id
        name
        surname
        patronymic
        age
        gender
        nation
    }
}
```

* Формат ответа:

```json
{
    "data": {
        "getFioById": {
            "age": 20,
            "gender": "male",
            "id": 86,
            "name": "Matvey",
            "nation": "RU",
            "patronymic": "Romanovich",
            "surname": "Popov"
        }
    }
}
```

### Получение списка пользователей

* Формат запроса:

```text
query GetFioByFilter {
    getFioByFilter(offset: 0, limit: 1, name: "Matvey", age: 20) {
        age
        gender
        id
        name
        nation
        patronymic
        surname
    }
}
```

* Формат ответа:

```json
{
    "data": {
        "getFioByFilter": [
            {
                "age": 20,
                "gender": "male",
                "id": 86,
                "name": "Matvey",
                "nation": "RU",
                "patronymic": "Romanovich",
                "surname": "Popov"
            }
        ]
    }
}
```

### Обновление пользователя

* Формат запроса:

```text
mutation UpdateFio {
    updateFio(
        gender: "male"
        nation: "RU"
        id: 86
        name: "Matvey"
        surname: "Popov"
        patronymic: "Romanovich"
        age: 21
    ) {
        age
        gender
        id
        name
        nation
        patronymic
        surname
    }
}

```

* Формат ответа:

```json
{
    "data": {
        "updateFio": {
            "age": 21,
            "gender": "male",
            "id": 86,
            "name": "Matvey",
            "nation": "RU",
            "patronymic": "Romanovich",
            "surname": "Popov"
        }
    }
}
```

### Удаление пользователя

* Формат запроса:

```text
mutation DeleteFio {
    deleteFio(id: 86)
}
```

* Формат ответа:

```json
{
    "data": {
        "deleteFio": true
    }
}
```
