from flask import Config, current_app

from util.minio_helper import MinioHelper


class Client:
    config: Config

    @staticmethod
    def init(config: Config) -> None:
        config = config
        current_app.config.setdefault('MINIO_CLIENT', MinioHelper(config['MINIO']))

    @staticmethod
    def get_oss() -> MinioHelper:
        return Client.config.get("MINIO_CLIENT")
