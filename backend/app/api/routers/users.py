from typing import List

from fastapi import APIRouter

from app.api.dependencies.core import DBSessionDep
from app.crud.user import get_users_by_count_tasks
from app.schemas.user import UserReadWithTasks

api_users_router = APIRouter(
    prefix="/api/users",
    tags=["api_users"],
    responses={404: {"description": "Not found"}},
)


@api_users_router.get("/", response_model=List[UserReadWithTasks],)
async def get_users(db_session: DBSessionDep):
    """Get any users sorted by task's count"""
    users = await get_users_by_count_tasks(db_session)
    return users
