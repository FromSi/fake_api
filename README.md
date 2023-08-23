# Fake Api

Простой инструмент для предоставления HTTP-API. Удобен, когда изучаешь новые технологии, а для этого нужен сервер с HTTP-API. [Документация HTTP-API](https://documenter.getpostman.com/view/2849977/2s9Y5VTiZb).

# Список endpoins
* `POST ../login` — регисрация полей **(получение JWT)**
* `GET ../show` — просмотр одного ресурса **(нужен JWT в Headers Authorization Bearer)**
* `GET ../list` — просмотр списка ресурсов **(нужен JWT в Headers Authorization Bearer)**
* `POST ../create` — создание ресурса **(нужен JWT в Headers Authorization Bearer)**
* `DELETE ../delete` — удаление ресурса **(нужен JWT в Headers Authorization Bearer)**
* `PATCH ../patch` — частичное обновление ресурса **(нужен JWT в Headers Authorization Bearer)**
* `PUT ../put` — полное обновление ресурса **(нужен JWT в Headers Authorization Bearer)**

# Demo (CLI)

`application/json`

```
curl --location 'https://fake-api.fromsi.net/show' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaWVsZHMiOlt7InR5cGUiOiJ1aW50MzIiLCJuYW1lIjoiaWQiLCJyZXF1aXJlZCI6dHJ1ZSwibWF4Ijo0Mjk0OTY3Mjk1LCJtaW4iOjB9LHsidHlwZSI6InN0cmluZ19maXJzdF9uYW1lIiwibmFtZSI6ImZpcnN0X25hbWUiLCJyZXF1aXJlZCI6dHJ1ZSwibWF4IjoyNTUsIm1pbiI6MX0seyJ0eXBlIjoic3RyaW5nX21pZGRsZV9uYW1lIiwibmFtZSI6Im1pZGRsZV9uYW1lIiwicmVxdWlyZWQiOnRydWUsIm1heCI6MjU1LCJtaW4iOjF9LHsidHlwZSI6InN0cmluZ19sYXN0X25hbWUiLCJuYW1lIjoibGFzdF9uYW1lIiwicmVxdWlyZWQiOnRydWUsIm1heCI6MjU1LCJtaW4iOjF9LHsidHlwZSI6InN0cmluZ19nZW5kZXIiLCJuYW1lIjoiZ2VuZGVyIiwicmVxdWlyZWQiOnRydWUsIm1heCI6MjU1LCJtaW4iOjF9LHsidHlwZSI6InN0cmluZ19zc24iLCJuYW1lIjoic3NuIiwicmVxdWlyZWQiOnRydWUsIm1heCI6MjU1LCJtaW4iOjF9LHsidHlwZSI6InN0cmluZ19ob2JieSIsIm5hbWUiOiJob2JieSIsInJlcXVpcmVkIjp0cnVlLCJtYXgiOjI1NSwibWluIjoxfSx7InR5cGUiOiJzdHJpbmdfZW1haWwiLCJuYW1lIjoiZW1haWwiLCJyZXF1aXJlZCI6dHJ1ZSwibWF4IjoyNTUsIm1pbiI6NX0seyJ0eXBlIjoic3RyaW5nX2NvdW50cnkiLCJuYW1lIjoiY291bnRyeSIsInJlcXVpcmVkIjp0cnVlLCJtYXgiOjI1NSwibWluIjoxfSx7InR5cGUiOiJzdHJpbmdfdXNlcm5hbWUiLCJuYW1lIjoidXNlcm5hbWUiLCJyZXF1aXJlZCI6dHJ1ZSwibWF4IjoyNTUsIm1pbiI6MX0seyJ0eXBlIjoic3RyaW5nX2ltYWdlX3VybCIsIm5hbWUiOiJhdmF0YXJfdXJsIiwicmVxdWlyZWQiOmZhbHNlLCJtYXgiOjE1MCwibWluIjoxNTB9LHsidHlwZSI6InN0cmluZ19pbWFnZV91cmwiLCJuYW1lIjoiYmFja2dyb3VuZF91cmwiLCJyZXF1aXJlZCI6ZmFsc2UsIm1heCI6NjAwLCJtaW4iOjIwMH0seyJ0eXBlIjoic3RyaW5nX2RhdGVfdGltZSIsIm5hbWUiOiJjcmVhdGVkX2F0IiwicmVxdWlyZWQiOmZhbHNlLCJtYXgiOjE5LCJtaW4iOjE5fV19.9Dyyw2qAp0HybzEJbhhyds83soTYAixFwy90rJyj0Hk'
```

`application/xml`

```
curl --location 'https://fake-api.fromsi.net/show' \
--header 'Content-Type: application/xml' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaWVsZHMiOlt7InR5cGUiOiJ1aW50MzIiLCJuYW1lIjoiaWQiLCJyZXF1aXJlZCI6dHJ1ZSwibWF4Ijo0Mjk0OTY3Mjk1LCJtaW4iOjB9LHsidHlwZSI6InN0cmluZ19maXJzdF9uYW1lIiwibmFtZSI6ImZpcnN0X25hbWUiLCJyZXF1aXJlZCI6dHJ1ZSwibWF4IjoyNTUsIm1pbiI6MX0seyJ0eXBlIjoic3RyaW5nX21pZGRsZV9uYW1lIiwibmFtZSI6Im1pZGRsZV9uYW1lIiwicmVxdWlyZWQiOnRydWUsIm1heCI6MjU1LCJtaW4iOjF9LHsidHlwZSI6InN0cmluZ19sYXN0X25hbWUiLCJuYW1lIjoibGFzdF9uYW1lIiwicmVxdWlyZWQiOnRydWUsIm1heCI6MjU1LCJtaW4iOjF9LHsidHlwZSI6InN0cmluZ19nZW5kZXIiLCJuYW1lIjoiZ2VuZGVyIiwicmVxdWlyZWQiOnRydWUsIm1heCI6MjU1LCJtaW4iOjF9LHsidHlwZSI6InN0cmluZ19zc24iLCJuYW1lIjoic3NuIiwicmVxdWlyZWQiOnRydWUsIm1heCI6MjU1LCJtaW4iOjF9LHsidHlwZSI6InN0cmluZ19ob2JieSIsIm5hbWUiOiJob2JieSIsInJlcXVpcmVkIjp0cnVlLCJtYXgiOjI1NSwibWluIjoxfSx7InR5cGUiOiJzdHJpbmdfZW1haWwiLCJuYW1lIjoiZW1haWwiLCJyZXF1aXJlZCI6dHJ1ZSwibWF4IjoyNTUsIm1pbiI6NX0seyJ0eXBlIjoic3RyaW5nX2NvdW50cnkiLCJuYW1lIjoiY291bnRyeSIsInJlcXVpcmVkIjp0cnVlLCJtYXgiOjI1NSwibWluIjoxfSx7InR5cGUiOiJzdHJpbmdfdXNlcm5hbWUiLCJuYW1lIjoidXNlcm5hbWUiLCJyZXF1aXJlZCI6dHJ1ZSwibWF4IjoyNTUsIm1pbiI6MX0seyJ0eXBlIjoic3RyaW5nX2ltYWdlX3VybCIsIm5hbWUiOiJhdmF0YXJfdXJsIiwicmVxdWlyZWQiOmZhbHNlLCJtYXgiOjE1MCwibWluIjoxNTB9LHsidHlwZSI6InN0cmluZ19pbWFnZV91cmwiLCJuYW1lIjoiYmFja2dyb3VuZF91cmwiLCJyZXF1aXJlZCI6ZmFsc2UsIm1heCI6NjAwLCJtaW4iOjIwMH0seyJ0eXBlIjoic3RyaW5nX2RhdGVfdGltZSIsIm5hbWUiOiJjcmVhdGVkX2F0IiwicmVxdWlyZWQiOmZhbHNlLCJtYXgiOjE5LCJtaW4iOjE5fV19.9Dyyw2qAp0HybzEJbhhyds83soTYAixFwy90rJyj0Hk'
```

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
