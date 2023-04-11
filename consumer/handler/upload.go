package handler

type Upload struct {
}

func (h *Upload) Handle(params string) (string, error) {
	// TODO: 需要先上线一个文件存储服务，或者阿里云OSS。然后修改请求的参数为文件路径。
	return "", nil
}
