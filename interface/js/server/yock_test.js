import { YockInterface } from "./yock"

var server = new YockInterface()

server.Call("SayHello", function (request) {
    console.log(request)
    return { Buf: "I'm nodejs" }
})

server.run()