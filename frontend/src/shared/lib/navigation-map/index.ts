interface NavigationParam {
	value: string,
	to: string | undefined
}

type NavigationKey = 'service_home' | 'service_users' | 'blog_posts' | 'blog_categories';

type NavigationMap = {
	[key in NavigationKey]: NavigationParam
}

const navMap: NavigationMap = {
	'service_home': {value: 'service-home', to: '/'},
	'service_users': {value: 'service-users', to: '/auth/users'},
	'blog_posts': {value: 'blog-posts', to: '/blog/posts'},
	'blog_categories': {value: 'blog-categories', to: '/blog/categories'},
};

export const getParam = (id: string): NavigationParam => {
	return navMap[id as NavigationKey];
};
