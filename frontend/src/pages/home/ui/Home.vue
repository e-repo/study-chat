<template>

	<v-container fluid>
		<v-row>
			<v-col cols="12">
        <div class="chat-container">
          <!-- Чат-сообщения -->
          <v-card class="chat-message received" v-for="msg in messages" :key="msg.id">
            <v-card-text>{{ msg.text }}</v-card-text>
          </v-card>

          <v-textarea
              v-model="newMessage"
              label="Введите сообщение"
              outlined
              :counter="500"
              maxlength="500"
              rows="2"
              @keyup.enter="sendMessage"
          ></v-textarea>
        </div>
			</v-col>
		</v-row>
	</v-container>

</template>

<script setup lang="ts">
import {EmitterService} from '@/shared/lib';
import {onMounted, reactive, ref} from 'vue';

const messages = reactive([
  { id: 1, text: 'Привет!' },
  { id: 2, text: 'Как дела?' }
]);

const newMessage = ref("")

const sendMessage = () => {
  if (newMessage.value.length > 0) {
    messages.push({
      id: messages.length + 1,
      text: newMessage.value
    });
    newMessage.value = '';
  }
}

onMounted(() => {
	EmitterService.dispatchComponentOnMountedEvent()
});

</script>

