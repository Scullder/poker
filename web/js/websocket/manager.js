class Event {
    constructor(type, payload) {
        this.type = type
        this.payload = payload
    }
}

export default class WebSocketManager {
    constructor(link, handler) {
        this.socket = this.initWebSocket(link)
        this.handler = handler
    }

    routeEvent = (event) => {
        if (event.type === undefined) {
            console.log("no 'type' field in event")
            return;
        }
    
        this.handler.call(event.type)
    }
    
    sendEvent = (eventName, payload) => {
        const event = new Event(eventName, payload)
        this.socket.send(JSON.stringify(event))
    }
    
    initWebSocket = (link) => {
        const socket = new WebSocket(link)
    
        socket.onopen = () => {
            console.log("Successfully Connected")
            socket.send("Hi From the Client!")
        }
    
        socket.onmessage = (msg) => {
            const data = JSON.parse(msg.data);
            const event = new Event(data.type, data.payload)
            routeEvent(event);
        }
    
        socket.onclose = event => {
            console.log("Socket Closed Connection: ", event)
            socket.send("Client Closed!")
        }
    
        socket.onerror = error => {
            console.log("Socket Error: ", error)
        }
    
        return socket
    }

}


