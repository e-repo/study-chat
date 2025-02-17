import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';
import checker from 'vite-plugin-checker';
import path from 'path';

export default defineConfig({
	plugins: [
		vue(),
		checker({
			eslint: {
				lintCommand: 'eslint "./**/*.{ts,tsx}"',
			},
		}),
	],
	server: {
		port: 3000,
		strictPort: true,
		host: true
	},
	base: '',
	resolve: {
		alias: {
			'@': path.resolve(__dirname, './src'),
		}
	}
});
