<template>
	<v-footer
		v-if="props.isDefaultSlotShow"
		:app="true"
		border
	>

		<v-container>
			<v-row justify="center" no-gutters>
				<v-col class="text-center">
					© {{ new Date().getFullYear() }} — <strong>Evening production</strong>
				</v-col>
			</v-row>
		</v-container>

	</v-footer>
</template>
<script setup lang="ts">

const props = defineProps({
	isDefaultSlotShow: null
});

</script>
<style scoped>
	.v-footer {
		max-height: 48px;
	}
</style>
