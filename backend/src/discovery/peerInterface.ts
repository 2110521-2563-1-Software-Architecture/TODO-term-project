export interface Peer {
    peerId: String
    connect(nodeId: String): Promise<any> /// initiate connection
    broadcast(message: String): Promise<any> /// find RTT
    getFile(fileId: String): Promise<any> /// get file from another node
    disconnect(): Promise<any>
}

