import { config } from 'dotenv-flow';
import { cleanEnv, num, str } from 'envalid';
config();

export const environment = cleanEnv(process.env, {
    NODE_ENV: str({ default: 'development', choices: ['development', 'production', 'stage', 'test'] }),
    PORT: num(),
    SERVER_URL: str(),
    LOG_DIR: str({ default: 'logs' })
});