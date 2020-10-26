import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { DiscoveryService } from './discovery/discovery.service';
import { DiscoveryModule } from './discovery/discovery.module';

@Module({
    imports: [DiscoveryModule],
    controllers: [AppController],
    providers: [AppService, DiscoveryService],
})
export class AppModule {}
