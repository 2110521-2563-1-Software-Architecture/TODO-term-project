import { Module } from '@nestjs/common';
import { DiscoveryService } from '@nestjs/core';

@Module({
    providers: [DiscoveryService],
    exports: [DiscoveryService],
})
export class DiscoveryModule { }
