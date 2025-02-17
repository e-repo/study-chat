import { AxiosError } from 'axios';
import { Violation } from '@/shared/lib';

type AlertType = 'error' | 'success' | 'warning' | 'info' | undefined;

export interface Alert {
	isShow: boolean;
	type: AlertType;
	message: string;
}

export const useProcessErrorResponse = (error: AxiosError, alert: Alert) => {
	if (error.response?.status !== 422) {
		alert.isShow = true;
		alert.type = 'error';

		return;
	}

	const violation = error.response?.data as Violation;

	alert.isShow = true;
	alert.message = violation.errors[0].detail || alert.message;
	alert.type = 'error';
};
