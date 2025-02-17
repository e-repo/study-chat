import { defineStore } from 'pinia';
import { CreateUser, UserApi } from '../index';
import { JwtParser } from '@/shared/lib';

interface AuthStoreUser {
	isAuthenticated: boolean,
	token: string|null,
	refreshToken: string|null,
}

interface UserToken {
	token: string,
	refreshToken: string,
}

export interface AuthStore {
	user: AuthStoreUser,
}

export interface UserFromToken
{
	id: string;
	firstName: string;
	email: string;
	roles: string[];
}

interface TokenPayload {
	user: UserFromToken;
	username: string;
}

export const useUserModel = defineStore({
	id: 'user',

	state: () => ({
		user: {
			isAuthenticated: false,
			token: null,
			refreshToken: null
		}
	} as AuthStore),

	actions: {
		async singIn(username: string, password: string) {
			const result = await UserApi.fetchToken(username, password);

			this.setToken({
				token: result.token,
				refreshToken: result.refreshToken
			});
		},
		async signUp(user: CreateUser) {
			await UserApi.requestCreateUser(user);
		},
		async block(userId: string) {
			await UserApi.block(userId);
		},
		async activate(userId: string) {
			await UserApi.activate(userId);
		},
		async requestResetPassword(email: string, registrationSource: string) {
			await UserApi.requestResetPassword(email, registrationSource);
		},
		async confirmResetPassword(token: string, newPassword: string) {
			await UserApi.confirmResetPassword(token, newPassword);
		},
		async confirmEmail(userId: string, token: string) {
			await UserApi.confirmEmail(userId, token);
		},
		logout() {
			this.$reset();
		},
		async refreshToken() {
			if (null === this.user.refreshToken) {
				throw new TypeError('Не найден токен для обновления.');
			}

			const result = await UserApi.refreshToken(this.user.refreshToken);

			this.setToken({
				token: result.token,
				refreshToken: result.refreshToken
			});
		},
		setToken(userToken: UserToken) {
			this.user.isAuthenticated = true;
			this.user.token = userToken.token;
			this.user.refreshToken = userToken.refreshToken;
		}
	},

	getters: {
		token(): string|null {
			return this.user.token;
		},
		isAuthenticated(): boolean {
			return this.user.isAuthenticated;
		},
		tokenPayload(): TokenPayload | null {
			if (null === this.user.token) {
				return null;
			}

			return JwtParser.parsePayload(this.user.token);
		},
		userFromToken(): UserFromToken {
			const tokenPayload = this.tokenPayload;

			if (null === tokenPayload) {
				throw Error('Токен не найден.');
			}

			tokenPayload.user.firstName = decodeURI(tokenPayload.user.firstName);
			return tokenPayload.user;
		}
	},

	persist: true
});


