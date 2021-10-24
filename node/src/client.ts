import { StorageClient } from './gen/storage_grpc_pb';
import { ReadRequest } from './gen/storage_pb';
import * as grpc from '@grpc/grpc-js';

const port = process.env.PORT || "50051"

const client = new StorageClient(`localhost:${port}`, grpc.ChannelCredentials.createInsecure());

const req = new ReadRequest();
req.setKey("test.txt");

client.read(req, (err, resp) => {
	if (err) {
		console.error(err);
	} else {
		console.log(resp);
	}
});