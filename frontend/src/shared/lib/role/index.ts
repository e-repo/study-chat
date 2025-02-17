export enum UserRole {
	ROLE_ADMIN = 'ROLE_ADMIN',
	ROLE_USER = 'ROLE_USER',
}

export const has = (role: UserRole, roles: string[]): boolean => {
	return roles.some(item => item === role);
};
