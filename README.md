# ez-server

## Запуск

```bash
docker-compose up --build
```

## cURL

Список задач
```
curl --location 'http://localhost:8080/task/list'
```

Добавить задачу
```
curl --location 'localhost:8080/task' \
--header 'Content-Type: application/json' \
--data '{"title":"Learn Go","desc":"Write a REST server"}'
```

Получить задачу
```
curl --location 'localhost:8080/task/1'
```

Обновить задачу
```
curl --location --request PUT 'localhost:8080/task/1' \
--header 'Content-Type: application/json' \
--data '{"title":"PULL","desc":"Write a REST server"}'
```

Удалить задачу
```
curl --location --request DELETE 'localhost:8080/task/1'
```