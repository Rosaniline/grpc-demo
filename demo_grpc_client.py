import grpc
from google.protobuf.empty_pb2 import Empty
from grpc._channel import _Rendezvous

from proto.person_pb2 import Person
from proto.service_pb2 import getPersonRequest
from proto.service_pb2_grpc import HuaFuStub


class Client(object):
    stub: HuaFuStub

    def __init__(self):
        channel = grpc.insecure_channel('localhost:50051')
        self.stub = HuaFuStub(channel)
        self.stub2 = HuaFuStub(channel)

    def get_person(self, id: int):
        try:
            res: Person = self.stub.GetPerson(getPersonRequest(id=id))
            print(res)

        except grpc.RpcError as e:
            print(e)

    def list(self):
        try:
            for ppl in self.stub.ListPeople(Empty(), timeout=2):
                print(ppl)

        except grpc.RpcError as e:
            print(e)

    def input_query(self):
        while True:
            i = input("\nEnter a id or 'q' to quit: \n")

            if i == "q":
                break
            try:
                user_id = int(i)
            except ValueError:
                continue

            yield getPersonRequest(id=user_id)

    def query(self):
        try:
            for p in self.stub.QueryPerson(self.input_query()):
                print(p)

        except grpc.RpcError as e:
            print(e)


if __name__ == '__main__':
    c = Client()

    # c.get_person(9487)

    # while True:
    #     c.get_person(int(input()))

    # c.list()

    c.query()


