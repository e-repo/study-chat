import { useHttpBearerToken, tryRefreshToken } from '@/shared/api';
import { useRefreshTokenListener } from '@/entities/user';

const http = useHttpBearerToken();

useRefreshTokenListener();

export const fetchServiceMenuItems = async () => {
	try {
		return (
			await http.get('/service/v1/menu', {
				params: {
					name: 'service'
				}
			})
		).data;
	} catch (error: unknown) {
		tryRefreshToken(error);
	}
};

export const fetchBlogMenuItems = async () => {
	try {
		return (
			await http.get('/service/v1/menu', {
				params: {
					name: 'blog'
				}
			})
		).data;
	} catch (error: unknown) {
		tryRefreshToken(error);
	}
};
