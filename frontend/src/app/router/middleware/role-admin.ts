import { MiddlewarePayload } from './types';
import { RouteLocationRaw } from 'vue-router';
import { useUserModel } from '@/entities/user';
import { Role } from '@/shared/lib';

export default function useRoleAdmin(payload: MiddlewarePayload): RouteLocationRaw | void {
	const userModel = useUserModel();
	const nextMiddleware = payload.nextMiddleware;

	if (! Role.has(Role.UserRole.ROLE_ADMIN, userModel.userFromToken.roles)) {
		return {
			name: 'Forbidden'
		};
	}

	if (nextMiddleware) {
		return nextMiddleware();
	}
}
