import axios, { AxiosError } from 'axios';
import { EmitterService } from '@/shared/lib';

export const TOKEN_NOT_FOUND_EVENT = 'token-not-found';

export const http = axios.create({
	baseURL: import.meta.env.VITE_BASE_URL,
	headers: {
		'Content-type': 'application/json'
	}
});

/**
 * trow REFRESH_TOKEN_EVENT_NAME
 */
export const useHttpBearerToken = () => {
	const localStorageToken = localStorage.getItem('user');
	let token: string | null = null;

	if (null !== localStorageToken) {
		token = JSON.parse(localStorageToken)?.user.token;
	}

	/**
	 * Для данного события создан хук useRefreshTokenListener()
	 * как обработчик данного события в клиентском коде
	 */
	if (null === token) {
		EmitterService.emitter.emit(TOKEN_NOT_FOUND_EVENT);
	}

	return axios.create({
		baseURL: import.meta.env.VITE_BASE_URL,
		headers: {
			'Content-type': 'application/json',
			'Authorization': `Bearer ${token}`
		}
	});
};

export const tryRefreshToken = (error: unknown): boolean => {
	if (! (error instanceof AxiosError)) {
		return false;
	}

	const responseData = error.response?.data;

	if (responseData.code === 401) {
		EmitterService.emitter.emit(TOKEN_NOT_FOUND_EVENT);
	}

	return true;
};

