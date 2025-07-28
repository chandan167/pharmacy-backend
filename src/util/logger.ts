import winston from 'winston';
// const path = require('path');
import { existsSync, mkdirSync } from 'fs';
import winstonDaily from 'winston-daily-rotate-file';
// import LokiTransport from 'winston-loki';
import util from 'util';
import * as sourceMapSupport from 'source-map-support';
import { environment } from '../config/environment';
import { blue, red, yellow, cyanBright, magenta } from 'colorette';

// Linking trace support
sourceMapSupport.install();
// logs dir
const logDir = environment.LOG_DIR;

if (!existsSync(logDir)) {
    mkdirSync(logDir);
}



// Define log format
const logFormat = winston.format.printf(({ timestamp, level, message, meta = {} }) => {
    level = level.toUpperCase();
    const customMeta = util.inspect(meta, {
        showHidden: false,
        depth: Infinity
    });
    return `${level} [${timestamp}] ${message}\n${'META'} ${customMeta}\n\n`;
});

/*
 * Log Level
 * error: 0, warn: 1, info: 2, http: 3, verbose: 4, debug: 5, silly: 6
 */
export const logger = winston.createLogger({
    defaultMeta: {
        meta: {}
    },
    format: winston.format.combine(
        winston.format.timestamp({
            format: 'YYYY-MM-DD HH:mm:ss'
        }),
        logFormat
    ),
    transports: [
        // new LokiTransport({
        //     host: 'http://localhost:3001', // Loki server URL
        //     labels: { job: 'pharmacy-app' },
        //     json: true,
        //     format: winston.format.json()
        // }),
        // debug log setting
        new winstonDaily({
            level: 'debug',
            datePattern: 'YYYY-MM-DD',
            dirname: `${logDir}/debug`, // log file /logs/debug/*.log in save
            filename: '%DATE%.log',
            maxFiles: 365, // 30 Days saved
            json: true,
            zippedArchive: false,
            handleExceptions: true,
        }),
        // error log setting
        new winstonDaily({
            level: 'error',
            datePattern: 'YYYY-MM-DD',
            dirname: `${logDir}/error`, // log file /logs/error/*.log in save
            filename: '%DATE%.log',
            maxFiles: 365, // 30 Days saved
            handleExceptions: true,
            json: true,
            zippedArchive: false
        }),
        new winstonDaily({
            level: 'info',
            datePattern: 'YYYY-MM-DD',
            dirname: `${logDir}/info`, // log file /logs/error/*.log in save
            filename: '%DATE%.log',
            maxFiles: 365, // 30 Days saved
            handleExceptions: true,
            json: true,
            zippedArchive: false
        })
    ]
});

if (environment.isDev) {
    const colorizeLevel = (level: string) => {
        switch (level) {
            case 'ERROR':
                return red(level);
            case 'INFO':
                return blue(level);
            case 'WARN':
                return yellow(level);
            default:
                return level;
        }
    };
    const logFormat = winston.format.printf(({ timestamp, level, message, meta = {} }) => {
        level = colorizeLevel(level.toUpperCase());
        const customMeta = util.inspect(meta, {
            showHidden: false,
            depth: Infinity,
            colors: true
        });
        timestamp = cyanBright(timestamp as string);
        return `${level} [${timestamp}] ${message}\n${magenta('META')} ${customMeta}\n\n`;
    });
    logger.add(
        new winston.transports.Console({
            format: winston.format.combine(
                winston.format.timestamp({
                    format: 'YYYY-MM-DD HH:mm:ss'
                }),
                logFormat
            )
        })
    );
}



export const stream = {
    write: (message: string) => {
        logger.info(message.substring(0, message.lastIndexOf('\n')));
    }
};
