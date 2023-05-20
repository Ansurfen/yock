import { join } from 'path'
import { loadPackageDefinition } from '@grpc/grpc-js'
import { loadSync } from '@grpc/proto-loader'

const PROTO_PATH = join(__dirname, '../yock.proto')
const packageDefinition = loadSync(PROTO_PATH, { keepCase: true, longs: String, enums: String, defaults: true, oneofs: true })
const protoDescriptor = loadPackageDefinition(packageDefinition)

const yock_proto = protoDescriptor.Yock

export default yock_proto