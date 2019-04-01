import grpc
import sequence_pb2_grpc as client
from google.protobuf import empty_pb2

channel = grpc.insecure_channel('service.127.0.0.1.xip.io:443')
stub = client.SequenceStub(channel)
resp = stub.Increment(empty_pb2.Empty())

print("""
Increment:
  ID:    {}
  Value: {}
""".format(resp.id, resp.value))
