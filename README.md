# Error by level

`error_by_level` - это пакет, позволяющий скрывать системные ошибки от клиентов приложения.

## Описание
При использовании пакета в случае возникновения ошибки в основном приложении для формирования ошибки клиенту 
(для web-сервисов, например, при формировании HTTP ответа) ошибка передаётся в метод `Handle()`.

Если переданная ошибка **на любом уровне вложенности** содержит ошибку, не реализующую интерфейс `AppError`, 
то, во избежании просачивания наружу технических особенностей работы приложения, такая ошибка считается системной.

## Использование собественных ошибок
Пакет уже содержит тип `ProcessError`, который может быть использован для создания ошибок, 
безопасных для передачи клиенту.

Кроме этого, `error_by_level` будет считать безопасными все ошибки, реализующие интерфейс `AppError`.
