import { http, tryRefreshToken, useHttpBearerToken } from '@/shared/api';
import { CreateUser } from '@/entities/user';

const httpBearerToken = useHttpBearerToken();

export const fetchToken = async (username: string, password: string) => {
	return (
		await http.post('/auth/login-check', {
			username,
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

export const requestResetPassword = async (email: string, registrationSource: string) => {
	return (
		await http.post('/auth/v1/user/request-reset-password', {
			email,
			registrationSource
		})
	).data;
};

export const requestCreateUser = async (user: CreateUser) => {
	try {
		return (
			await httpBearerToken.post('/auth/v1/user/sign-up', user)
		).data;
	} catch (error: unknown) {
		tryRefreshToken(error);
	}
};

export const activate = async (userId: string) => {
	try {
		return (
			await httpBearerToken.patch(`/auth/v1/user/${userId}/status`, {
				status: 'active'
			})
		).data;
	} catch (error: unknown) {
		tryRefreshToken(error);
	}
};

export const block = async (userId: string) => {
	try {
		return (
			await httpBearerToken.patch(`/auth/v1/user/${userId}/status`, {
				status: 'blocked'
			})
		).data;
	} catch (error: unknown) {
		tryRefreshToken(error);
	}
};

export const confirmResetPassword = async (token: string, password: string)=> {
	return (
		await http.post('/auth/v1/user/confirm-reset-password', {
			token,
			password,
		})
	).data;
};

export const confirmEmail = async (userId: string, token: string)=> {
	return (
		await http.post('/auth/v1/user/confirm-email', {
			userId,
			token,
		})
	).data;
};
