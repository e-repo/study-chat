<template>

	<v-app-bar v-if="props.isDefaultSlotShow">
		<template v-slot:prepend>
			<v-app-bar-nav-icon
				@click="toggleNavigation"
			></v-app-bar-nav-icon>
		</template>
		<v-app-bar-title>Чат</v-app-bar-title>
		<template v-slot:append>
			<v-menu
				transition="slide-y-transition"
			>
				<template v-slot:activator="{ props }">
					<v-btn
						icon="mdi-account"
						variant="outlined"
						v-bind="props"
						class="text-grey-darken-1"
					></v-btn>
				</template>
				<v-card>
					<v-list>
						<v-list-item
							prepend-icon="mdi-account-eye-outline"
							:subtitle="userModel.userFromToken.email"
							:title="userModel.userFromToken.firstName"
						>
						</v-list-item>
					</v-list>

					<v-divider></v-divider>

					<v-list>
						<v-list-item
							v-for="(item, i) in items"
							:key="i"
							:value="item"
							color="primary"
						>
							<template v-slot:prepend>
								<v-icon :icon="item.icon"></v-icon>
							</template>

							<v-list-item-title v-text="item.text"></v-list-item-title>
						</v-list-item>
					</v-list>

					<v-divider></v-divider>

					<v-card-actions>
						<v-spacer></v-spacer>

						<v-btn
							prepend-icon="mdi-logout"
							variant="text"
							@click="logout"
						>
							Выход
						</v-btn>
					</v-card-actions>
				</v-card>
			</v-menu>
		</template>
	</v-app-bar>

</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useUserModel } from '@/entities/user';
import { useRouter } from 'vue-router';
import { useNavigationModel } from '@/entities/navigation';

const userModel = useUserModel();
const navigationModel = useNavigationModel();
const router = useRouter();

const items = ref([
	{ text: 'Мои файлы', icon: 'mdi-folder' },
	{ text: 'Недавнее', icon: 'mdi-history' },
	{ text: 'Загрузки', icon: 'mdi-upload' },
]);

const props = defineProps({
	isDefaultSlotShow: null
});

const logout = () => {
	userModel.logout();

	router.push({name: 'Login'});
}

const toggleNavigation = (): void => {
	navigationModel.rail = ! navigationModel.rail;
}

</script>

<style scoped></style>
