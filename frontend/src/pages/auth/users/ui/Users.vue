<template>

	<v-container fluid>
		<v-row>
			<v-col>
				<v-breadcrumbs :items="breadcrumbs" class="pl-0"></v-breadcrumbs>
			</v-col>
		</v-row>

		<v-row>
			<v-col>
				<h1>Список пользователей</h1>
			</v-col>
		</v-row>

		<v-row>
			<v-col>
				<v-data-table-server
					:items-per-page="tableOptions.itemsPerPage"
					:items-per-page-options="[tableOptions.itemsPerPage]"
					:headers="tableOptions.headers"
					:items="tableOptions.serverItems"
					:items-length="tableOptions.totalItems"
					:loading="tableOptions.loading"
					item-value="name"
					@update:options="loadItems"
				>
					<template v-slot:top>

						<div class="d-flex">

							<v-text-field
								v-model="tableOptions.email"
								density="compact"
								class="sort-input mr-2"
								placeholder="Email..."
								type="email"
								hide-details
							></v-text-field>
							<v-text-field
								v-model="tableOptions.firstName"
								density="compact"
								class="sort-input"
								placeholder="Имя..."
								hide-details
							></v-text-field>

							<v-spacer></v-spacer>

							<v-dialog
								v-model="createUserForm.isShow"
								max-width="500"
							>
								<template v-slot:activator="{ props: activatorProps }">
									<v-btn
										v-bind="activatorProps"
										color="text-grey-darken-1"
										text="Добавить"
										variant="outlined"
									></v-btn>
								</template>

								<template v-slot:default="{ isActive }">
									<v-card
										min-width="340"
										elevation="4"
									>
										<v-card-item>
											<v-card-title>
												<h4 class="text-center">Создание пользователя</h4>
											</v-card-title>
											<v-alert
												v-if="createUserForm.serverError"
												text="При создании пользователя возникла ошибка на стороне сервера, обратитесь к администратору."
												type="error"
												variant="outlined"
											></v-alert>
										</v-card-item>

										<v-divider class="mx-4 mb-2"></v-divider>

										<v-card-text>

											<v-form
												v-model="createUserForm.isValid"
												@submit.prevent="submitCreateUser"
											>

												<v-text-field
													v-model="createUserForm.firstName"
													:rules="[firstNameRules.required, firstNameRules.counter]"
													label="Имя пользователя"
													variant="underlined"
													counter
													@click:append-inner="isPassShow = !isPassShow"
												></v-text-field>

												<v-text-field
													v-model="createUserForm.email"
													:rules="[emailRules.required, emailRules.email]"
													type="email"
													label="Email"
													variant="underlined"
												></v-text-field>

												<v-text-field
													v-model="createUserForm.password"
													:rules="[passRules.required, passRules.counter]"
													:type="isPassShow ? 'password' : 'text'"
													label="Пароль"
													variant="underlined"
													:append-inner-icon="isPassShow ? 'mdi-eye-off' : 'mdi-eye'"
													counter
													@click:append-inner="isPassShow = !isPassShow"
												></v-text-field>

												<v-select
													v-model="createUserForm.registrationSource"
													label="Источник регистрации"
													variant="underlined"
													disabled
													:items="[createUserForm.registrationSource]"
												></v-select>

												<div
													class="d-flex flex-wrap justify-end mt-4"
												>
													<v-btn
														text="Закрыть"
														color="error"
														variant="outlined"
														@click="isActive.value = false"
													></v-btn>

													<v-spacer></v-spacer>

													<v-btn
														:disabled="!createUserForm.isValid"
														:loading="createUserForm.loading"
														type="submit"
														color="success"
														variant="outlined"
													>
														Создать
													</v-btn>
												</div>

											</v-form>

										</v-card-text>

									</v-card>
								</template>
							</v-dialog>

						</div>
						<v-divider class="mt-4 mb-4"></v-divider>

					</template>
					<template v-slot:item.actions="{ item }">
						<v-icon
							:disabled="item.status === 'wait'"
							class="me-2"
							@click="changeStatus(item)"
						>
							{{ item.status === 'blocked' ? 'mdi-account-reactivate-outline' : 'mdi-account-cancel-outline' }}
						</v-icon>
						<v-icon
							@click="userInfoDialogOpen(item)"
						>
							mdi-information-outline
						</v-icon>
					</template>
				</v-data-table-server>
			</v-col>
		</v-row>

	</v-container>

	<v-dialog
		transition="dialog-top-transition"
		v-model="userInfoDialog.isShow"
		width="600px"
	>
		<v-card
			prepend-icon="mdi-information-outline"
			title="Профиль пользователя"
		>
			<template v-slot:text>
				<div
					class="text-center"
					v-if="userInfoDialog.user === null"
				>
					<v-progress-circular indeterminate></v-progress-circular>
				</div>
				<div v-else>
					<v-table density="compact">
						<tbody>
							<tr
								v-for="(item, key) in userInfoDialog.user"
							>
								<th>{{ key + 1 }}.</th>
								<td>{{ item.attribute }}:</td>
								<td>{{ item.value ? `${item.value}` : '-' }}</td>
							</tr>
						</tbody>
					</v-table>
				</div>
			</template>
			<template v-slot:actions>
				<v-btn
					class="ms-auto"
					variant="outlined"
					text="Закрыть"
					@click="userInfoDialogClose()"
				></v-btn>
			</template>
		</v-card>
	</v-dialog>

</template>

<script setup lang="ts">
import { onMounted, reactive, ref, watch } from 'vue';
import { EmitterService, FormHelper, List } from '@/shared/lib';
import { CreateUser, useRefreshTokenListener, useUserModel } from '@/entities/user';
import { UserFetcher, UserProfile } from '@/features/user';
import { RegistrationSource } from '@/shared/lib/form/sign-up/types';
import { AxiosError } from 'axios';

