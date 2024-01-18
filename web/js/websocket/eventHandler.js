export default class EventHandler {
    constructor() {
        this.handlers = new Map();
    }

    on = (eventName, callback) => {
        this.handlers.set(eventName, callback);
    }

    getHandler = (eventName) => {
        if (!this.handlers.has(eventName)) {
            console.log("unsupported event");
            return
        }

        return this.handlers.get(eventName)
    }

    call = (eventName) => {
        const callback = this.getHandler(eventName)
        callback();
    } 
}

