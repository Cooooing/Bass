package model

type Convert[T any] interface {
	ConvertToRpc() T
}

func ConvertToRpc[T1 Convert[T2], T2 any](model T1) T2 {
	return model.ConvertToRpc()
}

func ConvertToRpcList[T1 Convert[T2], T2 any](models []T1) []T2 {
	rsp := make([]T2, len(models))
	for i, model := range models {
		rsp[i] = model.ConvertToRpc()
	}
	return rsp
}
