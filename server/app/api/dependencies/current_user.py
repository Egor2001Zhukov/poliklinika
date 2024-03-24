from typing import Annotated

from fastapi import Depends

from app.schemas.user import UserRead
from app.utils.user import fastapi_users

СurrentUserDep = Annotated[UserRead, Depends(fastapi_users.current_user())]
