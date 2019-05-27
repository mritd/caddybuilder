package conf

var DNSPlugins = []Plugin{
	{
		Name:    "exoscale",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/exoscale",
	},
	{
		Name:    "vscale",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/vscale",
	},
	{
		Name:    "acmedns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/acmedns",
	},
	{
		Name:    "ovh",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/ovh",
	},
	{
		Name:    "nifcloud",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/nifcloud",
	},
	{
		Name:    "duckdns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/duckdns",
	},
	{
		Name:    "auroradns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/auroradns",
	},
	{
		Name:    "gandiv5",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/gandiv5",
	},
	{
		Name:    "azure",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/azure",
	},
	{
		Name:    "gandi",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/gandi",
	},
	{
		Name:    "linode",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/linode",
	},
	{
		Name:    "rackspace",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/rackspace",
	},
	{
		Name:    "httpreq",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/httpreq",
	},
	{
		Name:    "dnsimple",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/dnsimple",
	},
	{
		Name:    "digitalocean",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/digitalocean",
	},
	{
		Name:    "pdns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/pdns",
	},
	{
		Name:    "lightsail",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/lightsail",
	},
	{
		Name:    "otc",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/otc",
	},
	{
		Name:    "cloudflare",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/cloudflare",
	},
	{
		Name:    "dnspod",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/dnspod",
	},
	{
		Name:    "ns1",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/ns1",
	},
	{
		Name:    "glesys",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/glesys",
	},
	{
		Name:    "godaddy",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/godaddy",
	},
	{
		Name:    "dyn",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/dyn",
	},
	{
		Name:    "selectel",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/selectel",
	},
	{
		Name:    "conoha",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/conoha",
	},
	{
		Name:    "stackpath",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/stackpath",
	},
	{
		Name:    "namedotcom",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/namedotcom",
	},
	{
		Name:    "vultr",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/vultr",
	},
	{
		Name:    "cloudxns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/cloudxns",
	},
	{
		Name:    "dnsmadeeasy",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/dnsmadeeasy",
	},
	{
		Name:    "googlecloud",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/googlecloud",
	},
	{
		Name:    "namecheap",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/namecheap",
	},
	{
		Name:    "generic",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/generic",
	},
	{
		Name:    "fastdns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/fastdns",
	},
	{
		Name:    "rfc2136",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/rfc2136",
	},
	{
		Name:    "linodev4",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/linodev4",
	},
	{
		Name:    "transip",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/transip",
	},
	{
		Name:    "route53",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/route53",
	},
}

