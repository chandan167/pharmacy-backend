import { app } from './application';
import { environment } from './config/environment';
import { logger } from './util/logger';


const server = app.listen(environment.PORT);

(() => {
    try {
        // TODO: Database connection 

        logger.info(`APPLICATION_STARTED`, {
            meta: {
                PORT: environment.PORT,
                SERVER_URL: environment.SERVER_URL
            }
        });
    } catch (error) {
        if (error) {
            logger.error(`APPLICATION_STARTED`, {
                meta: error
            });
        }

        server.close((error) => {
            if (error) {
                logger.error(`APPLICATION_ERROR`, { meta: error });
            }
        });
        process.exit(1);
    }
})();