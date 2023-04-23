# serialization-test

## Сборка и запуск

Для запуска необходимо иметь установленный `docker` и `docker compose`.

Запуск осуществляется следующей командой:

```bash
make run
```

## Функционал и тестирование

Для пользователя доступен один UDP порт - 2000, а также следующий набор команд:

- `get_result gob` - тест "нативной" сериализации/десериализации `gob`

- `get_result json` - тест сериализации/десериализации `json`

- `get_result xml` - тест сериализации/десериализации `xml`

- `get_result protobuf` - тест сериализации/десериализации `protobuf`

- `get_result yaml` - тест сериализации/десериализации `yaml`

- `get_result message_pack` - тест сериализации/десериализации `message_pack`

- `get_result all` - тест сериализации всех вышеперечисленных форматов. Все запросы делаются в разных горутинах.

Все тесты проводятся 1000 раз и берется среднее значение.

Пример тестирования:

```bash
nc -u 127.0.0.1 2000
get_result gob
gob - 664 - 11.526µs - 37.22µs
get_result json
json - 762 - 2.196µs - 10.365µs
get_result xml
xml - 1209 - 9.686µs - 40.337µs
get_result protobuf
protobuf - 371 - 7.5µs - 6.242µs
get_result yaml
yaml - 927 - 44.186µs - 60.478µs
get_result message_pack
message_pack - 715 - 12.411µs - 8.426µs
get_result all
message_pack - 715 - 2.751µs - 3.952µs
protobuf - 371 - 1.389µs - 2.382µs
json - 762 - 1.899µs - 9.454µs
gob - 664 - 7.247µs - 20.284µs
xml - 1209 - 8.991µs - 43.761µs
yaml - 927 - 55.263µs - 58.31µs
```
