from app.models import User
from app.utils.auth import auth_backend
from app.utils.manager import get_user_manager
from fastapi_users import FastAPIUsers

fastapi_users = FastAPIUsers[User, int](
    get_user_manager,
    [auth_backend],
)
