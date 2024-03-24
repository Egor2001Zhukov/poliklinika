import pytest
from httpx import AsyncClient
from sqlalchemy.ext.asyncio import AsyncSession


@pytest.mark.asyncio
async def test_get_users(client):
    response = await client.get("/api/users/")
    assert response.status_code == 200


@pytest.mark.asyncio
async def test_get_current_user(register_and_login_user, client: AsyncClient):
    response = await client.get(
        "users/me",
        headers={"Cookie": "fastapiusersauth=" + client.cookies["fastapiusersauth"],},)
    assert response.json().get('email') == 'test@gmail.com'
    assert response.status_code == 200
