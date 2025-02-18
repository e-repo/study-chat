import { http, tryRefreshToken, useHttpBearerToken } from '@/shared/api';
import { CreateUser } from '@/entities/user';

const httpBearerToken = useHttpBearerToken();

export const fetchToken = async (email: string, password: string) => {
	return (
		await http.post('/auth', {
			email,
			password
		})
	).data;
};

export const refreshToken = async (refreshToken: string) => {
	return (
		await http.post('/auth/token-refresh', {
			refreshToken
		})
	).data;
};

export const requestCreateUser = async (user: CreateUser) => {
	try {
		return (
			await httpBearerToken.post('/users', user)
		).data;
	} catch (error: unknown) {
		tryRefreshToken(error);
	}
};
