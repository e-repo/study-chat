import { MiddlewarePayload } from './types';
import { RouteLocationRaw } from 'vue-router';
import { useUserModel } from '@/entities/user';

export default function useRoleUser(payload: MiddlewarePayload): RouteLocationRaw | void {
	const userModel = useUserModel();
	const nextMiddleware = payload.nextMiddleware;

	if (! userModel.isAuthenticated) {
		return {
			name: 'Login'
		};
	}

	if (nextMiddleware) {
		return nextMiddleware();
	}
}
