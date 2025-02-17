import { useUserModel } from '@/entities/user';
import { TOKEN_NOT_FOUND_EVENT } from '@/shared/api';
import { EmitterService }  from '@/shared/lib';

export const useRefreshTokenListener = () => {
	EmitterService.emitter.on(TOKEN_NOT_FOUND_EVENT, async () => {
		const userModel = useUserModel();

		await userModel.refreshToken();
		location.reload();
	});
};
