export const parsePayload = (token: string) => {
	try {
		return JSON.parse(
			decodeURIComponent(escape(atob(
				token.split('.')[1]
			)))
		);
	} catch (err) {
		return null;
	}
};
