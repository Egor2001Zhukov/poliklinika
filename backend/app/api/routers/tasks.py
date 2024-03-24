from typing import List

from fastapi import APIRouter, Depends

from app.api.dependencies.core import DBSessionDep
from app.api.dependencies.current_user import СurrentUserDep
from app.crud.task import create_task, delete_task, update_task, take_task_, start_task_, finish_task_, cancel_task_, \
    abandon_task_
from app.crud.task import get_task, get_user_task
from app.schemas.task import TaskAllSubtasksReadSchema, TaskUpdateSchema
from app.schemas.task import TaskCreateSchema, TaskReadSchema
from app.schemas.user import UserRead

api_tasks_router = APIRouter(
    prefix="/api/tasks",
    tags=["api_tasks"],
    responses={404: {"description": "Not found"}},
)


@api_tasks_router.get("/me", response_model=List[TaskReadSchema], )
async def task_my_list(db_session: DBSessionDep, user: СurrentUserDep):
    """Get any task details"""
    task = await get_user_task(db_session, user.id)
    return task


@api_tasks_router.get("/{task_id}", response_model=TaskAllSubtasksReadSchema, )
async def task_details(db_session: DBSessionDep, task_id: int, ):
    """Get any task details"""
    task = await get_task(db_session, task_id)
    return task


@api_tasks_router.post("/", response_model=TaskReadSchema, )
async def task_create(db_session: DBSessionDep, task_data: TaskCreateSchema, ):
    """Create task"""
    new_task = await create_task(db_session, task_data)
    return new_task


@api_tasks_router.delete("/{task_id}", response_model=dict, )
async def task_create(db_session: DBSessionDep, task_id: int, ):
    """Delete task"""
    message = await delete_task(db_session, task_id)
    return message


@api_tasks_router.patch("/{task_id}", response_model=TaskReadSchema, )
async def task_create(db_session: DBSessionDep, task_data: TaskUpdateSchema, task_id: int, ):
    """Update task"""
    task = await update_task(db_session, task_id, task_data)
    return task


@api_tasks_router.get("/{task_id}/take_task", response_model=dict, )
async def take_task(db_session: DBSessionDep, task_id: int, user: СurrentUserDep, ):
    """Take task by current user"""
    message = await take_task_(db_session, task_id, user.id)
    return message


@api_tasks_router.get("/{task_id}/start_task", response_model=dict, )
async def start_task(db_session: DBSessionDep, task_id: int, user: СurrentUserDep, ):
    """Take task by current user"""
    message = await start_task_(db_session, task_id, user.id)
    return message


@api_tasks_router.get("/{task_id}/finish_task", response_model=dict, )
async def finish_task(db_session: DBSessionDep, task_id: int, user: СurrentUserDep, ):
    """Take task by current user"""
    message = await finish_task_(db_session, task_id, user.id)
    return message


@api_tasks_router.get("/{task_id}/cancel_task", response_model=dict, )
async def cancel_task(db_session: DBSessionDep, task_id: int, user: СurrentUserDep, ):
    """Take task by current user"""
    message = await cancel_task_(db_session, task_id, user.id)
    return message


@api_tasks_router.get("/{task_id}/abandon_task", response_model=dict, )
async def abandon_task(db_session: DBSessionDep, task_id: int, user: СurrentUserDep, ):
    """Take task by current user"""
    message = await abandon_task_(db_session, task_id, user.id)
    return message
