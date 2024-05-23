package main

import (
	"context"
	"fmt"
	"github.com/bighuangbee/kratos-vue-admin/api/admin/v1"
	"github.com/bighuangbee/kratos-vue-admin/pkg/util/kitGrpc"
	"time"
)

func main() {

	//var discovery *nacos.Registry
	//var filter selector.NodeFilter
	//var naming naming_client.INamingClient
	//
	//naming, err := kitGrpc.NewNamingClient(
	//	bc.Discovery.IPAddr,
	//	uint64(bc.Discovery.Port),
	//	util.WithLogLevel(bc.LogLevel),
	//	util.WithNamespaceID(bc.Env),
	//)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if true {
	//	discovery = nacos.New(naming)
	//	filter = util.FilterVersion(bc.Version)
	//}

	opt := kitGrpc.GetConnOption{
		Endpoint: "localhost:9001",
		Logger:   nil,
		//Discovery:    discovery,
		//SelectFilter: filter,
		Caller:  "bc.Name",
		Timeout: time.Second * time.Duration(3),
	}

	cli, err := kitGrpc.GetGrpcClient(context.Background(), opt)
	if err != nil {
		fmt.Println(err)
		return
	}

	sysUserCli := v1.NewSysuserClient(cli)
	reply, err := sysUserCli.ListSysuser(context.Background(), &v1.ListSysuserRequest{
		//PageNum:  0,
		//PageSize: 0,
		//Username: "",
		//Phone:    "",
		//Status:   0,
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*reply)
	}

}
