<template>

	<v-container fluid>
		<v-row>
			<v-col cols="12">
				<h1 class="text-center">Привет {{ userModel.userFromToken.firstName }}.</h1>
				<h3 class="text-decoration-underline">Админка это:</h3>
			</v-col>
		</v-row>
		<v-row>
			<v-col xs="12" sm="3" md="4">
				<v-card>
					<v-card-item>
						<v-card-title>Системная информация</v-card-title>
						<v-card-subtitle>Сервисные разделы</v-card-subtitle>
					</v-card-item>

					<v-card-text>
						Сервисные разделы админ-панели состоящие из обобщенной информации всех разделов (модулей) проекта. Админ-панель предоставляет возможность управлять данным контентом. К нему например относятся пользователи сервисов
					</v-card-text>
				</v-card>
			</v-col>
			<v-col xs="12" sm="3" md="4">
				<v-card>
					<v-card-item>
						<v-card-title>Блог</v-card-title>
						<v-card-subtitle>Набор статей по теме защиты растений</v-card-subtitle>
					</v-card-item>

					<v-card-text>
						Классический блог представляющий пользователю исчерпывающую информацию по тематике защите растений. Состоит из множества статей по теме разбитых по категориям. Каждая статья дополнительно еще помечена метками, которые представяют собой возможность объединять статьи в более мелкие группы.
					</v-card-text>
				</v-card>
			</v-col>
			<v-col xs="12" sm="3" md="4">
				<v-card>
					<v-card-item>
						<v-card-title>Интернет магазин</v-card-title>
						<v-card-subtitle>Сопутствующие товары для защиты растений</v-card-subtitle>
					</v-card-item>

					<v-card-text>
						Представяет собой набор товаров разбитый по категориям для защиты растений. Основным функционалом является покупка в один клик, возможность бесплатной доставки, аналитика продаж, консультации при покупке, и интеграция с 1с.
					</v-card-text>
				</v-card>
			</v-col>
		</v-row>
	</v-container>

</template>

<script setup lang="ts">
import { useUserModel } from '@/entities/user';
import { EmitterService } from '@/shared/lib';
import { onMounted } from 'vue';

const userModel = useUserModel();

onMounted(() => {
	EmitterService.dispatchComponentOnMountedEvent()
});

</script>

<style scoped></style>
