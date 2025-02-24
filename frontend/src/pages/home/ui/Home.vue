<template>

	<v-container fluid>
		<v-row>
			<v-col cols="12">
        <div class="chat-container">
          <!-- Чат-сообщения -->
          <div
              class="chat-message-list bg-green-lighten-4"
              ref="chatContainer"
          >
            <div class="d-flex flex-column justify-end">
              <div
                  v-for="msg in messages"
                  :key="msg.id"
              >
                <v-card
                    class="chat-message received ma-2"
                    max-width="400"
                >
                  <v-card-item>
                    <v-card-subtitle>Пользователь - {{ msg.id }}</v-card-subtitle>
                  </v-card-item>
                  <v-card-text>{{ msg.text }}</v-card-text>
                </v-card>
              </div>
            </div>
          </div>
          <div class="mt-2">
            <v-textarea
                max
                v-model="newMessage"
                label="Введите сообщение"
                variant="outlined"
                class="mt-auto"
                :counter="500"
                maxlength="500"
                max-rows="1"
                rows="1"
                @keyup.enter="sendMessage"
            ></v-textarea>
          </div>
        </div>
			</v-col>
		</v-row>
	</v-container>

</template>

<script setup lang="ts">
import {EmitterService} from '@/shared/lib';
import {nextTick, onMounted, reactive, ref, watch} from 'vue';

const messages = reactive([
  { id: 1, text: 'Привет!' },
  { id: 2, text: 'О, привет!' },
  { id: 2, text: 'Как дела?' },
  { id: 1, text: 'Норм бро' },
  { id: 1, text: 'А у тебя?' },
  { id: 2, text: 'сойдет' },
  { id: 2, text: 'Че там нового в Эфке' },
]);

const newMessage = ref("")

const chatContainer = ref<HTMLElement | null>(null);

const scrollToBottom = () => {
  nextTick(() => {
    if (chatContainer.value) {
      chatContainer.value.scrollTop = chatContainer.value.scrollHeight;
    }
  });
};

const sendMessage = () => {
  if (newMessage.value.length > 0) {
    messages.push({
      id: messages.length + 1,
      text: newMessage.value
    });
    newMessage.value = '';
  }
}

watch(
    () => messages.length,
    () => {
      scrollToBottom();
    }
);

onMounted(() => {
  scrollToBottom();
	EmitterService.dispatchComponentOnMountedEvent()
});

</script>

<style scoped>
.chat-container {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 142px);
}

.chat-message-list {
  max-height: calc(100vh - 200px);
  overflow-y: auto
}
</style>

