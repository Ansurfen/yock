package server

import "github.com/ansurfen/yock/daemon/interface/client"

type DaemonOption struct {
	*client.DaemonOption
}

var Gopt = &DaemonOption{client.Gopt}
