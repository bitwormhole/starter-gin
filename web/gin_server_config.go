package web

type GinServerConfig struct {
	ErrorPageName         string // 只能位于根目录， 例如：‘/error.html’
	IndexPageNames        string // index.html, index.htm, index.xxx
	ContentTypeProperties string // like 'resources:///types.properties'
	ContextPath           string
	Host                  string
	Port                  int
	NotFoundStatusCode    int // default=404
}
