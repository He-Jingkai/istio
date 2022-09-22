package utils

import "github.com/containernetworking/cni/pkg/types"

type ProxyArgs struct {
	/* added by hjk, for offMesh only */
	PROXY_POD_NAME               types.UnmarshallableString // nolint: revive, stylecheck
	PROXY_POD_NAMESPACE          types.UnmarshallableString // nolint: revive, stylecheck
	PROXY_POD_INFRA_CONTAINER_ID types.UnmarshallableString // nolint: revive, stylecheck
	/* offMesh addition end */
}
