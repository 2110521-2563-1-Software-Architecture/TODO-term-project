import { Module } from '@nestjs/common';
import { ConfigModule } from '../config/config.module';
import { PubSubService } from './pub-sub.service';

@Module({
    imports: [ConfigModule],
    providers: [PubSubService],
    exports: [PubSubService],
})
export class PubSubModule {}
