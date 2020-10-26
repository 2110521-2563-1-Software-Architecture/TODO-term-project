export  interface  Peer{
    connect(nodeId : number): Promise<any> /// initiate connection
    broadcast(message : String): Promise<any> /// find RTT
    getFile(fileId : number): Promise<any> /// get file from another node
    disconnect() : Promise<any> 
}

