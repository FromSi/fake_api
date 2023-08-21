# Fake Api Examples
Примеры данных для работы с сервером.

* `fake-api.postman_collection_v2.1.json` — файл можно импортировать в `Postman` (все API с настроенными переменными)
* `example_payload.json` — данные для создания `JWT` по API `POST ../login` с `Content-Type: application/json`
* `example_payload.xml` — данные для создания `JWT` по API `POST ../login` с `Content-Type: application/xml`

## Список типов

### Числовые
* `uint8` — Беззнаковое 8-битное целое число
* `uint16` — Беззнаковое 16-битное целое число
* `uint32` — Беззнаковое 32-битное целое число
* `int8` — Знаковое 8-битное целое число
* `int16` — Знаковое 16-битное целое число
* `int32` — Знаковое 32-битное целое число

### Вещщественные
* `float` — Плавающее число (число с плавающей точкой). Это числовой тип данных, который может представлять дробные значения с плавающей точкой
* `latitude` — Широта. Обычно представляется числом с плавающей точкой и указывает на географическую широту местоположения. 0 обозначает экватор
* `longitude` — Долгота. Также представляется числом с плавающей точкой и указывает на географическую долготу местоположения. 0 обозначает меридиан Гринвича

### Булевойские
* `boolean` — логический тип данных, который может принимать два возможных значения: true (истина) или false (ложь). Может быть использовано для представления состояний включения/выключения, наличия/отсутствия и т.д.

### Исключение (ссылка на изображение)
`min` и `max` у него как `weight` и `height`
* `string_image_url` — ссылка на изображение

### Строки
* `string_name` — для имени
* `string_first_name` — для имени (первого имени)
* `string_middle_name` — для отчества
* `string_last_name` — для фамилии
* `string_gender` — для пола
* `string_ssn` — для номера социального страхования
* `string_hobby` — для хобби
* `string_email` — для адреса электронной почты
* `string_username` — для имени пользователя
* `string_country` — для названия страны
* `string_country_abr` — для аббревиатуры названия страны
* `string_city` — для названия города
* `string_state` — для названия региона (штата)
* `string_street` — для названия улицы
* `string_street_name` — для названия улицы (часть)
* `string_street_number` — для номера дома на улице
* `string_street_prefix` — для префикса улицы (например, "ул.")
* `string_street_suffix` — для суффикса улицы (например, "проезд")
* `string_zip` — для почтового индекса
* `string_gametag` — для игрового тега
* `string_beer_alcohol` — для содержания алкоголя в пиве
* `string_beer_blg` — для уровня сахаристости пива (BLG)
* `string_beer_hop` — для вида хмеля в пиве
* `string_beer_ibu` — для индекса горечи пива (IBU)
* `string_beer_malt` — для вида солода в пиве
* `string_beer_name` — для названия пива
* `string_beer_style` — для стиля пива
* `string_beer_yeast` — для вида дрожжей в пиве
* `string_noun` — для существительного
* `string_noun_common` — для общего существительного
* `string_noun_concrete` — для конкретного существительного
* `string_noun_abstract` — для абстрактного существительного
* `string_noun_collective_people` — для коллективного существительного (люди)
* `string_noun_collective_animal` — для коллективного существительного (животные)
* `string_noun_collective_thing` — для коллективного существительного (вещи)
* `string_noun_countable` — для исчисляемого существительного
* `string_noun_uncountable` — для неисчисляемого существительного
* `string_verb` — для глагола
* `string_verb_action` — для глагола действия
* `string_verb_linking` — для связывающего глагола
* `string_verb_helping` — для вспомогательного глагола
* `string_adverb` — для наречия
* `string_adverb_manner` — для наречия образа
* `string_adverb_degree` — для наречия степени
* `string_adverb_place` — для наречия места
* `string_adverb_time_definite` — для наречия определенного времени
* `string_adverb_time_indefinite` — для наречия неопределенного времени
* `string_adverb_frequency_definite` — для наречия определенной частоты
* `string_adverb_frequency_indefinite` — для наречия неопределенной частоты
* `string_preposition` — для предлога
* `string_preposition_simple` — для простого предлога
* `string_preposition_double` — для двойного предлога
* `string_preposition_compound` — для составного предлога
* `string_adjective` — для прилагательного
* `string_adjective_descriptive` — для описательного прилагательного
* `string_adjective_quantitative` — для количественного прилагательного
* `string_adjective_proper` — для нарицательного прилагательного
* `string_adjective_demonstrative` — для указательного прилагательного
* `string_adjective_possessive` — для притяжательного прилагательного
* `string_adjective_interrogative` — для вопросительного прилагательного
* `string_adjective_indefinite` — для неопределенного прилагательного
* `string_pronoun` — для местоимения
* `string_pronoun_personal` — для личного местоимения
* `string_pronoun_object` — для объектного местоимения
* `string_pronoun_possessive` — для притяжательного местоимения
* `string_pronoun_reflective` — для возвратного местоимения
* `string_pronoun_demonstrative` — для указательного местоимения
* `string_pronoun_interrogative` — для вопросительного местоимения
* `string_pronoun_relative` — для относительного местоимения
* `string_connective` — для союза/связки
* `string_connective_time` — для союза времени
* `string_connective_comparative` — для союза сравнения
* `string_connective_complaint` — для союза жалобы
* `string_connective_listing` — для союза перечисления
* `string_connective_casual` — для неформального союза
* `string_connective_examplify` — для союза примера
* `string_question` — для вопроса
* `string_quote` — для цитаты
* `string_phrase` — для фразы
* `string_word` — для слова
* `string_sentence` — для предложения
* `string_url` — для URL-адреса
* `string_uuid` — для UUID
* `string_color` — для цвета
* `string_hex_color` — для цвета в шестнадцатеричной нотации
* `string_safe_color` — для безопасного цвета
* `string_phone` — для номера телефона
* `string_phone_formatted` — для отформатированного номера телефона
* `string_credit_card` — для номера кредитной карты
* `string_currency` — для валюты
* `string_bitcoin_address` — для адреса Bitcoin
* `string_emoji` — для эмодзи
* `string_ipv4` — для IPv4-адреса
* `string_ipv6` — для IPv6-адреса
* `string_date` — для даты
* `string_date_time` — для даты и времени
* `string_time` — для времени
* `string_car_maker` — для производителя автомобиля
* `string_car_model` — для модели автомобиля
* `string_car_type` — для типа автомобиля
* `string_car_fuel_type` — для типа топлива автомобиля
* `string_car_transmission_type` — для типа трансмиссии автомобиля
* `string_fruit` — для фрукта
* `string_vegetable` — для овоща
* `string_breakfast` — для завтрака
* `string_lunch` — для обеда
* `string_dinner` — для ужина
* `string_snack` — для закуски
* `string_dessert` — для десерта
* `string_flip_a_coin` — для подбрасывания монеты
