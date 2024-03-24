from typing import List

from fastapi_users_db_sqlalchemy import SQLAlchemyBaseUserTable
from sqlalchemy import String, Boolean
from sqlalchemy.orm import Mapped, mapped_column, relationship

from . import Base


class User(SQLAlchemyBaseUserTable[int], Base):
    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True, index=True)
    first_name: Mapped[str] = mapped_column(
        String(length=50), nullable=False)
    last_name: Mapped[str] = mapped_column(
        String(length=50), nullable=False)
    email: Mapped[str] = mapped_column(
        String(length=320), unique=True, index=True, nullable=False)
    hashed_password: Mapped[str] = mapped_column(
        String(length=1024), nullable=False)
    position: Mapped[str] = mapped_column(
        String(length=150), nullable=False)
    tasks: Mapped[List["Task"]] = relationship(back_populates="user")
    is_active: Mapped[bool] = mapped_column(Boolean, default=True, nullable=False)
    is_superuser: Mapped[bool] = mapped_column(
        Boolean, default=False, nullable=False
    )
    is_verified: Mapped[bool] = mapped_column(
        Boolean, default=False, nullable=False
    )
