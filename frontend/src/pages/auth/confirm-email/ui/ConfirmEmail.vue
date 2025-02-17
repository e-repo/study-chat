<template>
	<v-progress-linear
		v-if="isShowConfirmProgress"
		color="success"
		indeterminate
	></v-progress-linear>
	<v-container class="h-100">
		<v-row class="h-100 align-center">
			<v-col class="d-flex justify-center">

				<h2 class="text-center">{{ message }}</h2>

			</v-col>
		</v-row>
	</v-container>

</template>

<script setup lang="ts">
import { useUserModel } from '@/entities/user';
import { useRoute } from 'vue-router';
import { onMounted, ref } from 'vue';

const isShowConfirmProgress = ref<boolean>(true);
const message = ref<string>('')

const userModel = useUserModel();
const route = useRoute();

onMounted(async () => {
	if (! route.query.userId  || route.query.token) {
		message.value = showAttributesError();
	}

	try {
		await userModel.confirmEmail(
			route.query.userId as string,
			route.query.token as string,
		);

		message.value = 'Email адрес подтвержден успешно, можете войти в систему.';
	} catch (error: any) {
		message.value = 'Возможно вы уже подтвердили email, попробуйте войти в систему или обратитесь к администратору';
	}

	isShowConfirmProgress.value = false;
});

const showAttributesError = (): string => 'Отсутствуют атрибуты подтверждения пароля.'

</script>

<style scoped></style>
