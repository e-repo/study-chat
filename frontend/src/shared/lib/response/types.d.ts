interface ViolationItem {
	detail: string | null;
	source: string;
	data: object;
}

export interface Violation {
	message: string;
	errors: ViolationItem[];
}

export interface ListMeta {
	limit: number;
	offset: number;
	total: number;
}

export interface List {
	data: object;
	meta: ListMeta;
}
