# Модуль для валидации приходящих данных с помощью pydantic
import datetime
from enum import Enum
from typing import Optional, List

from pydantic import BaseModel, ConfigDict


class TaskStatus(str, Enum):
    not_completed = "NotCompleted"
    in_progress = "InProgress"
    completed = "Completed"
    cancelled = "Canceled"


class TaskReadSchema(BaseModel):
    model_config = ConfigDict(from_attributes=True)

    id: int
    name: str
    parent_task_id: Optional[int] = None
    user_id: Optional[int] = None
    deadline: datetime.datetime
    status: TaskStatus


class TaskAllSubtasksReadSchema(BaseModel):
    model_config = ConfigDict(from_attributes=True)

    id: int
    name: str
    parent_task_id: Optional[int] = None
    user_id: Optional[int] = None
    deadline: datetime.datetime
    status: TaskStatus
    subtasks: List["TaskAllSubtasksReadSchema"] = []


class TaskCreateSchema(BaseModel):
    model_config = ConfigDict(from_attributes=True)

    name: str
    parent_task_id: Optional[int] = None
    user_id: Optional[int] = None
    deadline: datetime.datetime
    status: TaskStatus = "NotCompleted"


class TaskUpdateSchema(BaseModel):
    name: Optional[str] = None
    parent_task_id: Optional[int] = None
    user_id: Optional[int] = None
    deadline: Optional[datetime.datetime] = None
    status: Optional[TaskStatus] = None
