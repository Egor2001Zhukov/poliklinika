from pydantic_settings import BaseSettings


class Settings(BaseSettings):
    database_url: str
    echo_sql: bool = True
    test: bool = False
    project_name: str
    oauth_token_secret: str
    debug_logs: bool = True


settings = Settings()  # type: ignore
