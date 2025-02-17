export const parsePayload = (token: string) => {
	try {
		return JSON.parse(
			atob(
				token.split('.')[1]
			)
		);
	} catch (err) {
		return null;
	}
};
