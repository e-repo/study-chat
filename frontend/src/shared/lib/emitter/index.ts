import mitt from 'mitt';

export const emitter = mitt();

export const COMPONENT_ON_MOUNTED_EVENT = 'component-on-mounted';

export const dispatchComponentOnMountedEvent = () => {
	emitter.emit(
		COMPONENT_ON_MOUNTED_EVENT
	);
};
