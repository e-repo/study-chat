<template>

	<v-navigation-drawer
		v-if="props.isDefaultSlotShow"
		v-model="navigationModel.drawer"
		:rail="navigationModel.rail"
		permanent
	>

		<v-list density="compact">
			<v-list-subheader
				class="mt-8 text-decoration-underline"
			>
				<v-icon icon="mdi-cog-outline" class="mr-1"/>Система:
			</v-list-subheader>

			<v-divider></v-divider>

			<v-list-item
				v-for="(item, i) in serviceMenuItems"
				:key="i"
				:prepend-icon="`mdi-${item.icon}`"
				:title="item.title"
				:to="NavigationMap.getParam(item.id).to"
				:value="NavigationMap.getParam(item.id).value"
			></v-list-item>

			<v-list-subheader
				class="mt-8 text-decoration-underline"
			>
				<v-icon icon="mdi-text-recognition" class="mr-1"/>Блог:
			</v-list-subheader>

			<v-divider></v-divider>

			<v-list-item
				v-for="(item, i) in blogMenuItems"
				:key="i"
				:prepend-icon="`mdi-${item.icon}`"
				:title="item.title"
				:to="NavigationMap.getParam(item.id).to"
				:value="NavigationMap.getParam(item.id).value"
			></v-list-item>
		</v-list>
	</v-navigation-drawer>

</template>

<script setup lang="ts">
import { useNavigationModel } from '@/entities/navigation';
import { onMounted, ref } from 'vue';
import { useRefreshTokenListener } from '@/entities/user';
import { NavigationMap } from '@/shared/lib';
import { NavigationFetcher, MenuItem } from '@/features/navigation';

useRefreshTokenListener();

const navigationModel = useNavigationModel();

let serviceMenuItems = ref<MenuItem[]>([]);
let blogMenuItems = ref<MenuItem[]>([]);

const props = defineProps({
	isDefaultSlotShow: null
});

onMounted(async () => {
	serviceMenuItems.value = (await NavigationFetcher.getServiceMenuItems()).data;
	blogMenuItems.value = (await NavigationFetcher.getBlogMenuItems()).data;
});

</script>

<style scoped></style>
