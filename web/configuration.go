package web

// GinServerConfig 结构表示gin服务容器的配置
type GinServerConfig struct {
	NoMethodStatus   int    // ${server.no-method-status}, default=404
	NoMethodPage     string // ${server.no-method-page}, like 'resource:///www/http404.html'
	NoResourceStatus int    // ${server.no-resource-status}, default=404
	NoResourcePage   string // ${server.no-resource-page}, like 'resource:///www/http404.html'

	IndexPageNames        string // ${server.index-page-names}, like 'index.html, index.htm, index.xxx'
	StaticResourceFolder  string // ${server.static-resource-folder}, like 'resource:///static/'
	ContentTypeProperties string // ${server.content-type-properties}, like 'resource:///types.properties'
	APIContextPath        string // ${server.api-context-path}, like ''| '/' | '/appname'
	StaticContextPath     string // ${server.static-context-path}, like ''| '/' | '/appname'

	Host string // ${server.host}, like '0.0.0.0'
	Port int    // ${server.port}, like '8080'
}
