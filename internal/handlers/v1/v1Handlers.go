package handlers

type v1Handlers struct {
	urls *URLHandlers
}

func GetV1Handlers() *v1Handlers {
	return &v1Handlers{
		urls: &URLHandlers{},
	}
}
