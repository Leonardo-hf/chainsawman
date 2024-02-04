from flask import Config

from util import VulAPI, MinioHelper


class Client:
    config: Config

    @classmethod
    def init(cls, config: Config) -> None:
        config.setdefault('MINIO_CLIENT', MinioHelper(config['MINIO']))
        config.setdefault('OSV_CLIENT', VulAPI(config['OSV_API']))
        cls.config = config

    @classmethod
    def get_oss(cls) -> MinioHelper:
        return cls.config.get("MINIO_CLIENT")

    @classmethod
    def get_osv(cls) -> VulAPI:
        return cls.config.get("OSV_CLIENT")