var HttpPlugins = []Plugin{
	{
		Name:    "supervisor",
		Type:    "http",
		Package: "github.com/lucaslorentz/caddy-supervisor",
	},
	{
		Name:    "realip",
		Type:    "http",
		Package: "github.com/captncraig/caddy-realip",
	},
	{
		Name:    "git",
		Type:    "http",
		Package: "github.com/abiosoft/caddy-git",
	},
	{
		Name:    "proxyprotocol",
		Type:    "http",
		Package: "github.com/mastercactapus/caddy-proxyprotocol",
	},
	{
		Name:    "locale",
		Type:    "http",
		Package: "github.com/simia-tech/caddy-locale",
	},
	{
		Name:    "cache",
		Type:    "http",
		Package: "github.com/nicolasazrak/caddy-cache",
	},
	{
		Name:    "minify",
		Type:    "http",
		Package: "github.com/hacdias/caddy-minify",
	},
	{
		Name:    "geoip",
		Type:    "http",
		Package: "github.com/kodnaplakal/caddy-geoip",
	},
	{
		Name:    "authz",
		Type:    "http",
		Package: "github.com/casbin/caddy-authz",
	},
	{
		Name:    "filter",
		Type:    "http",
		Package: "github.com/echocat/caddy-filter",
	},
	{
		Name:    "ipfilter",
		Type:    "http",
		Package: "github.com/pyed/ipfilter",
	},
	{
		Name:    "ratelimit",
		Type:    "http",
		Package: "github.com/xuqingfeng/caddy-rate-limit",
	},
	{
		Name:    "expires",
		Type:    "http",
		Package: "github.com/epicagency/caddy-expires",
	},
	{
		Name:    "forwardproxy",
		Type:    "http",
		Package: "github.com/caddyserver/forwardproxy",
	},
	{
		Name:    "cors",
		Type:    "http",
		Package: "github.com/captncraig/cors/caddy",
	},
	{
		Name:    "s3browser",
		Type:    "http",
		Package: "github.com/techknowlogick/caddy-s3browser",
	},
	{
		Name:    "nobots",
		Type:    "http",
		Package: "github.com/Xumeiquer/nobots",
	},
	{
		Name:    "login",
		Type:    "http",
		Package: "github.com/tarent/loginsrv/caddy",
	},
	{
		Name:    "reauth",
		Type:    "http",
		Package: "github.com/freman/caddy-reauth",
	},
	{
		Name:    "extauth",
		Type:    "http",
		Package: "github.com/BTBurke/caddy-extauth",
	},
	{
		Name:    "jwt",
		Type:    "http",
		Package: "github.com/BTBurke/caddy-jwt",
	},
	{
		Name:    "jsonp",
		Type:    "http",
		Package: "github.com/pschlump/caddy-jsonp",
	},
	{
		Name:    "upload",
		Type:    "http",
		Package: "blitznote.com/src/caddy.upload",
	},
	{
		Name:    "multipass",
		Type:    "http",
		Package: "github.com/namsral/multipass/caddy",
	},
	{
		Name:    "datadog",
		Type:    "http",
		Package: "github.com/payintech/caddy-datadog",
	},
	{
		Name:    "prometheus",
		Type:    "http",
		Package: "github.com/miekg/caddy-prometheus",
	},
	{
		Name:    "pubsub",
		Type:    "http",
		Package: "github.com/jung-kurt/caddy-pubsub",
	},
	{
		Name:    "cgi",
		Type:    "http",
		Package: "github.com/jung-kurt/caddy-cgi",
	},
	{
		Name:    "filebrowser",
		Type:    "http",
		Package: "github.com/filebrowser/caddy",
	},
	{
		Name:    "webdav",
		Type:    "http",
		Package: "github.com/hacdias/caddy-webdav",
	},
	{
		Name:    "mailout",
		Type:    "http",
		Package: "github.com/SchumacherFM/mailout",
	},
	{
		Name:    "awses",
		Type:    "http",
		Package: "github.com/miquella/caddy-awses",
	},
	{
		Name:    "awslambda",
		Type:    "http",
		Package: "github.com/coopernurse/caddy-awslambda",
	},
	{
		Name:    "grpc",
		Type:    "http",
		Package: "github.com/pieterlouw/caddy-grpc",
	},
	{
		Name:    "gopkg",
		Type:    "http",
		Package: "github.com/zikes/gopkg",
	},
	{
		Name:    "restic",
		Type:    "http",
		Package: "github.com/restic/caddy",
	},
	{
		Name:    "wkd",
		Type:    "http",
		Package: "github.com/emersion/caddy-wkd",
	},
	{
		Name:    "dyndns",
		Type:    "http",
		Package: "github.com/linkonoid/caddy-dyndns",
	},
}

var OtherPlugins = []Plugin{
	{
		Name:    "net",
		Type:    "server",
		Package: "github.com/pieterlouw/caddy-net/caddynet",
	},
	{
		Name:    "dns",
		Type:    "server",
		Package: "github.com/coredns/coredns/core/dnsserver",
	},
	{
		Name:    "docker",
		Type:    "caddyfile",
		Package: "github.com/lucaslorentz/caddy-docker-proxy/plugin",
	},
	{
		Name:    "service",
		Type:    "hook",
		Package: "github.com/hacdias/caddy-service",
	},
}

