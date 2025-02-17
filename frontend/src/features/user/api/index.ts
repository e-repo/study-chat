import { useHttpBearerToken, tryRefreshToken } from '@/shared/api';
import { useRefreshTokenListener } from '@/entities/user';
import { UserListFilter } from '@/features/user';

const http = useHttpBearerToken();

useRefreshTokenListener();

export const fetchUserList = async (filter: UserListFilter)=> {
	try {
		return (
			await http.get('/auth/v1/users', {
				params: filter
			})
		).data;
	} catch (error: unknown) {
		tryRefreshToken(error);
	}
};

export const fetchUserById = async (userId: string)=> {
	try {
		return (
			await http.get(`/auth/v1/user/${userId}`)
		).data;
	} catch (error: unknown) {
		tryRefreshToken(error);
	}
};
