import * as NavigationApi from '../api';

export const NavigationFetcher = {
	async getServiceMenuItems() {
		return await NavigationApi.fetchServiceMenuItems();
	},
	async getBlogMenuItems() {
		return await NavigationApi.fetchBlogMenuItems();
	},
};
