import redis


class CSMClient:

    def __init__(self, host='localhost', port=6379):
        pool = redis.ConnectionPool(
            host=host, port=port, decode_responses=True)  # decode_responses=True 这样写存的数据是字符串格式
        self.rdc = redis.StrictRedis(connection_pool=pool)

    def createGroup(self, stream_name, group_name):
        if self.rdc.exists(stream_name):
            self.rdc.delete(stream_name)
        self.rdc.xgroup_create(stream_name, group_name, id=0, mkstream=True)

    def send(self, stream_name, values):
        print(values)
        self.rdc.xadd(stream_name, values)


if __name__ == '__main__':
    client = CSMClient()
    client.createGroup("import", "import_consumers")
