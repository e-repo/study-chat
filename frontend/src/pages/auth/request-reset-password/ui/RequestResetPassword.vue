<template>

	<v-container class="h-100">
		<v-row class="h-100 align-center">
			<v-col class="d-flex justify-center">

				<v-card
					min-width="340"
					width="400"
					elevation="4"
				>
					<v-card-item>
						<v-card-title>
							<h4 class="text-center">Изменение пароля</h4>
						</v-card-title>
						<v-alert
							v-if="alert.isShow"
							:text="alert.message"
							:type="alert.type"
							variant="outlined"
						></v-alert>
					</v-card-item>

					<v-divider class="mx-4 mb-2"></v-divider>

					<v-card-text>

						<v-form
							v-model="resetPasswordForm.isValid"
							@submit.prevent="onSubmit"
						>

							<v-text-field
								v-model="resetPasswordForm.email"
								:rules="[emailRules.required, emailRules.email]"
								type="email"
								label="Email"
								variant="underlined"
							></v-text-field>

							<div
								class="d-flex flex-wrap justify-end mt-4"
							>
								<v-btn
									to="/login"
									color="error"
									variant="outlined"
									type="button"
								>
									Вход
								</v-btn>
								<v-btn
									class="ml-2"
									:disabled="!resetPasswordForm.isValid"
									:loading="resetPasswordForm.loading"
									type="submit"
									color="success"
								>
									Изменить пароль
								</v-btn>
							</div>

						</v-form>

					</v-card-text>
				</v-card>

			</v-col>
		</v-row>
	</v-container>

</template>
<script setup lang="ts">
import { FormHelper } from '@/shared/lib';
import { reactive } from 'vue';
import { useUserModel } from '@/entities/user';
import { AxiosError } from 'axios';
import { Alert, useProcessErrorResponse } from '@/shared/ui/alert';

const REGISTRATION_SOURCE = 'admin_panel';

const userModel = useUserModel();

interface RequestResetPasswordForm {
	isValid: boolean;
	email: string | null;
	registrationSource: string;
	loading: boolean;
}

const resetPasswordForm = reactive<RequestResetPasswordForm>({
	isValid: false,
	email: null,
	registrationSource: REGISTRATION_SOURCE,
	loading: false
})

const alert = reactive<Alert>({
	isDefaultSlotShow: false,
	type: undefined,
	message: '',
});

const onSubmit = async (): Promise<void> => {
	if (! resetPasswordForm.isValid) {
		return;
	}

	resetPasswordForm.loading = true;

	try {
		await userModel.requestResetPassword(
			resetPasswordForm.email as string,
			resetPasswordForm.registrationSource
		);

		alert.isShow = true;
		alert.type = 'success';
		alert.message = 'На указанную почту было отправлено письмо, следуйте указаниям в письме.'

		resetPasswordForm.loading = false;
	} catch (error: any) {
		resetPasswordForm.loading = false;
		resetPasswordForm.isValid = false;

		if (error instanceof AxiosError) {
			alert.message = 'Ошибка сброса пароля, проверьте правильность ввода email или обратитесь к администратору';

			useProcessErrorResponse(error, alert);
		}
	}
};

const emailRules = {
	required: FormHelper.requiredRule,
	email: (value: string): FormHelper.RuleType => FormHelper.emailPattern.test(value) || 'Некорректный \'Email\'',
};
</script>
<style scoped></style>
