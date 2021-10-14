import asyncio
from gen.interview.storage import StorageBase, ReadResponse
import grpclib

class StorageService(StorageBase):
  async def read(self, key: str) -> "ReadResponse":
    raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)
  

async def main():
  server = grpclib.server.Server([StorageService()])
  await server.start("127.0.0.1", 50051)
  await server.wait_closed()


if __name__ == '__main__':
  loop = asyncio.get_event_loop()
  loop.run_until_complete(main())