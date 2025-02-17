import { RouteLocationNormalized, RouteLocationRaw } from 'vue-router';

export interface MiddlewarePayload {
	to?: RouteLocationNormalized;
	from?: RouteLocationNormalized;
	nextMiddleware?: NextMiddlewareCallable | null;
}

export type NextMiddlewareCallable = () => RouteLocationRaw | void;

export interface MiddlewareContext {
	to?: RouteLocationNormalized;
	from?: RouteLocationNormalized;
}

export type Middleware = (payload: MiddlewarePayload) => RouteLocationRaw | void;