import datetime
from typing import List

from sqlalchemy import String, ForeignKey, DateTime, func
from sqlalchemy.orm import mapped_column, Mapped, relationship

from . import Base


class Task(Base):
    __tablename__ = "task"

    id: Mapped[int] = mapped_column(primary_key=True, index=True)
    name: Mapped[str] = mapped_column(String(30))
    parent_task_id: Mapped[int] = mapped_column(ForeignKey("task.id"), nullable=True)
    user_id: Mapped[int] = mapped_column(ForeignKey("user.id"), nullable=True)
    deadline: Mapped[datetime.datetime] = mapped_column(DateTime(timezone=True), server_default=func.now())
    status: Mapped[str] = mapped_column(String(30))
    parent_task: Mapped["Task"] = relationship("Task", remote_side=[id], back_populates="subtasks",
                                               foreign_keys=[parent_task_id])
    subtasks: Mapped[List["Task"]] = relationship("Task", back_populates="parent_task", foreign_keys=[parent_task_id],
                                                  cascade="all, delete-orphan")
    user: Mapped["User"] = relationship(back_populates="tasks")
