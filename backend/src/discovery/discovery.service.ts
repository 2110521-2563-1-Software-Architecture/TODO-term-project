import { Injectable } from '@nestjs/common';
import { Peer } from './peerInterface';
@Injectable()
export class DiscoveryService {
    fileLocationList: Array<String>;
    nodeId: String;
    fileId: String;
    peer: Peer;
    constructor(fileId: String,peer: Peer) {
        this.fileId = fileId
        this.peer = peer
    }
    setFileLocation(): void {
        ///connect to some api and get list of file location
        // this.fileLocationList = some list 
    }
    selectNode(): String {
        ///broadcast message to every node in fileLocationList and check RTT
        ///then, choose the best node (lowest RTT)
        ///can 
        // this.nodeId = best node
        return this.nodeId
    }
    // may be end at this line if we move file transfer fuction to oter service
    getFile(): Promise<any> {
        ///file transfer
        return this.peer.getFile(this.fileId)
    }
    checkSum(): Boolean {
        ///check file changing
        return false;
    }


}
