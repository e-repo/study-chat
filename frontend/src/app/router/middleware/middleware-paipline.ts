import { Middleware, MiddlewareContext, NextMiddlewareCallable } from './types';

export default function useMiddlewarePipeline(
	context: MiddlewareContext,
	middlewares: Middleware[],
	index: number
): NextMiddlewareCallable | null {
	const nextMiddleware = middlewares[index];

	if (! nextMiddleware) {
		return null;
	}

	return () => {
		const nextPipeline = useMiddlewarePipeline(context, middlewares, index + 1);

		return nextMiddleware({ ...context, nextMiddleware: nextPipeline });
	};
}
