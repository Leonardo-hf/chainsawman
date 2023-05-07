import json

from CSMClient import CSMClient

TOP = 'import'
GROUP = 'import_consumers'
client = None


class UpdateBody:
    def __init__(self, graphId, edges):
        self.graphId = graphId
        self.edges = edges

    def toJson(self):
        temp = self.__dict__
        return json.dumps(temp)


class Message:
    def __init__(self, Opt, Entity, body):
        self.Opt = Opt
        self.Entity = Entity
        self.body = body


def initClient():
    global client
    client = CSMClient()
    client.createGroup(TOP, GROUP)


def upload(body):
    if client is None:
        initClient()
    msg=Message(1,1,body)
    client.send(TOP,msg.__dict__)


if __name__ == '__main__':
    ha = UpdateBody(1, [[1, 2], [3, 4]])
    print(ha.toJson())
