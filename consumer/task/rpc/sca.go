package rpc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"io"
	"net/http"
)

type ScaClient struct {
	url string
}

type ScaConfig struct {
	Url string
}

func InitScaClient(cfg *ScaConfig) ScaClient {
	return ScaClient{
		url: cfg.Url,
	}
}

type DepsResponse struct {
	Packages PackageDeps `json:"packages"`
}

type PackageDeps struct {
	Modules []ModuleDeps `json:"modules"`
}

type ModuleDeps struct {
	Purl         string `json:"purl"`
	Lang         string `json:"lang"`
	Path         string `json:"path"`
	Dependencies []Dep  `json:"dependencies"`
}

type Dep struct {
	Purl  string
	Limit string
	Scope string
}

func (c *ScaClient) GetDeps(p string, lang string) (*DepsResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%v/search?package=%v&lang=%v", c.url, p, lang))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	deps := &DepsResponse{}
	err = jsonx.Unmarshal(body, deps)
	if err != nil {
		return nil, err
	}
	return deps, nil
}
