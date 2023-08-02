from minio import Minio


class MinioHelper:
    def __init__(self, cfg):
        self._client = Minio(
            cfg['ENDPOINT'],
            access_key=cfg['ACCESS_KEY_ID'],
            secret_key=cfg['SECRET_ACCESS_KEY'],
            secure=cfg['USE_SSL']
        )
        self._bucket = cfg['BUCKET']

    def fetch(self, name):
        response = self._client.get_object(self._bucket, name)
        data = response.data
        response.close()
        response.release_conn()
        return data
