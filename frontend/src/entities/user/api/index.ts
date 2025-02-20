import { http, tryRefreshToken, useHttpBearerToken } from '@/shared/api';
import { CreateUser } from '@/entities/user';

const httpBearerToken = useHttpBearerToken();

export const fetchToken = async (email: string, password: string) => {
	return (
		await http.post('/sign-in', {
			email,
			password
		})
	).data;
};

export const refreshToken = async (refreshToken: string) => {
	return (
		await http.post('/token-refresh', {
			refreshToken
		})
	).data;
};

export const requestCreateUser = async (user: CreateUser) => {
	try {
		return (
			await httpBearerToken.post('/sign-up', user)
		).data;
	} catch (error: unknown) {
		tryRefreshToken(error);
	}
};
