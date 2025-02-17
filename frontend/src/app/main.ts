import { createApp } from 'vue';
import App from './App.vue';
import router from './router/router.ts';
import { createPinia } from 'pinia';
import piniaPluginPersistedState from 'pinia-plugin-persistedstate';

// Vuetify
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

import '@mdi/font/css/materialdesignicons.css';
import { aliases, mdi } from 'vuetify/iconsets/mdi';
import { ru } from 'vuetify/locale';

const pinia = createPinia();

const vuetify = createVuetify({
	locale: {
		locale: 'ru',
		messages: { ru }
	},
	theme: {
	},
	icons: {
		defaultSet: 'mdi',
		aliases,
		sets: {
			mdi,
		},
	},
	components,
	directives,
});

createApp(App)
	.use(router)
	.use(
		pinia.use(piniaPluginPersistedState)
	)
	.use(vuetify)
	.mount('#app');
