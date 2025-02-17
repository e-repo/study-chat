import { AuthLayout } from '@/shared/ui/layout';
import { useRoleUser, useRoleAdmin, useMiddlewarePipeline } from './middleware';
import { Middleware, MiddlewareContext, MiddlewarePayload } from './middleware/types';

import {
	createRouter,
	createWebHistory,
	RouteLocationNormalized,
	RouteLocationRaw,
	Router,
	RouteRecordRaw
} from 'vue-router';

const routes: RouteRecordRaw[] = [
	{
		path: '/',
		name: 'Home',
		component: () => import('@/pages/home'),
		meta: {
			middleware: [
				useRoleUser,
				useRoleAdmin,
			] as Middleware[]
		}
	},
	{
		path: '/auth/users',
		name: 'Users',
		component: () => import('@/pages/auth/users'),
		meta: {
			middleware: [
				useRoleUser,
				useRoleAdmin,
			] as Middleware[]
		}
	},
	{
		path: '/login',
		name: 'Login',
		component: () => import('@/pages/auth/login'),
		meta: {
			layout: AuthLayout,
		}
	},
	{
		path: '/confirm-reset-password',
		name: 'ConfirmPassword',
		component: () => import('@/pages/auth/confirm-password'),
		meta: {
			layout: AuthLayout,
		}
	},
	{
		path: '/confirm-email',
		name: 'ConfirmEmail',
		component: () => import('@/pages/auth/confirm-email'),
		meta: {
			layout: AuthLayout,
		}
	},
	{
		path: '/request-reset-password',
		name: 'RequestResetPassword',
		component: () => import('@/pages/auth/request-reset-password'),
		meta: {
			layout: AuthLayout,
		}
	},
	{
		path: '/forbidden',
		name: 'Forbidden',
		component: () => import('@/pages/auth/forbidden'),
		meta: {
			layout: AuthLayout,
		}
	},
	{
		path: '/:catchAll(.*)',
		name: 'NotFound',
		component: () => import('@/pages/auth/not-found'),
		meta: {
			layout: AuthLayout,
		}
	},
];

const router: Router = createRouter({
	routes,
	history: createWebHistory(import.meta.env.BASE_URL),
});

router.beforeEach((to: RouteLocationNormalized, from: RouteLocationNormalized): RouteLocationRaw | void => {
	if (! to.meta.middleware) {
		return;
	}

	const middlewares = to.meta.middleware as Middleware[];

	const context: MiddlewareContext = {
		to,
		from,
	};

	const payload: MiddlewarePayload = {
		...context,
		nextMiddleware: useMiddlewarePipeline(context, middlewares, 1)
	};

	return middlewares[0](payload);
});


export default router;

