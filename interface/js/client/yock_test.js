import { YockInterface } from "./yock"

var cli = new YockInterface(9090)

cli.Call({ Func: "SayHello" }, (request) => {
    console.log(request)
})