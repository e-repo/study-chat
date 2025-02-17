<template>
	<v-container class="h-100">
		<v-row class="h-100 align-center">
			<v-col class="d-flex justify-center">

				<div>
					<h1>Доступ к ресурсу запрещен!</h1>
					<div class="d-flex justify-center mt-2">
						<v-btn
							to="/login"
							color="success"
							type="button"
						>
							Вход
						</v-btn>
					</div>
				</div>

			</v-col>

		</v-row>
	</v-container>
</template>

<script setup lang="ts">
</script>
<style scoped></style>