useRefreshTokenListener();

const userModel = useUserModel();

type SortItem = { key: string, order?: boolean | 'asc' | 'desc' }

interface UserDialogItems
{
	attribute: string;
	value: string|null|undefined;
}

interface UserInfoDialog
{
	isShow: boolean;
	user: UserDialogItems[]|null;
}

interface LoadParam
{
	page: number;
	itemsPerPage: number;
	sortBy: SortItem[];
}

let userInfoDialog = reactive<UserInfoDialog>({
	isShow: false,
	user: null
});

let userList = ref<UserProfile[]>([])
let currentLoadUsersOptions: LoadParam | null = null;

const isPassShow = ref<boolean>(true);

interface CreateUserForm {
	serverError: boolean;
	isValid: boolean;
	firstName: string | null
	email: string | null;
	password: string | null;
	registrationSource: RegistrationSource;
	loading: boolean;
	isShow: boolean;
}

const createUserForm = reactive<CreateUserForm>({
	serverError: false,
	isValid: false,
	firstName: null,
	email: null,
	password: null,
	registrationSource: 'admin_panel',
	loading: false,
	isShow: false
})

const emailRules = {
	required: FormHelper.requiredRule,
	email: (value: string): FormHelper.RuleType => FormHelper.emailPattern.test(value) || 'Некорректный \'Email\'',
};

const passRules = {
	required: FormHelper.requiredRule,
	counter: (value: string): FormHelper.RuleType => {
		if (value.length < 6) {
			return 'Длинна пароля не менее 6-ти символов';
		}

		return value.length <= 20 || 'Максимальное число символов 20';
	}
};

const firstNameRules = {
	required: FormHelper.requiredRule,
	counter: (value: string): FormHelper.RuleType => {
		if (value.length < 3) {
			return 'Длинна имени не менее 3-eх символов';
		}

		return value.length <= 50 || 'Максимальное число символов 50';
	}
};

const submitCreateUser = async (): Promise<void> => {
	if (! createUserForm.isValid) {
		return;
	}

	createUserForm.loading = true;

	try {
		await userModel.signUp(<CreateUser>{
			firstName: createUserForm.firstName,
			email: createUserForm.email,
			password: createUserForm.password,
			registrationSource: createUserForm.registrationSource,
		});

		createUserForm.loading = false;
		createUserForm.isShow = false;

		await loadItems(currentLoadUsersOptions);
	} catch (error: any) {
		createUserForm.loading = false;
		createUserForm.isValid = false;

		if (error instanceof AxiosError) {
			createUserForm.serverError = true;

			return;
		}
	}
};

const breadcrumbs = ref<string[]>(['Система', 'Пользователи']);

const tableOptions = reactive({
	itemsPerPage: 5,
	headers: [
		{ title: 'E-mail', key: 'email', sortable: false },
		{ title: 'Имя', key: 'firstName', sortable: false },
		{ title: 'Роль', key: 'role', sortable: false },
		{ title: 'Статус', key: 'status', sortable: false },
		{ title: 'Дата создания', key: 'createdAt', sortable: false },
		{ title: 'Действия', key: 'actions', sortable: false },
	],
	serverItems: userList,
	loading: true,
	totalItems: 0,
	firstName: '',
	email: '',
});

const changeStatus = async (item: UserProfile) => {
	const isConfirm = confirm(`Вы уверены что хотите изменить статус пользователя ${item.email}?`);

	if (! isConfirm) {
		return;
	}

	switch (item.status) {
		case "blocked": {
			await userModel.activate(item.id);
			break;
		}
		case "active": {
			await userModel.block(item.id);
			break;
		}
	}

	await loadItems(currentLoadUsersOptions);
}

const userInfoDialogOpen = async (item: UserProfile) => {
	userInfoDialog.isShow = true;

	const user = <UserProfile>await UserFetcher.getUserById(item.id);

	userInfoDialog.user = [
		{attribute: 'Имя', value: user.firstName},
		{attribute: 'Фамилия', value: user.lastName},
		{attribute: 'Email', value: user.email},
		{attribute: 'Источник регистрации', value: user.registrationSource},
	];
}

const userInfoDialogClose = () => {
	userInfoDialog.isShow = false;
	userInfoDialog.user = null;
}

const loadItems = async (options: LoadParam | null) => {
	if (null === options) {
		return;
	}

	console.log(tableOptions.email);

	currentLoadUsersOptions = options;
	tableOptions.loading = true;

	const result = await UserFetcher.getUserList({
		offset: (options.page - 1) * options.itemsPerPage,
		limit: options.itemsPerPage,
		...(tableOptions.email.trim() && { email: tableOptions.email.trim() }),
		...(tableOptions.firstName.trim() && { firstName: tableOptions.firstName.trim() })
	}) as List

	tableOptions.serverItems = <UserProfile[]>result.data;
	tableOptions.totalItems = result.meta.total;
	tableOptions.loading = false;
};

watch(() => tableOptions.firstName, async(): Promise<void> => {
	const isFirstName = tableOptions.firstName.trim().length > 1

	if (!tableOptions.firstName.trim() || isFirstName) {
		await loadItems(currentLoadUsersOptions);
	}
})

watch(() => tableOptions.email, async (): Promise<void> => {
	const isEmail = FormHelper.emailPattern.test(
		tableOptions.email.trim()
	);

	if (!tableOptions.email.trim() || isEmail) {
		await loadItems(currentLoadUsersOptions);
	}
})

onMounted(() => {
	EmitterService.dispatchComponentOnMountedEvent();
});

</script>

<style scoped>

.sort-input {
	max-width: 350px;
}

</style>
