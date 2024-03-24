from typing import Optional, List

from fastapi_users import schemas

from app.schemas.task import TaskReadSchema


class UserReadWithTasks(schemas.BaseUser[int]):
    first_name: str
    last_name: str
    position: str
    tasks: List[TaskReadSchema] = []


class UserRead(schemas.BaseUser[int]):
    first_name: str
    last_name: str
    position: str


class UserCreate(schemas.BaseUserCreate):
    first_name: str
    last_name: str
    position: str


class UserUpdate(schemas.BaseUserUpdate):
    first_name: Optional[str] = None
    last_name: Optional[str] = None
    position: Optional[str] = None
