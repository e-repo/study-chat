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
							v-model="restoreForm.isValid"
							@submit.prevent="onSubmit"
						>

							<v-text-field
								v-model="restoreForm.password"
								:rules="[passRules.required, passRules.counter]"
								:type="isPassShow ? 'password' : 'text'"
								label="Пароль"
								variant="underlined"
								:append-inner-icon="isPassShow ? 'mdi-eye-off' : 'mdi-eye'"
								counter
								@click:append-inner="isPassShow = !isPassShow"
							></v-text-field>

							<v-text-field
								v-model="restoreForm.repeatPassword"
								:rules="[newPassRules.required, newPassRules.counter, newPassRules.equal]"
								:type="isNewPassShow ? 'password' : 'text'"
								label="Новый пароль"
								variant="underlined"
								:append-inner-icon="isNewPassShow ? 'mdi-eye-off' : 'mdi-eye'"
								counter
								@click:append-inner="isNewPassShow = !isNewPassShow"
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
									:disabled="!restoreForm.isValid"
									:loading="restoreForm.loading"
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
import { reactive, ref } from 'vue';
import { useUserModel } from '@/entities/user';
import { Alert, useProcessErrorResponse } from '@/shared/ui/alert';
import { useRoute } from 'vue-router';
import { AxiosError } from 'axios';

const userModel = useUserModel();
const route = useRoute();

interface RestoreForm {
	isValid: boolean;
	password: string | null;
	repeatPassword: string | null;
	loading: boolean;
}

const restoreForm = reactive<RestoreForm>({
	isValid: false,
	password: null,
	repeatPassword: null,
	loading: false,
})

const alert = reactive<Alert>({
	isDefaultSlotShow: false,
	type: undefined,
	message: '',
});

const onSubmit = async (): Promise<void> => {
	if (! restoreForm.isValid) {
		return;
	}

	if (! route.query.token) {
		return;
	}

	restoreForm.loading = true;

	try {
		await userModel.confirmResetPassword(
			route.query.token as string,
			restoreForm.password as string,
		);

		alert.isShow = true;
		alert.type = 'success';
		alert.message = 'Пароль успешно изменен.'

		restoreForm.loading = false;
	} catch (error: any) {
		restoreForm.loading = false;
		restoreForm.isValid = false;

		if (error instanceof AxiosError) {
			alert.message = 'Ошибка сброса пароля, проверьте правильность ввода email или обратитесь к администратору';

			useProcessErrorResponse(error, alert);
		}
	}
};

const isPassShow = ref<boolean>(true);
const isNewPassShow = ref<boolean>(true);

const passRules = {
	required: FormHelper.requiredRule,
	counter: (value: string): FormHelper.RuleType => {
		if (value.length <= 6) {
			return 'Длинна пароля не менее 6-ми символов';
		}

		return value.length <= 20 || 'Максимальное число символов 20';
	}
};

const newPassRules = {
	required: FormHelper.requiredRule,
	counter: (value: string): FormHelper.RuleType => {
		if (value.length <= 6) {
			return 'Длинна пароля не менее 6-ми символов';
		}

		return value.length <= 20 || 'Максимальное число символов 20';
	},
	equal: (): FormHelper.RuleType =>
		restoreForm.password === restoreForm.repeatPassword || 'Пароли не совпадают.',
};
</script>

<style scoped></style>
