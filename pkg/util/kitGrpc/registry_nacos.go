package kitGrpc

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport"
	"os"

	nodeSelector "github.com/go-kratos/kratos/v2/selector"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type ClientOption func(o *constant.ClientConfig)

func processNacosClientParam(ipAddr string, port uint64, opts ...ClientOption) vo.NacosClientParam {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(ipAddr, port),
	}

	path := fmt.Sprintf("/tmp/nacos/%d/", os.Getpid())
	cc := constant.ClientConfig{
		TimeoutMs:           5000,
		NamespaceId:         "public",
		NotLoadCacheAtStart: true,
		LogLevel:            "error",
		LogDir:              path + "log",
		CacheDir:            path + "cache",
	}

	for _, o := range opts {
		o(&cc)
	}
	return vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	}
}

func NewNamingClient(ipAddr string, port uint64, opts ...ClientOption) (naming_client.INamingClient, error) {
	return clients.NewNamingClient(processNacosClientParam(ipAddr, port, opts...))
}

func NewConfigClient(ipAddr string, port uint64, opts ...ClientOption) (config_client.IConfigClient, error) {
	return clients.NewConfigClient(processNacosClientParam(ipAddr, port, opts...))
}

func WithTimeout(timeout uint64) ClientOption {
	return func(o *constant.ClientConfig) {
		if timeout > 0 {
			o.TimeoutMs = timeout * 1e6
		}
	}
}

func WithNamespaceID(ns string) ClientOption {
	return func(o *constant.ClientConfig) {
		if ns != "" {
			o.NamespaceId = ns
		}
	}
}

func WithLogLevel(lv string) ClientOption {
	return func(o *constant.ClientConfig) {
		if lv != "" {
			o.LogLevel = lv
		}
	}
}

////

const (
	BaseVersion        = "latest"
	VersionKeyInHeader = "x-md-global-version"
)

func FilterVersion(version string) nodeSelector.NodeFilter {
	return func(ctx context.Context, nodes []nodeSelector.Node) []nodeSelector.Node {
		if tr, ok := transport.FromServerContext(ctx); ok {
			version = tr.RequestHeader().Get(VersionKeyInHeader)
		}
		filterFn := func(match string) []nodeSelector.Node {
			var filters []nodeSelector.Node
			for _, n := range nodes {
				if n.Version() == match {
					filters = append(filters, n)
				}
			}
			return filters
		}

		r := filterFn(version)
		// 降级
		if len(r) == 0 && version != BaseVersion {
			r = filterFn(BaseVersion)
		}
		return r
	}
}