var PluginsMap = map[string]Plugin{
	"exoscale": {
		Name:    "exoscale",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/exoscale",
	},
	"vscale": {
		Name:    "vscale",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/vscale",
	},
	"acmedns": {
		Name:    "acmedns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/acmedns",
	},
	"ovh": {
		Name:    "ovh",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/ovh",
	},
	"nifcloud": {
		Name:    "nifcloud",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/nifcloud",
	},
	"duckdns": {
		Name:    "duckdns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/duckdns",
	},
	"auroradns": {
		Name:    "auroradns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/auroradns",
	},
	"gandiv5": {
		Name:    "gandiv5",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/gandiv5",
	},
	"azure": {
		Name:    "azure",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/azure",
	},
	"gandi": {
		Name:    "gandi",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/gandi",
	},
	"linode": {
		Name:    "linode",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/linode",
	},
	"rackspace": {
		Name:    "rackspace",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/rackspace",
	},
	"httpreq": {
		Name:    "httpreq",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/httpreq",
	},
	"dnsimple": {
		Name:    "dnsimple",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/dnsimple",
	},
	"digitalocean": {
		Name:    "digitalocean",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/digitalocean",
	},
	"pdns": {
		Name:    "pdns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/pdns",
	},
	"lightsail": {
		Name:    "lightsail",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/lightsail",
	},
	"otc": {
		Name:    "otc",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/otc",
	},
	"cloudflare": {
		Name:    "cloudflare",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/cloudflare",
	},
	"dnspod": {
		Name:    "dnspod",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/dnspod",
	},
	"ns1": {
		Name:    "ns1",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/ns1",
	},
	"glesys": {
		Name:    "glesys",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/glesys",
	},
	"godaddy": {
		Name:    "godaddy",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/godaddy",
	},
	"dyn": {
		Name:    "dyn",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/dyn",
	},
	"selectel": {
		Name:    "selectel",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/selectel",
	},
	"conoha": {
		Name:    "conoha",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/conoha",
	},
	"stackpath": {
		Name:    "stackpath",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/stackpath",
	},
	"namedotcom": {
		Name:    "namedotcom",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/namedotcom",
	},
	"vultr": {
		Name:    "vultr",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/vultr",
	},
	"cloudxns": {
		Name:    "cloudxns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/cloudxns",
	},
	"dnsmadeeasy": {
		Name:    "dnsmadeeasy",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/dnsmadeeasy",
	},
	"googlecloud": {
		Name:    "googlecloud",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/googlecloud",
	},
	"namecheap": {
		Name:    "namecheap",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/namecheap",
	},
	"generic": {
		Name:    "generic",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/generic",
	},
	"fastdns": {
		Name:    "fastdns",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/fastdns",
	},
	"rfc2136": {
		Name:    "rfc2136",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/rfc2136",
	},
	"linodev4": {
		Name:    "linodev4",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/linodev4",
	},
	"transip": {
		Name:    "transip",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/transip",
	},
	"route53": {
		Name:    "route53",
		Type:    "dns",
		Package: "github.com/caddyserver/dnsproviders/route53",
	},

	"supervisor": {
		Name:    "supervisor",
		Type:    "http",
		Package: "github.com/lucaslorentz/caddy-supervisor",
	},
	"realip": {
		Name:    "realip",
		Type:    "http",
		Package: "github.com/captncraig/caddy-realip",
	},
	"git": {
		Name:    "git",
		Type:    "http",
		Package: "github.com/abiosoft/caddy-git",
	},
	"proxyprotocol": {
		Name:    "proxyprotocol",
		Type:    "http",
		Package: "github.com/mastercactapus/caddy-proxyprotocol",
	},
	"locale": {
		Name:    "locale",
		Type:    "http",
		Package: "github.com/simia-tech/caddy-locale",
	},
	"cache": {
		Name:    "cache",
		Type:    "http",
		Package: "github.com/nicolasazrak/caddy-cache",
	},
	"minify": {
		Name:    "minify",
		Type:    "http",
		Package: "github.com/hacdias/caddy-minify",
	},
	"geoip": {
		Name:    "geoip",
		Type:    "http",
		Package: "github.com/kodnaplakal/caddy-geoip",
	},
	"authz": {
		Name:    "authz",
		Type:    "http",
		Package: "github.com/casbin/caddy-authz",
	},
	"filter": {
		Name:    "filter",
		Type:    "http",
		Package: "github.com/echocat/caddy-filter",
	},
	"ipfilter": {
		Name:    "ipfilter",
		Type:    "http",
		Package: "github.com/pyed/ipfilter",
	},
	"ratelimit": {
		Name:    "ratelimit",
		Type:    "http",
		Package: "github.com/xuqingfeng/caddy-rate-limit",
	},
	"expires": {
		Name:    "expires",
		Type:    "http",
		Package: "github.com/epicagency/caddy-expires",
	},
	"forwardproxy": {
		Name:    "forwardproxy",
		Type:    "http",
		Package: "github.com/caddyserver/forwardproxy",
	},
	"cors": {
		Name:    "cors",
		Type:    "http",
		Package: "github.com/captncraig/cors/caddy",
	},
	"s3browser": {
		Name:    "s3browser",
		Type:    "http",
		Package: "github.com/techknowlogick/caddy-s3browser",
	},
	"nobots": {
		Name:    "nobots",
		Type:    "http",
		Package: "github.com/Xumeiquer/nobots",
	},
	"login": {
		Name:    "login",
		Type:    "http",
		Package: "github.com/tarent/loginsrv/caddy",
	},
	"reauth": {
		Name:    "reauth",
		Type:    "http",
		Package: "github.com/freman/caddy-reauth",
	},
	"extauth": {
		Name:    "extauth",
		Type:    "http",
		Package: "github.com/BTBurke/caddy-extauth",
	},
	"jwt": {
		Name:    "jwt",
		Type:    "http",
		Package: "github.com/BTBurke/caddy-jwt",
	},
	"jsonp": {
		Name:    "jsonp",
		Type:    "http",
		Package: "github.com/pschlump/caddy-jsonp",
	},
	"upload": {
		Name:    "upload",
		Type:    "http",
		Package: "blitznote.com/src/caddy.upload",
	},
	"multipass": {
		Name:    "multipass",
		Type:    "http",
		Package: "github.com/namsral/multipass/caddy",
	},
	"datadog": {
		Name:    "datadog",
		Type:    "http",
		Package: "github.com/payintech/caddy-datadog",
	},
	"prometheus": {
		Name:    "prometheus",
		Type:    "http",
		Package: "github.com/miekg/caddy-prometheus",
	},
	"pubsub": {
		Name:    "pubsub",
		Type:    "http",
		Package: "github.com/jung-kurt/caddy-pubsub",
	},
	"cgi": {
		Name:    "cgi",
		Type:    "http",
		Package: "github.com/jung-kurt/caddy-cgi",
	},
	"filebrowser": {
		Name:    "filebrowser",
		Type:    "http",
		Package: "github.com/filebrowser/caddy",
	},
	"webdav": {
		Name:    "webdav",
		Type:    "http",
		Package: "github.com/hacdias/caddy-webdav",
	},
	"mailout": {
		Name:    "mailout",
		Type:    "http",
		Package: "github.com/SchumacherFM/mailout",
	},
	"awses": {
		Name:    "awses",
		Type:    "http",
		Package: "github.com/miquella/caddy-awses",
	},
	"awslambda": {
		Name:    "awslambda",
		Type:    "http",
		Package: "github.com/coopernurse/caddy-awslambda",
	},
	"grpc": {
		Name:    "grpc",
		Type:    "http",
		Package: "github.com/pieterlouw/caddy-grpc",
	},
	"gopkg": {
		Name:    "gopkg",
		Type:    "http",
		Package: "github.com/zikes/gopkg",
	},
	"restic": {
		Name:    "restic",
		Type:    "http",
		Package: "github.com/restic/caddy",
	},
	"wkd": {
		Name:    "wkd",
		Type:    "http",
		Package: "github.com/emersion/caddy-wkd",
	},
	"dyndns": {
		Name:    "dyndns",
		Type:    "http",
		Package: "github.com/linkonoid/caddy-dyndns",
	},

	"net": {
		Name:    "net",
		Type:    "server",
		Package: "github.com/pieterlouw/caddy-net/caddynet",
	},
	"dns": {
		Name:    "dns",
		Type:    "server",
		Package: "github.com/coredns/coredns/core/dnsserver",
	},
	"docker": {
		Name:    "docker",
		Type:    "caddyfile",
		Package: "github.com/lucaslorentz/caddy-docker-proxy/plugin",
	},
	"service": {
		Name:    "service",
		Type:    "hook",
		Package: "github.com/hacdias/caddy-service",
	},
}
