from typing import Sequence

from fastapi import Depends
from fastapi_users_db_sqlalchemy import SQLAlchemyUserDatabase
from sqlalchemy import select, func
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import joinedload

from app.database import get_db_session
from app.models import Task as TaskDBModel
from app.models import User as UserDBModel


async def get_user_db(session: AsyncSession = Depends(get_db_session)):
    yield SQLAlchemyUserDatabase(session, UserDBModel)


async def get_users_by_count_tasks(session: AsyncSession = Depends(get_db_session)) -> Sequence[UserDBModel]:
    users_sorted_by_count_tasks = (await (
        session.scalars(select(UserDBModel)
                        .options(joinedload(UserDBModel.tasks))
                        .join(TaskDBModel)
                        .group_by(UserDBModel.id)
                        .order_by(func.count(TaskDBModel.id).desc())  # Сортируем по убыванию количества задач
                        ))).unique().all()
    return users_sorted_by_count_tasks
