<template>
	<v-layout ref="app" class="flex-column">
		<slot name="header"></slot>
		<slot name="navigation"></slot>

		<v-main>
			<slot />
		</v-main>

		<slot name="footer"></slot>
	</v-layout>
</template>

<script setup lang="ts"></script>

<style scoped></style>
