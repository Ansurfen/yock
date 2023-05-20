import { YockInterface as _YockInterface } from './proto'
import { credentials } from '@grpc/grpc-js'

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
    client

    /**
     * 
     * @param {number} port 
     */
    constructor(port) {
        this.client = new _YockInterface(`localhost:${port}`, credentials.createInsecure())
    }

    /**
     * 
     * @param {CallRequest} request 
     * @param {Function} handle
     */
    Call(request, handle) {
        if (this.client == null) {
            handle({ Ok: false, Buf: "sys err" })
        }
        this.client.Call(request, function (err, response) {
            if (err) {
                handle({ Ok: false, Buf: "sys err" })
            }
            handle(response)
        })
    }
}

export const YockInterface = YockInterface