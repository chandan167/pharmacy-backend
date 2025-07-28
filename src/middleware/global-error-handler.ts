import { ErrorRequestHandler, NextFunction, Request, Response } from 'express';
import { HttpError, isHttpError } from 'http-errors';
import { getReasonPhrase, StatusCodes } from 'http-status-codes';
import { environment } from '../config/environment';

export const globalErrorHandler: ErrorRequestHandler = (error: Record<string, unknown>, _: Request, res: Response, __: NextFunction) => {
    const status = (error as HttpError).status || StatusCodes.INTERNAL_SERVER_ERROR;
    const message = error.message || getReasonPhrase(status);
    if (isHttpError(error)) {
        return res.status(error.status).json({ message: error.message });
    }

    if (!environment.isProd) {
        return res.status(status).json({ message });
    }

    // Always return a response in production if not handled above
    return res.status(StatusCodes.INTERNAL_SERVER_ERROR).json({ message: getReasonPhrase(StatusCodes.INTERNAL_SERVER_ERROR) });
};