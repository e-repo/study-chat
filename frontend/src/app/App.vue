<template>
	<v-app>
		<component
			:is="layout"
		>
			<template v-slot:header>
				<Header :is-default-slot-show="isDefaultSlotShow" />
			</template>

			<template v-slot:navigation>
				<Navigation :is-default-slot-show="isDefaultSlotShow" />
			</template>

			<template v-slot:default>
				<router-view />
			</template>

			<template v-slot:footer>
				<Footer :is-default-slot-show="isDefaultSlotShow" />
			</template>
		</component>
	</v-app>
</template>


<script setup lang="ts">
import { useRoute } from 'vue-router';
import { computed, ref } from 'vue';
import { DefaultLayout } from '@/shared/ui/layout';
import { Header } from '@/widgets/header';
import { Navigation } from '@/widgets/navigation';
import { Footer } from '@/widgets/footer';
import { EmitterService } from '@/shared/lib';

let isDefaultSlotShow = ref<boolean>(false);

const route = useRoute();

const layout = computed(() => route.meta.layout || DefaultLayout);

EmitterService.emitter.on(
	EmitterService.COMPONENT_ON_MOUNTED_EVENT,
	() => {
		isDefaultSlotShow.value = true;
	}
);
</script>

<style></style>
