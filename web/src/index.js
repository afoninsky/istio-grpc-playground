import { grpc } from 'grpc-web-client'
import { Empty } from '../proto/streamer_pb'
import { Streamer } from '../proto/streamer_pb_service'

const messages = document.querySelector('#messages')

function write(text) {
  messages.appendChild(document.createTextNode(text + "\n"))
}

// https://github.com/improbable-eng/grpc-web/blob/master/client/grpc-web-client/docs/invoke.md
function connect(host) {
  write(`Connection to ${host}`)
  grpc.invoke(Streamer.Subscribe, {
    debug: true,
    request: new Empty(),
    host,
    onMessage: message => {
        write(`message from the server: ${JSON.stringify(message.toObject())}`)
    },
    onEnd: (code, msg) => {
      write(`Request ended with code ${code} and message "${msg}"`)
    }
  })
}


connect('http://localhost:8080')
