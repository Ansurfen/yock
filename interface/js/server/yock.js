import { Server, ServerCredentials } from "@grpc/grpc-js"

/**
 * @typedef {object} CallRequest
 * @property {String} Fn function name
 * @property {String} Arg function argument
 */

/**
 * @typedef {object} CallResponse
 * @property {String} Buf
 */

class YockInterface {
    server
    dict

    constructor() {
        this.server = new Server()
        this.dict = new Map()
    }

    /**
     * 
     * @param {String} func 
     * @param {Function} huloCall 
     */
    Call(func, huloCall) {
        this.dict.set(func, huloCall)
    }

    run() {
        this.server.addService(hulo_proto.HuloInterface.service, { Call: this.callHandle.bind(this) })
        const args = process.argv.slice(2);
        let port = 0;
        for (let i = 0; i < args.length; i++) {
            if (args[i] === '-p') {
                port = parseInt(args[i + 1], 10);
                break;
            }
        }
        if (port === 0) {
            console.log("invalid port")
            process.exit(1);
        }
        this.server.bindAsync(`localhost:${port}`, ServerCredentials.createInsecure(), () => {
            this.server.start()
            console.log('grpc server started')
        })
    }
}

export const YockInterface = YockInterface