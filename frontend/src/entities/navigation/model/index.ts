import { defineStore } from 'pinia';

interface Drawer {
	drawer: boolean;
	rail: boolean;
}

export const useNavigationModel = defineStore({
	id: 'navigation',

	state: () => ({
		drawer: true,
		rail: false
	} as Drawer),

	actions: {
		close(): void {
			this.rail = true;
		},
		open(): void {
			this.rail = false;
		}
	},

	persist: true
});
