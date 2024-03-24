# TaskTrackerPro

Backend REST API-приложения для создания и распределения задач.

### Стек и назначение:

Вся структура приложения и настройки реализованы для работы в контейнерах Docker

Как основа выбран FastApi, он позволяет быстро и удобно принимать запросы, имеет возможность выполнять запросы и
работать в асинхронном режиме, а также удобно реализованы интеграции с Pydentic, FastApiUsers и тесно связан с
аннотацией типов от Python.

Для валидации и сериализации данных выбран Pydantic

Для связи с базой данных выбрана ORM и Core от SQLAlchemy 2.0, а также для создания и управлении миграций выбран Alembic

Для управления пользователями и аутентификацией выбрана библиотека FastApiUsers

Все функции выполняются в асинхронном режиме для обеспечения высокой производительности приложения

### Начало работы:

1. Необходимо переименовать файл _.env.template_ в _.env.docker_ и дописать все значения переменных
2. Создать кластер контейнеров и запустить их с помощью команды `docker-compose up`
3. Далее нужно выполнить команду `docker-compose exec backend alembic upgrade head` для применения миграций

### Работа приложения:
Все функции и схемы описаны в документации от OpenApi по адресу _/api/docs_

### Тестирование:
C покрытием эндпоинтов:
docker-compose run --rm -e TEST=1 -e DATABASE_URL="postgresql+asyncpg://test_user:password@test-postgres:5432/test_db" backend pytest --cov=app/api/routers/

Без покрытия:
docker-compose run --rm -e TEST=1 -e DATABASE_URL="postgresql+asyncpg://test_user:password@test-postgres:5432/test_db" backend pytest
