export interface UserListFilter {
	offset: number;
	limit: number;
	firstName?: string;
	lastName?: string;
	email?: string;
	role?: string;
	status?: string;
}

export type UserStatus = 'wait' | 'active' | 'blocked';

export interface UserProfile {
	id: string;
	email: string;
	firstName: string;
	role: string;
	status: UserStatus;
	createdAt: string;
	lastName?: string;
	registrationSource?: string|null;
}
