// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha3

import (
	cors "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/cors/v2"
	grpcweb "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/grpc_web/v2"
	router "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/router/v2"
	httpinspector "github.com/envoyproxy/go-control-plane/envoy/config/filter/listener/http_inspector/v2"
	originaldst "github.com/envoyproxy/go-control-plane/envoy/config/filter/listener/original_dst/v2"
	originalsrc "github.com/envoyproxy/go-control-plane/envoy/config/filter/listener/original_src/v2alpha1"
	tlsinspector "github.com/envoyproxy/go-control-plane/envoy/config/filter/listener/tls_inspector/v2"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	fault "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/fault/v3"
	http_conn "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"

	"istio.io/istio/pilot/pkg/networking/util"
	alpn_filter "istio.io/istio/security/proto/envoy/config/filter/http/alpn/v2alpha1"
)

const OriginalSrc = "envoy.listener.original_src"

// Define static filters to be reused across the codebase. This avoids duplicate marshaling/unmarshaling
// This should not be used for filters that will be mutated
var (
	corsFilter = &http_conn.HttpFilter{
		Name: wellknown.CORS,
		ConfigType: &http_conn.HttpFilter_TypedConfig{
			TypedConfig: util.MessageToAny(&cors.Cors{}),
		},
	}
	faultFilter = &http_conn.HttpFilter{
		Name: wellknown.Fault,
		ConfigType: &http_conn.HttpFilter_TypedConfig{
			TypedConfig: util.MessageToAny(&fault.HTTPFault{}),
		},
	}
	routerFilter = &http_conn.HttpFilter{
		Name: wellknown.Router,
		ConfigType: &http_conn.HttpFilter_TypedConfig{
			TypedConfig: util.MessageToAny(&router.Router{}),
		},
	}
	grpcWebFilter = &http_conn.HttpFilter{
		Name: wellknown.GRPCWeb,
		ConfigType: &http_conn.HttpFilter_TypedConfig{
			TypedConfig: util.MessageToAny(&grpcweb.GrpcWeb{}),
		},
	}

	tlsInspectorFilter = &listener.ListenerFilter{
		Name: wellknown.TlsInspector,
		ConfigType: &listener.ListenerFilter_TypedConfig{
			TypedConfig: util.MessageToAny(&tlsinspector.TlsInspector{}),
		},
	}
	httpInspectorFilter = &listener.ListenerFilter{
		Name: wellknown.HttpInspector,
		ConfigType: &listener.ListenerFilter_TypedConfig{
			TypedConfig: util.MessageToAny(&httpinspector.HttpInspector{}),
		},
	}
	originalDestinationFilter = &listener.ListenerFilter{
		Name: wellknown.OriginalDestination,
		ConfigType: &listener.ListenerFilter_TypedConfig{
			TypedConfig: util.MessageToAny(&originaldst.OriginalDst{}),
		},
	}
	originalSrcFilter = &listener.ListenerFilter{
		Name: OriginalSrc,
		ConfigType: &listener.ListenerFilter_TypedConfig{
			TypedConfig: util.MessageToAny(&originalsrc.OriginalSrc{}),
		},
	}
	alpnFilter = &http_conn.HttpFilter{
		Name: AlpnFilterName,
		ConfigType: &http_conn.HttpFilter_TypedConfig{
			TypedConfig: util.MessageToAny(&alpn_filter.FilterConfig{
				AlpnOverride: []*alpn_filter.FilterConfig_AlpnOverride{
					{
						UpstreamProtocol: alpn_filter.FilterConfig_HTTP10,
						AlpnOverride:     mtlsHTTP10ALPN,
					},
					{
						UpstreamProtocol: alpn_filter.FilterConfig_HTTP11,
						AlpnOverride:     mtlsHTTP11ALPN,
					},
					{
						UpstreamProtocol: alpn_filter.FilterConfig_HTTP2,
						AlpnOverride:     mtlsHTTP2ALPN,
					},
				},
			}),
		},
	}
)
