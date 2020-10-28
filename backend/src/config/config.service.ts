import { Injectable } from '@nestjs/common';

@Injectable()
export class ConfigService {
    private getEnv(name: string): string {
        const env = process.env[name];
        if (name) {
            return env;
        }
        throw new Error(`${name} is undefined`);
    }

    get googleProjectId(): string {
        return this.getEnv('GOOGLE_PROJECT_ID');
    }

    get gCloudServiceKey() {
        const key = this.getEnv('GCLOUD_SERVICE_KEY');
        return JSON.parse(key);
    }
}
