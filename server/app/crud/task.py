from typing import Sequence

from fastapi import HTTPException
from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import joinedload, selectinload

from app.models import Task as TaskDBModel
from app.schemas.task import TaskCreateSchema, TaskUpdateSchema


async def get_task_from_db_by_id(db_session: AsyncSession, task_id: int) -> TaskDBModel:
    return (await db_session.scalars(select(TaskDBModel).where(TaskDBModel.id == task_id))).first()


async def change_task_status(db_session: AsyncSession, task_id: int, user_id: int, status: str = None,
                             abandon: bool = False) -> None:
    task = await get_task_from_db_by_id(db_session, task_id)
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")
    if task.user_id != user_id:
        raise HTTPException(status_code=409, detail="Another user took the task, or you didn’t take it")
    if abandon:
        task.user_id = None
        task.status = "NotCompleted"
    else:
        task.status = status
    await db_session.commit()


async def get_task(db_session: AsyncSession, task_id: int) -> TaskDBModel:
    task = (await db_session.scalars(select(TaskDBModel).options(joinedload(TaskDBModel.parent_task),
                                                                 selectinload(TaskDBModel.subtasks,
                                                                              recursion_depth=float("inf"))).where(
        TaskDBModel.id == task_id))).first()
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")
    return task


async def create_task(db_session: AsyncSession, task_data: TaskCreateSchema) -> TaskDBModel:
    new_task = TaskDBModel(**task_data.dict())
    db_session.add(new_task)
    await db_session.commit()
    await db_session.refresh(new_task)
    return new_task


async def delete_task(db_session: AsyncSession, task_id: int) -> dict:
    task = await get_task_from_db_by_id(db_session, task_id)
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")
    await db_session.delete(task)
    await db_session.commit()
    return {"status": 200, "message": f"Task {task_id} deleted successfully"}


async def update_task(db_session: AsyncSession, task_id: int, task_data: TaskUpdateSchema) -> TaskDBModel:
    task = await get_task_from_db_by_id(db_session, task_id)
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")
    for field, value in task_data.model_dump(exclude_unset=True).items():
        setattr(task, field, value)
    await db_session.commit()
    await db_session.refresh(task)
    return task


async def get_user_task(db_session: AsyncSession, user_id: int) -> Sequence[TaskDBModel]:
    tasks = (await db_session.scalars(select(TaskDBModel).where(TaskDBModel.user_id == user_id))).all()
    if not tasks:
        return []
    return tasks


async def take_task_(db_session: AsyncSession, task_id: int, user_id: int, ) -> dict:
    task = await get_task_from_db_by_id(db_session, task_id)
    if not task:
        raise HTTPException(status_code=404, detail=f"Task {task_id} not found")
    if task.user_id == user_id:
        raise HTTPException(status_code=418, detail="You have already taken this task")
    if task.user_id:
        raise HTTPException(status_code=409, detail="Task is being performed by another user")
    task.user_id = user_id
    await db_session.commit()
    return {"status": 200, "message": f"You took on the job task {task_id}"}


async def start_task_(db_session: AsyncSession, task_id: int, user_id: int, ) -> dict:
    task = (await db_session.scalars(select(TaskDBModel).options(joinedload(TaskDBModel.parent_task),
                                                                 selectinload(TaskDBModel.subtasks,
                                                                              recursion_depth=1)).where(
        TaskDBModel.id == task_id))).first()
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")
    if task.user_id != user_id:
        raise HTTPException(status_code=409, detail="Another user took the task, or you didn’t take it")
    if task.subtasks:
        for subtask in task.subtasks:
            if subtask.status != "Completed":
                raise HTTPException(status_code=406,
                                    detail=f"Before this task can be executed, the subtask {subtask.id} must be executed")
    task.status = "InProgress"
    await db_session.commit()
    return {"status": 200, "message": f"You have accepted a task {task_id} to complete"}


async def finish_task_(db_session: AsyncSession, task_id: int, user_id: int, ) -> dict:
    await change_task_status(db_session, task_id, user_id, status="Completed")
    return {"status": 200, "message": f"You have completed the task {task_id}"}


async def cancel_task_(db_session: AsyncSession, task_id: int, user_id: int, ) -> dict:
    await change_task_status(db_session, task_id, user_id, status="Canceled")
    return {"status": 200, "message": f"You canceled the task {task_id}"}


async def abandon_task_(db_session: AsyncSession, task_id: int, user_id: int, ) -> dict:
    await change_task_status(db_session, task_id, user_id, abandon=True)
    return {"status": 200, "message": f"You abandoned the task {task_id}"}
