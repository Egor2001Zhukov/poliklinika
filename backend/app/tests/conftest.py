import asyncio
from contextlib import ExitStack
from typing import AsyncGenerator, Any

import pytest
import pytest_asyncio
from alembic.config import Config
from alembic.migration import MigrationContext
from alembic.operations import Operations
from alembic.script import ScriptDirectory
from asyncpg import Connection
from httpx import AsyncClient

from app.config import settings
from app.database import Base, get_db_session, sessionmanager
from app.main import app as actual_app


@pytest.fixture(autouse=True)
def app():
    with ExitStack():
        yield actual_app


@pytest_asyncio.fixture
async def client() -> AsyncGenerator[AsyncClient, Any]:
    async with AsyncClient(app=actual_app, base_url="http://test") as c:
        yield c


def run_migrations(connection: Connection):
    config = Config("app/alembic.ini")
    config.set_main_option("script_location", "app/alembic")
    config.set_main_option("sqlalchemy.url", settings.database_url)
    script = ScriptDirectory.from_config(config)

    def upgrade(rev, context):
        return script._upgrade_revs("head", rev)

    context = MigrationContext.configure(connection, opts={"target_metadata": Base.metadata, "fn": upgrade})

    with context.begin_transaction():
        with Operations.context(context):
            context.run_migrations()


@pytest.fixture(scope="session")
def event_loop():
    loop = asyncio.get_event_loop_policy().new_event_loop()
    yield loop
    loop.close()


@pytest_asyncio.fixture(scope="session", autouse=True)
async def setup_database():
    # Run alembic migrations on test DB
    async with sessionmanager.connect() as connection:
        await connection.run_sync(run_migrations)
    yield

    # Teardown
    await sessionmanager.close()


# Each test function is a clean slate
@pytest_asyncio.fixture(scope="function", autouse=True)
async def transactional_session():
    async with sessionmanager.session() as session:
        try:
            await session.begin()
            yield session
        finally:
            await session.rollback()  # Rolls back the outer transaction


@pytest_asyncio.fixture(scope="function")
async def db_session(transactional_session):
    yield transactional_session


@pytest_asyncio.fixture(scope="function", autouse=True)
async def session_override(app, db_session):
    async def get_db_session_override():
        yield db_session

    app.dependency_overrides[get_db_session] = get_db_session_override


# User register and login
@pytest_asyncio.fixture(scope="function")
async def register_and_login_user(client: AsyncClient):
    test_user_email = "test@gmail.com"
    test_user_password = "test_password"
    user_data = {
        "email": test_user_email,
        "password": test_user_password,
        "first_name": "test_first_name",
        "last_name": "test_last_name",
        "position": "test_position", }
    await client.post("/auth/register", json=user_data)

    response = await client.post(
        "/auth/jwt/login",
        data={"username": test_user_email, "password": test_user_password},
    )
    client.cookies.update(response.cookies)

    return client.cookies
