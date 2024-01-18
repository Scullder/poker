/**
 * Client should send data only about himself
 * Information about client updates only with websocket events
 */
export default class Client {
    constructor(manager) {
        this.manager = manager
        this.data = {}
    } 

    joinToGame = (data) => {
        this.manager.sendEvent('joinToGame', data)
    }

    // send client information
    updateRequest = (data) => {
        this.manager.sendEvent('playerUpdate', data)
    }

    updated = (data) => {
        
    }
}