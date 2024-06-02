package rpc

type AlgoService interface {
	SubmitAlgo(jar string, entryPoint string, args map[string]interface{}) (string, error)

	StopAlgo(appID string) error
}
