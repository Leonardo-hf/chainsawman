from flask import Config

from util.minio_helper import MinioHelper


class Client:
    config: Config

    @classmethod
    def init(cls, config: Config) -> None:
        config.setdefault('MINIO_CLIENT', MinioHelper(config['MINIO']))
        cls.config = config

    @classmethod
    def get_oss(cls) -> MinioHelper:
        return cls.config.get("MINIO_CLIENT")
