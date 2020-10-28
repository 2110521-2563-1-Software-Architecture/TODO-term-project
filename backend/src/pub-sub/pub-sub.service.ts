import { Message, PubSub, Subscription } from '@google-cloud/pubsub';
import { Injectable } from '@nestjs/common';
import { ConfigService } from '../config/config.service';

@Injectable()
export class PubSubService {
    private client: PubSub;
    private subscription: Subscription;

    constructor(config: ConfigService) {
        this.client = new PubSub({
            projectId: config.googleProjectId,
            credentials: config.gCloudServiceKey,
        });
    }

    async publish(topicName: string, attributes: Message['attributes']) {
        try {
            const messageId = await this.client
                .topic(topicName)
                .publish(null, attributes);
            console.log(`Message ${messageId} published.`);
        } catch (err) {
            console.error(err.stack);
        }
    }

    async subscribe(subscriptionName: string) {
        this.subscription = this.client.subscription(subscriptionName);
        this.subscription.on('message', this.messageHandler);
    }

    unsubscribe() {
        this.subscription.removeListener('message', this.messageHandler);
    }

    private messageHandler(message: Message) {
        console.log(message.id);
        console.log(message.attributes);

        message.ack();
    }
}
