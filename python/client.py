import asyncio
from gen.interview.storage import StorageStub, ReadRequest

from grpclib.client import Channel


async def main():
    channel = Channel(host="127.0.0.1", port=50051)
    service = StorageStub(channel)
    response = await service.read(key="test.txt")
    print(response)

    # don't forget to close the channel when done!
    channel.close()


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())