# Fake Api

Инструмент для предоставления HTTP-API. Удобен, когда изучаешь новые технологии, а для этого нужен сервер с HTTP-API.

# Список endpoins
* `POST ../login` — регисрация полей **(получение JWT)**
* `GET ../show` — просмотр одного ресурса **(нужен JWT в Headers Authorization Bearer)**
* `GET ../list` — просмотр списка ресурсов **(нужен JWT в Headers Authorization Bearer)**
* `POST ../create` — создание ресурса **(нужен JWT в Headers Authorization Bearer)**
* `DELETE ../delete` — удаление ресурса **(нужен JWT в Headers Authorization Bearer)**
* `PATCH ../patch` — частичное обновление ресурса **(нужен JWT в Headers Authorization Bearer)**
* `PUT ../put` — полное обновление ресурса **(нужен JWT в Headers Authorization Bearer)**

# Принцип работы
1. Импортировать в Postman файл [fake-api.postman_collection_v2.1.json](https://github.com/FromSi/fake_api/blob/master/examples/fake-api.postman_collection_v2.1.json)
2. Создать `JWT` или использовать переменные Postman, в котором будут описанны поля сущностей, через API `POST ../login`
   - Пример `request body` с `Content-Type` [`application/json`](https://github.com/FromSi/fake_api/blob/master/examples/example_payload.json)
   - Пример `request body` с `Content-Type` [`application/xml`](https://github.com/FromSi/fake_api/blob/master/examples/example_payload.xml)
3. `JWT` нужно вставить в `Headers Authorization Bearer` в HTTP запросе

# Enviromens
|Key                            |Default  |Description       |
|:------------------------------|:--------|:-----------------|
|`FAKE_API_HOST`                |`0.0.0.0`|Хост IPv4         |
|`FAKE_API_PORT`                |`8080`   |Порт TCP          |
|`FAKE_API_MAX_FIELDS_IN_OBJECT`|`50`     |Макс. кол-во полей|
