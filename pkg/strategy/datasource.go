package strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"prometheus-manager/pkg/httpx"
)

var _ Query = (*Datasource)(nil)

type Datasource struct {
	// 数据源类型
	Category string
	// 地址
	Domain string
}

// Query 调用数据源查询数据
//
//	curl 'https://<domain>/api/v1/query?query=go_memstats_sys_bytes&time=1704785907'
func (d *Datasource) Query(_ context.Context, expr string, duration int64) (*QueryResponse, error) {
	params, err := ParseQuery(map[string]any{
		"query": expr,
		"time":  duration,
	})
	if err != nil {
		return nil, err
	}

	hx := httpx.NewHttpX()
	hx.SetHeader(map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9",
	})
	getResponse, err := hx.GET(fmt.Sprintf("%s%s?%s", d.Domain, apiV1Query, params))
	if err != nil {
		return nil, err
	}
	defer getResponse.Body.Close()
	var allResp QueryResponse
	if err = json.NewDecoder(getResponse.Body).Decode(&allResp); err != nil {
		return nil, err
	}
	return &allResp, nil
}

// NewDatasource 实例化数据源对象
func NewDatasource(category, domain string) *Datasource {
	return &Datasource{
		Category: category,
		Domain:   strings.TrimSuffix(domain, "/"),
	}
}