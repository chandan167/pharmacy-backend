import express, { Application, NextFunction, Request, Response } from 'express';
import morgan from 'morgan';
import cors from 'cors';
import { NotFound } from 'http-errors';
import { globalErrorHandler } from './middleware/global-error-handler';
import path from 'path';
import i18n from 'i18n';

export const app: Application = express();

// i18n configuration
i18n.configure({
    locales: ['en', 'es'],
    directory: path.join(__dirname, '..', 'locales'),
    defaultLocale: 'en',
    queryParameter: 'lang',
    header: 'accept-language',
    cookie: 'locale',
    fallbacks: {
        'es-MX': 'es',
        'fr-CA': 'fr',
        '*': 'en' // Fallback everything else to English
    },
    autoReload: true,
    updateFiles: false,
    syncFiles: false
});

app.use(morgan('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cors());
app.use(i18n.init);
// Optional: Manual locale override via query parameter
app.use((req: Request, _: Response, next: NextFunction) => {
    if (req.query.lang) {
        req.setLocale(req.query.lang as string);
    }
    next();
});

app.use((req: Request, _: Response, next: NextFunction) => {
    next(new NotFound(req.__('route.not.found')));
});

app.use(globalErrorHandler);
