<template>

	<v-container fluid>
		<v-row>
			<v-col cols="12">
				<h1 class="text-center">Привет {{ userModel.userFromToken.firstName }}.</h1>
			</v-col>
		</v-row>
	</v-container>

</template>

<script setup lang="ts">
import {useUserModel} from '@/entities/user';
import {EmitterService} from '@/shared/lib';
import {onMounted} from 'vue';

const userModel = useUserModel();

onMounted(() => {
	EmitterService.dispatchComponentOnMountedEvent()
});

</script>

