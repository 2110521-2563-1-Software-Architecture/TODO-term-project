import { PubSub } from '@google-cloud/pubsub';
import { Injectable } from '@nestjs/common';
import { ConfigService } from '../config/config.service';

@Injectable()
export class PubSubService {
    private client: PubSub;
    constructor(config: ConfigService) {
        this.client = new PubSub({
            projectId: config.googleProjectId,
            credentials: config.gCloudServiceKey,
        });
    }
}
