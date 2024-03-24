import pytest
from httpx import AsyncClient


@pytest.mark.asyncio
async def test_get_tasks_me(register_and_login_user, client: AsyncClient):
    response = await client.get(
        "/api/tasks/me",
        headers={"Cookie": "fastapiusersauth=" + client.cookies["fastapiusersauth"], }, )
    assert response.status_code == 200


@pytest.mark.asyncio
async def test_create_task(register_and_login_user, client: AsyncClient):
    task = {
        "name": "test_task",
        "deadline": "2023-11-20T14:24:32.584Z",
        "status": "NotCompleted"
    }
    response = await client.post(
        "/api/tasks/", json=task,
        headers={"Cookie": "fastapiusersauth=" + client.cookies["fastapiusersauth"], }, )
    assert response.status_code == 200
