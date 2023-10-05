package rpc

import (
	"context"
	"fmt"
	"github.com/3-shake/livy-go"
	"github.com/zeromicro/go-zero/core/jsonx"
	"net/http"
	"net/url"
	"strings"
)

type livyClient struct {
	svc    *livy.Service
	master string
}

type LivyConfig struct {
	Addr          string
	SparkMasterUI string
}

func InitLivyClient(cfg *LivyConfig) AlgoService {
	svc := livy.NewService(context.Background())
	svc.BasePath = cfg.Addr
	return &livyClient{svc: svc, master: cfg.SparkMasterUI}
}

func (l *livyClient) SubmitAlgo(jar string, entryPoint string, args map[string]interface{}) (string, error) {
	argsJSON, err := jsonx.MarshalToString(args)
	argsJSON = fmt.Sprintf("{%v}", strings.Replace(argsJSON[1:len(argsJSON)-1], "\"", "\\\"", -1))
	if err != nil {
		return "", err
	}
	batch := l.svc.Batches.Insert(&livy.InsertBatchRequest{
		File:      jar,
		ClassName: entryPoint,
		Args:      []string{argsJSON},
		Conf: map[string]string{
			"spark.driver.extraJavaOptions": "--add-opens=java.base/java.lang.invoke=ALL-UNNAMED --add-opens=java.base/java.nio=ALL-UNNAMED --add-opens=java.base/java.util=ALL-UNNAMED --add-opens=java.base/sun.nio.ch=ALL-UNNAMED",
		},
	})
	res, err := batch.Do()
	if err != nil {
		return "", err
	}
	return res.AppID, nil
}

func (l *livyClient) StopAlgo(appID string) error {
	urlValues := url.Values{}
	urlValues.Add("id", appID)
	urlValues.Add("terminate", "true")
	_, err := http.Post(l.master, "id=${app-id}&terminate=true", strings.NewReader(urlValues.Encode()))
	return err
}
