import * as grpc from "@grpc/grpc-js";
import { IStorageServer, StorageService } from './gen/storage_grpc_pb';
import { ReadRequest, ReadResponse } from './gen/storage_pb';

class StorageServer implements IStorageServer {
	[name: string]: grpc.UntypedHandleCall;

	read(call: grpc.ServerUnaryCall<ReadRequest, ReadResponse>, callback: grpc.sendUnaryData<ReadResponse>): void {
		callback({ code: grpc.status.UNIMPLEMENTED, message: "UNIMPLEMENTED"})
	}
}

function main() {
	const server = new grpc.Server();

	server.addService(StorageService, new StorageServer())

	const port = process.env.PORT || "50051"

	server.bindAsync(`localhost:${port}`, grpc.ServerCredentials.createInsecure(), (error) => {
		console.log(`Listening on ${port}`);
		server.start();
	});
}

main();

