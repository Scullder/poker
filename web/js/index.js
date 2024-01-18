import WebSocketManager from './websocket/manager.js'
import EventHandler from './websocket/eventHandler.js';
import Client from './client/client.js';

let client;
let handler = new EventHandler;

// will not be called before WebSocketManager is init
handler.on('joined', (payload) => {
    client.updated(payload)
})

$(document).on(('load'), () => {
    const manager = new WebSocketManager("ws://127.0.0.1:8080/websocket", handler);
    client = new Client(manager)
})

$('#join').on('click', () => {
    client.join({
        name: 'PlayerOne',
    })
})
