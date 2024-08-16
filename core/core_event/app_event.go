package core_event

import (
	"container/list"
	"unsafe"
)

type AppEventHandler func(event string, args ...any) error

type AppEventContext struct {
	items map[string]*list.List
}

var appEventContext = &AppEventContext{
	items: make(map[string]*list.List),
}

func GetAppEventContext() *AppEventContext {
	return appEventContext
}

func (context *AppEventContext) On(event string, handler AppEventHandler) {
	value, ok := context.items[event]

	if !ok {
		value = list.New()
		context.items[event] = value
	}

	value.PushBack(handler)
}

func (context *AppEventContext) Emit(event string, args ...any) error {
	value, ok := context.items[event]

	if !ok {
		return nil
	}

	for i := value.Front(); i != nil; i = i.Next() {
		err := (*((*AppEventHandler)(unsafe.Pointer(&i.Value))))(event, args)

		if err != nil {
			return err
		}
	}

	return nil
}
