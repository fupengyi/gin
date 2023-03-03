package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/render"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"time"
)

Gin 是一个用 Go 编写的 Web 框架。它具有类似 martini 的 API，由于 httprouter，性能最高可提高 40 倍。如果您需要性能和良好的生产力，您会喜欢 Gin。
Gin的主要特点是： 零配置路由器 快速地 中间件支持 无崩溃 JSON 验证 Routes分组 错误管理 渲染内置 可扩展


Getting started
Prerequisites
	Go：三个最新主要版本中的任何一个（现在需要 1.16+ 版本）。


Getting Gin
使用 Go 模块支持，只需添加以下导入
	import "github.com/gin-gonic/gin"
到你的代码，然后go [build|run|test] 将自动获取必要的依赖项。
否则，运行以下 Go 命令来安装 gin 包：
	$ go get -u github.com/gin-gonic/gin


Running Gin
首先你需要导入 Gin 包来使用 Gin，一个最简单的例子就像下面的 example.go：
package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
并使用 Go 命令运行演示：
	# run example.go and visit 0.0.0.0:8080/ping on browser
	$ go run example.go
	

Learn more examples
Quick Start
	学习和实践更多示例，请阅读包含 API 示例和构建标签的 Gin 快速入门。
Examples
	在 Gin 示例存储库中展示了 Gin 的各种用例的许多现成的示例。


Documentation
请参阅 API 文档和包说明。


Middlewares
你可以在 gin-contrib 找到很多有用的 Gin 中间件。


Users
使用 Gin Web 框架的很棒的项目列表。
gorush：用 Go 编写的推送通知服务器。
fnproject：容器原生、云不可知的无服务器平台。
photoprism：由 Go 和 Google TensorFlow 提供支持的个人照片管理。
lura：带有中间件的超高性能 API 网关。
picfit：用 Go 编写的图像大小调整服务器。
dkron：分布式容错作业调度系统。


Documentation ¶
Overview ¶
包 gin 实现了一个名为 gin 的 HTTP Web 框架。
有关 gin 的更多信息，请参阅 https://gin-gonic.com/。


Constants ¶
View Source
const (		// Content-Type MIME 最常见的数据格式。
	MIMEJSON              = binding.MIMEJSON
	MIMEHTML              = binding.MIMEHTML
	MIMEXML               = binding.MIMEXML
	MIMEXML2              = binding.MIMEXML2
	MIMEPlain             = binding.MIMEPlain
	MIMEPOSTForm          = binding.MIMEPOSTForm
	MIMEMultipartPOSTForm = binding.MIMEMultipartPOSTForm
	MIMEYAML              = binding.MIMEYAML
	MIMETOML              = binding.MIMETOML
)

const (		// 可信平台
	// PlatformGoogleAppEngine when running on Google App Engine. Trust X-Appengine-Remote-Addr		// 在 Google App Engine 上运行时的 PlatformGoogleAppEngine。信任 X-Appengine-Remote-Addr
	// for determining the client's IP																// 用于确定客户端的 IP
	PlatformGoogleAppEngine = "X-Appengine-Remote-Addr"
	// PlatformCloudflare when using Cloudflare's CDN. Trust CF-Connecting-IP for determining		// 使用 Cloudflare 的 CDN 时的 PlatformCloudflare。信任 CF-Connecting-IP 来确定
	// the client's IP																				// 客户端的IP
	PlatformCloudflare = "CF-Connecting-IP"
)

const (
	// DebugMode indicates gin mode is debug.					// DebugMode 表示 gin 模式是 debug。
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.				// ReleaseMode 表示 gin 模式是 release。
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.						// TestMode 表示 gin 模式是 test。
	TestMode = "test"
)

const AuthUserKey = "user"										// AuthUserKey 是基本身份验证中用户凭证的 cookie 名称。
const BindKey = "_gin-gonic/gin/bindkey"						// BindKey 表示默认的绑定键。
const BodyBytesKey = "_gin-gonic/gin/bodybyteskey"				// BodyBytesKey 表示默认的正文字节密钥。
const ContextKey = "_gin-gonic/gin/contextkey"					// ContextKey 是 Context 为其返回自身的键。
const EnvGinMode = "GIN_MODE"									// EnvGinMode 指示 gin 模式的环境名称。
const Version = "v1.9.0"										// Version 是当前 gin 框架的版本。


Variables ¶
var DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)			// DebugPrintRouteFunc 表示调试日志输出格式。
var DefaultErrorWriter io.Writer = os.Stderr														// DefaultErrorWriter 是 Gin 用来调试错误的默认 io.Writer
var DefaultWriter io.Writer = os.Stdout																// DefaultWriter 是 Gin 用于调试输出和中间件输出（如 Logger() 或 Recovery()）的默认 io.Writer。
																									// 请注意，Logger 和 Recovery 都提供自定义方法来配置其输出 io.Writer。要在 Windows 中支持着色，请使用：
																									// import "github.com/mattn/go-colorable"
																									// gin.DefaultWriter = colorable.NewColorableStdout()


Functions ¶
func CreateTestContext(w http.ResponseWriter) (c *Context, r *Engine)				// CreateTestContext 返回用于测试目的的新引擎和上下文
func Dir(root string, listDirectory bool) http.FileSystem							// Dir 返回一个可以被 http.FileServer() 使用的 http.FileSystem。
																					// 它在 router.Static() 内部使用。如果 listDirectory == true，那么它的工作方式与 http.Dir() 相同，否则它返回一个阻止 http.FileServer() 列出目录文件的文件系统。
func DisableBindValidation()														// DisableBindValidation 关闭默认验证器。
func DisableConsoleColor()															// DisableConsoleColor 禁用控制台中的颜色输出。
func EnableJsonDecoderDisallowUnknownFields()										// EnableJsonDecoderDisallowUnknownFields 为 binding.EnableDecoderDisallowUnknownFields 设置 true 以调用 JSON 解码器实例上的 DisallowUnknownFields 方法。
func EnableJsonDecoderUseNumber()													// EnableJsonDecoderUseNumber 为 binding.EnableDecoderUseNumber 设置 true 以调用 JSON Decoder 实例上的 UseNumber 方法。
func ForceConsoleColor()															// ForceConsoleColor 在控制台中强制输出颜色。
func IsDebugging() bool																// 如果框架在调试模式下运行，则 IsDebugging 返回 true。使用 SetMode(gin.ReleaseMode) 禁用调试模式。
func Mode() string																	// Mode 返回当前 gin 模式。
func SetMode(value string)															// SetMode 根据输入字符串设置 gin 模式。



Types ¶
type Accounts map[string]string														// Accounts 为授权登录的用户/通行证列表定义键/值。

type Context struct {																				// Context 是 gin 最重要的部分。例如，它允许我们在中间件之间传递变量、管理流程、验证请求的 JSON 和呈现 JSON 响应。
	Request *http.Request
	Writer  ResponseWriter
	
	Params Params
	
	// Keys is a key/value pair exclusively for the context of each request.						// Keys 是专用于每个请求的上下文的键/值对。
	Keys map[string]any
	
	// Errors is a list of errors attached to all the handlers/middlewares who used this context.	// Errors 是附加到使用此上下文的所有处理程序/中间件的错误列表。
	Errors errorMsgs
	
	// Accepted defines a list of manually accepted formats for content negotiation.				// Accepted 定义了用于内容协商的手动接受格式列表。
	Accepted []string
	// contains filtered or unexported fields
}
1.func CreateTestContextOnly(w http.ResponseWriter, r *Engine) (c *Context)			// CreateTestContextOnly 返回基于引擎的新上下文以用于测试目的
2.func (c *Context) Abort()															// Abort 阻止挂起的处理程序被调用。请注意，这不会停止当前处理程序。假设您有一个验证当前请求是否已获得授权的授权中间件。如果授权失败（例如：密码不匹配），调用 Abort 以确保不调用此请求的其余处理程序。
3.func (c *Context) AbortWithError(code int, err error) *Error						// AbortWithError 在内部调用 `AbortWithStatus()` 和 `Error()`。此方法停止链，写入状态代码并将指定的错误推送到 `c.Errors`。有关更多详细信息，请参阅 Context.Error() 。
4.func (c *Context) AbortWithStatus(code int)										// AbortWithStatus 调用 Abort() 并写入具有指定状态代码的标头。例如，对请求进行身份验证的失败尝试可以使用：context.AbortWithStatus(401)。
5.func (c *Context) AbortWithStatusJSON(code int, jsonObj any)						// AbortWithStatusJSON 在内部调用“Abort()”，然后调用“JSON”。此方法停止链，写入状态代码并返回 JSON 正文。它还将 Content-Type 设置为“application/json”。
6.func (c *Context) AddParam(key, value string)										// AddParam 将参数添加到上下文并用给定值替换路径参数键以用于 e2e 测试目的示例路由：“/user/:id” AddParam(“id”, 1) 结果：“/user/1”
7.func (c *Context) AsciiJSON(code int, obj any)									// AsciiJSON 将给定的结构序列化为 JSON 到具有 unicode 到 ASCII 字符串的响应主体中。它还将 Content-Type 设置为“application/json”。
8.func (c *Context) Bind(obj any) error												// Bind 检查 Method 和 Content-Type 以自动选择绑定引擎，根据“Content-Type”标头使用不同的绑定，例如：
																					// "application/json" --> JSON binding
																					// "application/xml"  --> XML binding
																					// 如果 Content-Type == "application/json" 使用 JSON 或 XML 作为 JSON 输入，它将请求的主体解析为 JSON。它将 json 有效负载解码为指定为指针的结构。如果输入无效，它会写入 400 错误并在响应中设置 Content-Type 标头“text/plain”。
9.func (c *Context) BindHeader(obj any) error										// BindHeader 是 c.MustBindWith(obj, binding.Header) 的快捷方式。
10.func (c *Context) BindJSON(obj any) error										// BindJSON 是 c.MustBindWith(obj, binding.JSON) 的快捷方式。
11.func (c *Context) BindQuery(obj any) error										// BindQuery 是 c.MustBindWith(obj, binding.Query) 的快捷方式。
12.func (c *Context) BindTOML(obj interface{}) error								// BindTOML 是 c.MustBindWith(obj, binding.TOML) 的快捷方式。
13.func (c *Context) BindUri(obj any) error											// BindUri 使用 binding.Uri 绑定传递的结构指针。如果发生任何错误，它将使用 HTTP 400 中止请求。
14.func (c *Context) BindWith(obj any, b binding.Binding) error						// BindWith 使用指定的绑定引擎绑定传递的结构指针。查看绑定包。
15.func (c *Context) BindXML(obj any) error											// BindXML 是 c.MustBindWith(obj, binding.BindXML) 的快捷方式。
16.func (c *Context) BindYAML(obj any) error										// BindYAML 是 c.MustBindWith(obj, binding.YAML) 的快捷方式。
17.func (c *Context) ClientIP() string												// ClientIP 实现了一种尽力而为算法来返回真实的客户端 IP。
																					// 它在后台调用 c.RemoteIP() 来检查远程 IP 是否是受信任的代理。
																					// 如果是，它将尝试解析 Engine.RemoteIPHeaders 中定义的标头（默认为 [X-Forwarded-For, X-Real-Ip]）。
																					// 如果标头在语法上无效或远程 IP 不对应于受信任的代理，则返回远程 IP（来自 Request.RemoteAddr）。
18.func (c *Context) ContentType() string											// ContentType 返回请求的 Content-Type 标头。
19.func (c *Context) Cookie(name string) (string, error)							// Cookie 返回请求中提供的命名 cookie，如果未找到，则返回 ErrNoCookie。并返回未转义的命名 cookie。如果多个 cookie 与给定名称匹配，则只会返回一个 cookie。
20.func (c *Context) Copy() *Context												// Copy 返回当前上下文的副本，可以在请求范围之外安全地使用它。当必须将上下文传递给 goroutine 时，必须使用它。
21.func (c *Context) Data(code int, contentType string, data []byte)				// Data 将一些数据写入正文流并更新 HTTP 代码。
22.func (c *Context) DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string)	// DataFromReader 将指定的阅读器写入正文流并更新 HTTP 代码。
23.func (c *Context) Deadline() (deadline time.Time, ok bool)						// deadline 在c.Request没有Context的时候返回没有deadline（ok==false）。
24.func (c *Context) DefaultPostForm(key, defaultValue string) string				// DefaultPostForm 在 POST urlencoded 表单或多部分表单存在时返回指定的键，否则返回指定的 defaultValue 字符串。请参阅：PostForm() 和 GetPostForm() 了解更多信息。
25.func (c *Context) DefaultQuery(key, defaultValue string) string					// DefaultQuery 返回键控 url 查询值（如果存在），否则返回指定的 defaultValue 字符串。请参阅：Query() 和 GetQuery() 了解更多信息。
																					// GET /?name=Manu&lastname=
																					// c.DefaultQuery("name", "unknown") == "Manu"
																					// c.DefaultQuery("id", "none") == "none"
																					// c.DefaultQuery("lastname", "none") == ""
26.func (c *Context) Done() <-chan struct{}											// 当 c.Request 没有上下文时，Done 返回 nil（chan 将永远等待）。
27.func (c *Context) Err() error													// 当 c.Request 没有 Context 时，Err 返回 nil。
28.func (c *Context) Error(err error) *Error										// Error 将错误附加到当前上下文。错误被推送到错误列表。为请求解析期间发生的每个错误调用 Error 是个好主意。
																					// 中间件可用于收集所有错误并将它们一起推送到数据库、打印日志或将其附加到 HTTP 响应中。
																					// 如果 err 为 nil，错误将 panic。
29.func (c *Context) File(filepath string)											// File 以高效的方式将指定文件写入主体流。
30.func (c *Context) FileAttachment(filepath, filename string)						// FileAttachment 以高效的方式将指定的文件写入主体流 在客户端，通常会使用给定的文件名下载文件
31.func (c *Context) FileFromFS(filepath string, fs http.FileSystem)				// FileFromFS 以高效的方式将指定的文件从 http.FileSystem 写入主体流。
32.func (c *Context) FormFile(name string) (*multipart.FileHeader, error)			// FormFile 返回所提供表单键的第一个文件。
33.func (c *Context) FullPath() string												// FullPath 返回匹配的路由完整路径。对于未找到的路由，返回一个空字符串。
																					// router.GET("/user/:id", func(c *gin.Context) {
																					//  	c.FullPath() == "/user/:id" // true
																					// })
34.func (c *Context) Get(key string) (value any, exists bool)						// Get 返回给定键的值，即：(value, true)。如果该值不存在，则返回 (nil, false)
35.func (c *Context) GetBool(key string) (b bool)									// GetBool 返回与键关联的值作为布尔值。
36.func (c *Context) GetDuration(key string) (d time.Duration)						// GetDuration 返回与键关联的值作为持续时间。
37.func (c *Context) GetFloat64(key string) (f64 float64)							// GetFloat64 以 float64 形式返回与键关联的值。
38.func (c *Context) GetHeader(key string) string									// GetHeader 从请求标头返回值。
39.func (c *Context) GetInt(key string) (i int)										// GetInt 以整数形式返回与键关联的值。
40.func (c *Context) GetInt64(key string) (i64 int64)								// GetInt64 以整数形式返回与键关联的值。
41.func (c *Context) GetPostForm(key string) (string, bool)							// GetPostForm 类似于 PostForm(key)。
																					// 它从 POST urlencoded 表单或多部分表单返回指定的键，当它存在时（value，true）（即使该值是空字符串），否则它返回（“”，false）。
																					// 例如，在更新用户电子邮件的 PATCH 请求期间：
																					//     email=mail@example.com  -->  ("mail@example.com", true) := GetPostForm("email") // set email to "mail@example.com"
																					//	   email=                  -->  ("", true) := GetPostForm("email") // set email to ""
																					//                            -->  ("", false) := GetPostForm("email") // do nothing with email
42.func (c *Context) GetPostFormArray(key string) (values []string, ok bool)		// GetPostFormArray 为给定的表单键返回一段字符串，加上一个布尔值是否至少有一个值存在于给定的键。
43.func (c *Context) GetPostFormMap(key string) (map[string]string, bool)			// GetPostFormMap 为给定的表单键返回一个映射，加上一个布尔值是否至少存在一个给定键的值。
44.func (c *Context) GetQuery(key string) (string, bool)							// GetQuery 类似于 Query()，如果它存在 `(value, true)`（即使该值为空字符串），它返回键控 url 查询值，否则返回 `("", false)`。
																					// 它是 `c.Request.URL.Query().Get(key)` 的快捷方式
																					// GET /?name=Manu&lastname=
																					// ("Manu", true) == c.GetQuery("name")
																					// ("", false) == c.GetQuery("id")
																					// ("", true) == c.GetQuery("lastname")
45.func (c *Context) GetQueryArray(key string) (values []string, ok bool)			// GetQueryArray 返回给定查询键的字符串片段，以及一个布尔值，以确定给定键是否存在至少一个值。
46.func (c *Context) GetQueryMap(key string) (map[string]string, bool)				// GetQueryMap 返回给定查询键的映射，加上一个布尔值，以确定给定键是否存在至少一个值。
47.func (c *Context) GetRawData() ([]byte, error)									// GetRawData 返回流数据。
48.func (c *Context) GetString(key string) (s string)								// GetString 以字符串形式返回与键关联的值。
49.func (c *Context) GetStringMap(key string) (sm map[string]any)					// GetStringMap 返回与键关联的值作为接口映射。
50.func (c *Context) GetStringMapString(key string) (sms map[string]string)			// GetStringMapString 返回与键关联的值作为字符串映射。
51.func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string)	// GetStringMapStringSlice 将与键关联的值作为到字符串切片的映射返回。
52.func (c *Context) GetStringSlice(key string) (ss []string)						// GetStringSlice 返回与键关联的值作为字符串切片。
53.func (c *Context) GetTime(key string) (t time.Time)								// GetTime 返回与键关联的值作为时间。
54.func (c *Context) GetUint(key string) (ui uint)									// GetUint 将与键关联的值作为无符号整数返回。
55.func (c *Context) GetUint64(key string) (ui64 uint64)							// GetUint64 以无符号整数形式返回与键关联的值。
56.func (c *Context) HTML(code int, name string, obj any)							// HTML 呈现由其文件名指定的 HTTP 模板。它还会更新 HTTP 代码并将 Content-Type 设置为“text/html”。请参阅 http://golang.org/doc/articles/wiki/
57.func (c *Context) Handler() HandlerFunc											// Handler 返回主处理程序。
58.func (c *Context) HandlerName() string											// HandlerName 返回主处理程序的名称。例如，如果处理程序是“handleGetUsers()”，此函数将返回“main.handleGetUsers”。
59.func (c *Context) HandlerNames() []string										// HandlerNames 按照 HandlerName() 的语义按降序返回此上下文的所有已注册处理程序的列表
60.func (c *Context) Header(key, value string)										// Header 是 c.Writer.Header().Set(key, value) 的智能快捷方式。它在响应中写入一个标头。如果 value == ""，此方法删除标头 `c.Writer.Header().Del(key)`
61.func (c *Context) IndentedJSON(code int, obj any)								// IndentedJSON 将给定的结构序列化为漂亮的 JSON（缩进 + 结束行）到响应主体中。
																					// 它还将 Content-Type 设置为“application/json”。
																					// 警告：我们建议仅将此用于开发目的，因为打印漂亮的 JSON 会消耗更多的 CPU 和带宽。 请改用 Context.JSON() 。
62.func (c *Context) IsAborted() bool												// 如果当前上下文已中止，则 IsAborted 返回 true。
63.func (c *Context) IsWebsocket() bool												// 如果请求标头指示客户端正在启动 websocket 握手，则 IsWebsocket 返回 true。
64.func (c *Context) JSON(code int, obj any)										// JSON 将给定的结构序列化为 JSON 到响应主体中。它还将 Content-Type 设置为“application/json”。
65.func (c *Context) JSONP(code int, obj any)										// JSONP 将给定的结构作为 JSON 序列化到响应主体中。它将填充添加到响应主体，以从与客户端位于不同域中的服务器请求数据。它还将 Content-Type 设置为“application/javascript”。
66.func (c *Context) MultipartForm() (*multipart.Form, error)						// MultipartForm 是解析的多部分表单，包括文件上传。
67.func (c *Context) MustBindWith(obj any, b binding.Binding) error					// MustBindWith 使用指定的绑定引擎绑定传递的结构指针。如果发生任何错误，它将使用 HTTP 400 中止请求。查看绑定包。
68.func (c *Context) MustGet(key string) any										// MustGet 返回给定键的值（如果存在），否则它会崩溃。
69.func (c *Context) Negotiate(code int, config Negotiate)							// 根据可接受的Accept格式, Negotiate调用不同的Render。
70.func (c *Context) NegotiateFormat(offered ...string) string						// NegotiateFormat 返回可接受的 Accept 格式。
71.func (c *Context) Next()															// Next 应该只在中间件内部使用。它在调用处理程序内执行链中的挂起处理程序。请参阅 GitHub 中的示例。
72.func (c *Context) Param(key string) string										// Param 返回 URL 参数的值。它是 c.Params.ByName(key) 的快捷方式
																					// router.GET("/user/:id", func(c *gin.Context) {
																					//    // a GET request to /user/john
																					//    id := c.Param("id") // id == "/john"
																					//    // a GET request to /user/john/
																					//    id := c.Param("id") // id == "/john/"
																					// })
73.func (c *Context) PostForm(key string) (value string)							// PostForm 在 POST urlencoded 表单或多部分表单存在时返回指定的键，否则返回空字符串 `("")`。
74.func (c *Context) PostFormArray(key string) (values []string)					// PostFormArray 返回给定表单键的字符串片段。切片的长度取决于具有给定键的参数的数量。
75.func (c *Context) PostFormMap(key string) (dicts map[string]string)				// PostFormMap 返回给定表单键的映射。
76.func (c *Context) ProtoBuf(code int, obj any)									// ProtoBuf 将给定的结构序列化为 ProtoBuf 到响应主体中。
77.func (c *Context) PureJSON(code int, obj any)									// PureJSON 将给定的结构序列化为 JSON 到响应主体中。与 JSON 不同，PureJSON 不会用它们的 unicode 实体替换特殊的 html 字符。
78.func (c *Context) Query(key string) (value string)								// Query 返回键控 url 查询值（如果存在），否则返回空字符串 `("") 。它是 `c.Request.URL.Query().Get(key)` 的快捷方式
																					// GET /path?id=1234&name=Manu&value=
																					// 	 c.Query("id") == "1234"
																					//	 c.Query("name") == "Manu"
																					//	 c.Query("value") == ""
																					//	 c.Query("wtf") == ""
79.func (c *Context) QueryArray(key string) (values []string)						// QueryArray 返回给定查询键的字符串片段。切片的长度取决于具有给定键的参数的数量。
80.func (c *Context) QueryMap(key string) (dicts map[string]string)					// QueryMap 返回给定查询键的映射。
81.func (c *Context) Redirect(code int, location string)							// Redirect 返回到特定位置的 HTTP 重定向。
82.func (c *Context) RemoteIP() string												// RemoteIP 从 Request.RemoteAddr 解析 IP，规范化并返回 IP（不带端口）。
83.func (c *Context) Render(code int, r render.Render)								// Render 写入响应头并调用 render.Render 来渲染数据。
84.func (c *Context) SSEvent(name string, message any)								// SSEvent 将服务器发送的事件写入正文流。
85.func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error	// SaveUploadedFile 将表单文件上传到特定的dst。
86.func (c *Context) SecureJSON(code int, obj any)									// SecureJSON 将给定的结构作为安全 JSON 序列化到响应正文中。如果给定的结构是数组值，则默认将“while(1)”添加到响应主体。它还将 Content-Type 设置为“application/json”。
87.func (c *Context) Set(key string, value any)										// Set 用于存储专用于此上下文的新键/值对。如果 c.Keys 以前没有使用过，它也会延迟初始化。
88.func (c *Context) SetAccepted(formats ...string)									// SetAccepted 设置接受头数据。
89.func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)		// SetCookie 将 Set-Cookie 标头添加到 ResponseWriter 的标头中。提供的 cookie 必须具有有效名称。无效的 cookie 可能会被静默删除。
90.func (c *Context) SetSameSite(samesite http.SameSite)							// SetSameSite 与 cookie
91.func (c *Context) ShouldBind(obj any) error										// ShouldBind 检查 Method 和 Content-Type 以自动选择绑定引擎，根据“Content-Type”标头使用不同的绑定，例如：
																					// "application/json" --> JSON binding
																					// "application/xml"  --> XML binding
																					// 如果 Content-Type == "application/json" 使用 JSON 或 XML 作为 JSON 输入，它将请求的主体解析为 JSON。
																					// 它将 json 有效负载解码为指定为指针的结构。与 c.Bind() 类似，但此方法不会将响应状态代码设置为 400 或在输入无效时中止。
92.func (c *Context) ShouldBindBodyWith(obj any, bb binding.BindingBody) (err error)// ShouldBindBodyWith 与 ShouldBindWith 类似，但它将请求体存储到上下文中，并在再次调用时重用。
																					// 注意：此方法在绑定之前读取正文。所以如果你只需要调用一次，你应该使用 ShouldBindWith 以获得更好的性能。
93.func (c *Context) ShouldBindHeader(obj any) error								// ShouldBindHeader 是 c.ShouldBindWith(obj, binding.Header) 的快捷方式。
94.func (c *Context) ShouldBindJSON(obj any) error									// ShouldBindJSON 是 c.ShouldBindWith(obj, binding.JSON) 的快捷方式。
95.func (c *Context) ShouldBindQuery(obj any) error									// ShouldBindQuery 是 c.ShouldBindWith(obj, binding.Query) 的快捷方式。
96.func (c *Context) ShouldBindTOML(obj interface{}) error							// ShouldBindTOML 是 c.ShouldBindWith(obj, binding.TOML) 的快捷方式。
97.func (c *Context) ShouldBindUri(obj any) error									// ShouldBindUri 使用指定的绑定引擎绑定传递的结构指针。
98.func (c *Context) ShouldBindWith(obj any, b binding.Binding) error				// ShouldBindWith 使用指定的绑定引擎绑定传递的结构指针。查看绑定包。
99.func (c *Context) ShouldBindXML(obj any) error									// ShouldBindXML 是 c.ShouldBindWith(obj, binding.XML) 的快捷方式。
100.func (c *Context) ShouldBindYAML(obj any) error									// ShouldBindYAML 是 c.ShouldBindWith(obj, binding.YAML) 的快捷方式。
101.func (c *Context) Status(code int)												// Status 设置 HTTP 响应代码。
102.func (c *Context) Stream(step func(w io.Writer) bool) bool						// Stream 发送流式响应并返回一个布尔值，表示“客户端是否在流中间断开连接”
103.func (c *Context) String(code int, format string, values ...any)				// String 将给定的字符串写入响应主体。
104.func (c *Context) TOML(code int, obj interface{})								// TOML 将给定的结构作为 TOML 序列化到响应主体中。
105.func (c *Context) Value(key any) any											// Value 返回与 key 上下文关联的值，如果没有值与 key 关联，则返回 nil。使用相同的键连续调用 Value 会返回相同的结果。
106.func (c *Context) XML(code int, obj any)										// XML 将给定的结构作为 XML 序列化到响应正文中。它还将 Content-Type 设置为“application/xml”。
107.func (c *Context) YAML(code int, obj any)										// YAML 将给定的结构序列化为 YAML 到响应正文中。





type Engine struct {	// Engine 是框架的实例，它包含混合器、中间件和配置设置。使用 New() 或 Default() 创建 Engine 实例
	RouterGroup
	
	// RedirectTrailingSlash enables automatic redirection if the current route can't be matched but a		// RedirectTrailingSlash 如果当前路由无法匹配但启用自动重定向
	// handler for the path with (without) the trailing slash exists.										// 存在（没有）尾部斜杠的路径处理程序。
	// For example if /foo/ is requested but a route only exists for /foo, the								// 例如，如果请求 /foo/ 但仅存在 /foo 的路由，则
	// client is redirected to /foo with http status code 301 for GET requests								// 客户端被重定向到 /foo，GET 请求的 http 状态代码为 301
	// and 307 for all other request methods.																// 和 307 用于所有其他请求方法。
	RedirectTrailingSlash bool
	
	// RedirectFixedPath if enabled, the router tries to fix the current request path, if no				// RedirectFixedPath 如果启用，路由器尝试修复当前请求路径，如果没有
	// handle is registered for it.																			// 句柄已为其注册。
	// First superfluous path elements like ../ or // are removed.											// 第一个多余的路径元素，如 ../ 或 // 被删除。
	// Afterwards the router does a case-insensitive lookup of the cleaned path.							// 之后路由器对清理后的路径进行不区分大小写的查找。
	// If a handle can be found for this route, the router makes a redirection								// 如果可以找到该路由的句柄，则路由器进行重定向
	// to the corrected path with status code 301 for GET requests and 307 for								// 更正后的路径，GET 请求的状态码为 301，GET 请求的状态码为 307
	// all other request methods.																			// 所有其他请求方法。
	// For example /FOO and /..//Foo could be redirected to /foo.											// 例如 /FOO 和 /..//Foo 可以重定向到 /foo。
	// RedirectTrailingSlash is independent of this option.													// RedirectTrailingSlash 独立于此选项。
	RedirectFixedPath bool
	
	// HandleMethodNotAllowed if enabled, the router checks if another method is allowed for the			// HandleMethodNotAllowed 如果启用，路由器检查是否允许另一个方法
	// current route, if the current request can not be routed.												// 当前路由，如果当前请求无法路由。
	// If this is the case, the request is answered with 'Method Not Allowed'								// 如果是这种情况，请求将以 'Method Not Allowed' 回答
	// and HTTP status code 405.																			// 和 HTTP 状态代码 405。
	// If no other Method is allowed, the request is delegated to the NotFound								// 如果不允许其他方法，则将请求委托给 NotFound
	// handler.																								// 处理程序。
	HandleMethodNotAllowed bool
	
	// ForwardedByClientIP if enabled, client IP will be parsed from the request's headers that				// ForwardedByClientIP 如果启用，客户端 IP 将从请求的标头中解析
	// match those stored at `(*gin.Engine).RemoteIPHeaders`. If no IP was									// 匹配存储在 `(*gin.Engine).RemoteIPHeaders` 中的那些。如果没有IP
	// fetched, it falls back to the IP obtained from														// 已获取，它会回退到从中获取的 IP
	// `(*gin.Context).Request.RemoteAddr`.																	// `(*gin.Context).Request.RemoteAddr`。
	ForwardedByClientIP bool
	
	// AppEngine was deprecated.																			// AppEngine 已弃用。
	// Deprecated: USE `TrustedPlatform` WITH VALUE `gin.PlatformGoogleAppEngine` INSTEAD					// 已弃用：使用带有值 `gin.PlatformGoogleAppEngine` 的 `TrustedPlatform` 代替
	// #726 #755 If enabled, it will trust some headers starting with										// #726 #755 如果启用，它将信任一些以
	// 'X-AppEngine...' for better integration with that PaaS.												// 'X-AppEngine...' 以便更好地与该 PaaS 集成。
	AppEngine bool
	
	// UseRawPath if enabled, the url.RawPath will be used to find parameters.								// UseRawPath 如果启用，url.RawPath 将用于查找参数。
	UseRawPath bool
	
	// UnescapePathValues if true, the path value will be unescaped.										// UnescapePathValues 如果为真，路径值将被取消转义。
	// If UseRawPath is false (by default), the UnescapePathValues effectively is true,						// 如果 UseRawPath 为 false（默认情况下），则 UnescapePathValues 实际上为 true，
	// as url.Path gonna be used, which is already unescaped.												// 作为 url.Path 将被使用，它已经未转义。
	UnescapePathValues bool
	
	// RemoveExtraSlash a parameter can be parsed from the URL even with extra slashes.						// RemoveExtraSlash 即使有额外的斜杠，也可以从 URL 中解析参数。
	// See the PR #1817 and issue #1644																		// 请参阅 PR #1817 和问题 #1644
	RemoveExtraSlash bool
	
	// RemoteIPHeaders list of headers used to obtain the client IP when									// RemoteIPHeaders headers列表，用于获取客户端IP时
	// `(*gin.Engine).ForwardedByClientIP` is `true` and													// `(*gin.Engine).ForwardedByClientIP` 为 `true` 并且
	// `(*gin.Context).Request.RemoteAddr` is matched by at least one of the								// `(*gin.Context).Request.RemoteAddr` 与至少一个匹配
	// network origins of list defined by `(*gin.Engine).SetTrustedProxies()`.								// 由 `(*gin.Engine).SetTrustedProxies()` 定义的列表的网络来源。
	RemoteIPHeaders []string
	
	// TrustedPlatform if set to a constant of value gin.Platform*, trusts the headers set by				// TrustedPlatform 如果设置为值 gin.Platform* 的常量，则信任由
	// that platform, for example to determine the client IP												// 该平台，例如确定客户端IP
	TrustedPlatform string
	
	// MaxMultipartMemory value of 'maxMemory' param that is given to http.Request's ParseMultipartForm		// 给 http.Request 的 ParseMultipartForm 的 'maxMemory' 参数的 MaxMultipartMemory 值
	// method call.																							// 方法调用。
	MaxMultipartMemory int64
	
	// UseH2C enable h2c support.																			// UseH2C 启用 h2c 支持。
	UseH2C bool
	
	// ContextWithFallback 在 Context.Request.Context() 不为 nil 时启用回退 Context.Deadline()、Context.Done()、Context.Err() 和 Context.Value()。
	// ContextWithFallback enable fallback Context.Deadline(), Context.Done(), Context.Err() and Context.Value() when Context.Request.Context() is not nil.
	ContextWithFallback bool
	
	HTMLRender render.HTMLRender
	FuncMap    template.FuncMap
	// contains filtered or unexported fields
}
1.func Default() *Engine											// Default 返回一个 Engine 实例，其中已经附加了 Logger 和 Recovery 中间件。
2.func New() *Engine												// New 返回一个新的空白 Engine 实例，没有附加任何中间件。默认情况下，配置为：
																	// - RedirectTrailingSlash: true
																	// - RedirectFixedPath: false
																	// - HandleMethodNotAllowed: false
																	// - ForwardedByClientIP: true
																	// - UseRawPath: false
																	// - UnescapePathValues: true
3.func (engine *Engine) Delims(left, right string) *Engine			// Delims 设置模板左右 delims 并返回一个 Engine 实例。
4.func (engine *Engine) HandleContext(c *Context)					// HandleContext 重新进入已被重写的上下文。这可以通过将 c.Request.URL.Path 设置为新目标来完成。免责声明：您可以循环自己来处理这个问题，明智地使用。
5.func (engine *Engine) Handler() http.Handler
6.func (engine *Engine) LoadHTMLFiles(files ...string)				// LoadHTMLFiles 加载一片 HTML 文件并将结果与 HTML 渲染器相关联。
7.func (engine *Engine) LoadHTMLGlob(pattern string)				// LoadHTMLGlob 加载由 glob 模式标识的 HTML 文件，并将结果与 HTML 渲染器相关联。
8.func (engine *Engine) NoMethod(handlers ...HandlerFunc)			// NoMethod 设置当 Engine.HandleMethodNotAllowed = true 时调用的处理程序。
9.func (engine *Engine) NoRoute(handlers ...HandlerFunc)			// NoRoute 为 NoRoute 添加处理程序。它默认返回 404 代码。
10.func (engine *Engine) Routes() (routes RoutesInfo)				// Routes 返回注册路由的一部分，包括一些有用的信息，例如：http 方法、路径和处理程序名称。
11.func (engine *Engine) Run(addr ...string) (err error)			// Run 将路由器附加到 http.Server 并开始侦听和处理 HTTP 请求。它是 http.ListenAndServe(addr, router) 的快捷方式
																	// 注意：除非发生错误，否则此方法将无限期地阻止调用 goroutine。
12.func (engine *Engine) RunFd(fd int) (err error)					// RunFd 将路由器附加到 http.Server 并通过指定的文件描述符开始侦听和服务 HTTP 请求。
																	// 注意：除非发生错误，否则此方法将无限期地阻塞调用 goroutine。
13.func (engine *Engine) RunListener(listener net.Listener) (err error)	// RunListener 将路由器附加到 http.Server 并通过指定的 net.Listener 开始侦听和服务 HTTP 请求
14.func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error)		// RunTLS 将路由器附加到 http.Server 并开始侦听和服务 HTTPS（安全）请求。它是 http.ListenAndServeTLS(addr, certFile, keyFile, router) 的快捷方式
																				// 注意：除非发生错误，否则此方法将无限期地阻止调用 goroutine。
15.func (engine *Engine) RunUnix(file string) (err error)			// RunUnix 将路由器附加到 http.Server 并开始通过指定的 unix 套接字（即文件）侦听和服务 HTTP 请求。
																	// 注意：除非发生错误，否则此方法将无限期地阻塞调用 goroutine。
16.func (engine *Engine) SecureJsonPrefix(prefix string) *Engine	// SecureJsonPrefix 设置在 Context.SecureJSON 中使用的 secureJSONPrefix。
17.func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)	// ServeHTTP 符合 http.Handler 接口。
18.func (engine *Engine) SetFuncMap(funcMap template.FuncMap)		// SetFuncMap 设置用于模板的 FuncMap。
19.func (engine *Engine) SetHTMLTemplate(templ *template.Template)	// SetHTMLTemplate 将模板与 HTML 渲染器相关联。
20.func (engine *Engine) SetTrustedProxies(trustedProxies []string) error		// SetTrustedProxies 设置网络源列表（IPv4 地址、IPv4 CIDR、IPv6 地址或 IPv6 CIDR），当“(*gin.Engine).ForwardedByClientIP”为“true”时，从中信任包含备用客户端 IP 的请求标头。
																				// `TrustedProxies` 功能默认启用，默认情况下它也信任所有代理。如果要禁用此功能，请使用 Engine.SetTrustedProxies(nil)，然后 Context.ClientIP() 将直接返回远程地址。
21.func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes		// 使用将全局中间件附加到路由器。即通过 Use() 附加的中间件将包含在每个请求的处理程序链中。甚至 404、405、静态文件...例如，这是记录器或错误管理中间件的正确位置。



type Error struct {		// Error 表示错误的规范。
	Err  error
	Type ErrorType
	Meta any
}
1.func (msg Error) Error() string									// Error 实现错误接口。
2.func (msg *Error) IsType(flags ErrorType) bool					// IsType 判断一个错误。
3.func (msg *Error) JSON() any										// JSON 创建格式正确的 JSON
4.func (msg *Error) MarshalJSON() ([]byte, error)					// MarshalJSON 实现了 json.Marshaller 接口。
5.func (msg *Error) SetMeta(data any) *Error						// SetMeta 设置错误的元数据。
6.func (msg *Error) SetType(flags ErrorType) *Error					// SetType 设置错误的类型。
7.func (msg *Error) Unwrap() error									// Unwrap 返回包装的错误，以允许与 errors.Is()、errors.As() 和 errors.Unwrap() 的互操作性



type ErrorType uint64	// ErrorType 是 gin 规范中定义的无符号 64 位错误代码。
const (
	// ErrorTypeBind is used when Context.Bind() fails.					// ErrorTypeBind 在 Context.Bind() 失败时使用。
	ErrorTypeBind ErrorType = 1 << 63
	// ErrorTypeRender is used when Context.Render() fails.				// ErrorTypeRender 在 Context.Render() 失败时使用。
	ErrorTypeRender ErrorType = 1 << 62
	// ErrorTypePrivate indicates a private error.						// ErrorTypePrivate 表示私有错误。
	ErrorTypePrivate ErrorType = 1 << 0
	// ErrorTypePublic indicates a public error.						// ErrorTypePublic 表示公共错误。
	ErrorTypePublic ErrorType = 1 << 1
	// ErrorTypeAny indicates any other error.							// ErrorTypeAny 指示任何其他错误。
	ErrorTypeAny ErrorType = 1<<64 - 1
	// ErrorTypeNu indicates any other error.							// ErrorTypeNu 表示任何其他错误。
	ErrorTypeNu = 2
)


type H map[string]any													// H 是 map[string]interface{} 的快捷方式
1.func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error	// MarshalXML 允许类型 H 与 xml.Marshal 一起使用。


type HandlerFunc func(*Context)											// HandlerFunc 将 gin 中间件使用的处理程序定义为返回值。
1.func BasicAuth(accounts Accounts) HandlerFunc							// BasicAuth 返回一个基本 HTTP 授权中间件。它以 map[string]string 作为参数，其中键是用户名，值是密码。
2.func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc	// BasicAuthForRealm 返回一个基本的 HTTP 授权中间件。它以一个 map[string]string 作为参数，其中键是用户名，值是密码，以及 Realm 的名称。如果领域为空，则默认使用“需要授权”。 （参见 http://tools.ietf.org/html/rfc2617#section-1.2）
3.func Bind(val any) HandlerFunc										// Bind 是给定接口对象的辅助函数，并返回一个 Gin 中间件。
4.func CustomRecovery(handle RecoveryFunc) HandlerFunc					// CustomRecovery 返回一个从任何恐慌中恢复的中间件，并调用提供的句柄函数来处理它。
5.func CustomRecoveryWithWriter(out io.Writer, handle RecoveryFunc) HandlerFunc		// CustomRecoveryWithWriter 为给定的 writer 返回一个中间件，它从任何恐慌中恢复并调用提供的 handle func 来处理它。
6.func ErrorLogger() HandlerFunc										// ErrorLogger 为任何错误类型返回一个 HandlerFunc。
7.func ErrorLoggerT(typ ErrorType) HandlerFunc							// ErrorLoggerT 返回给定错误类型的 HandlerFunc。
8.func Logger() HandlerFunc												// Logger 实例是一个 Logger 中间件，它将日志写入 gin.DefaultWriter。默认情况下，gin.DefaultWriter = os.Stdout。
9.func LoggerWithConfig(conf LoggerConfig) HandlerFunc					// LoggerWithConfig 实例一个带有配置的 Logger 中间件。
10.func LoggerWithFormatter(f LogFormatter) HandlerFunc					// LoggerWithFormatter 实例一个具有指定日志格式功能的 Logger 中间件。
11.func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc// LoggerWithWriter 实例具有指定写入器缓冲区的 Logger 中间件。示例：os.Stdout，以写入模式打开的文件，套接字...
12.func Recovery() HandlerFunc											// Recovery 返回一个中间件，该中间件可以从任何 panic 中恢复并在出现 panic 时写入 500。
13.func RecoveryWithWriter(out io.Writer, recovery ...RecoveryFunc) HandlerFunc		// RecoveryWithWriter 为给定的 writer 返回一个中间件，该中间件从任何 panic 中恢复并写入 500（如果有）。
14.func WrapF(f http.HandlerFunc) HandlerFunc							// WrapF 是一个用于包装 http.HandlerFunc 并返回一个 Gin 中间件的辅助函数。
15.func WrapH(h http.Handler) HandlerFunc								// WrapH 是一个用于包装 http.Handler 并返回一个 Gin 中间件的辅助函数。



type HandlersChain []HandlerFunc										// HandlersChain 定义了一个 HandlerFunc 切片。
1.func (c HandlersChain) Last() HandlerFunc								// Last 返回链中的最后一个处理程序。即最后一个处理程序是主要处理程序。



type IRouter interface {												// IRouter 定义了所有路由器句柄接口，包括单个和组路由器。
	IRoutes
	Group(string, ...HandlerFunc) *RouterGroup
}



type IRoutes interface {												// IRoutes 定义了所有路由器句柄接口。
	Use(...HandlerFunc) IRoutes
	
	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes
	Match([]string, string, ...HandlerFunc) IRoutes
	
	StaticFile(string, string) IRoutes
	StaticFileFS(string, string, http.FileSystem) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}



type LogFormatter func(params LogFormatterParams) string				// LogFormatter 给出传递给 LoggerWithFormatter 的格式化程序函数的签名



type LogFormatterParams struct {										// LogFormatterParams 是在记录时间到来时将传递给任何格式化程序的结构
	Request *http.Request
	
	// TimeStamp shows the time after the server returns a response.				// TimeStamp 显示服务器返回响应后的时间。
	TimeStamp time.Time
	// StatusCode is HTTP response code.											// StatusCode 是 HTTP 响应代码。
	StatusCode int
	// Latency is how much time the server cost to process a certain request.		// Latency 是服务器处理某个请求所花费的时间。
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.									// ClientIP 等于 Context 的 ClientIP 方法。
	ClientIP string
	// Method is the HTTP method given to the request.								// Method 是给请求的 HTTP 方法。
	Method string
	// Path is a path the client requests.											// Path 是客户端请求的路径。
	Path string
	// ErrorMessage is set if error has occurred in processing the request.			// 如果在处理请求时发生错误，则设置 ErrorMessage。
	ErrorMessage string
	
	// BodySize is the size of the Response Body									// BodySize 是 Response Body 的大小
	BodySize int
	// Keys are the keys set on the request's context.								// Keys 是在请求的上下文中设置的键。
	Keys map[string]any
	// contains filtered or unexported fields
}
1.func (p *LogFormatterParams) IsOutputColor() bool						// IsOutputColor 表示是否可以将颜色输出到日志中。
2.func (p *LogFormatterParams) MethodColor() string						// MethodColor 是用于将 http 方法正确记录到终端的 ANSI 颜色。
3.func (p *LogFormatterParams) ResetColor() string						// ResetColor 重置所有转义属性。
4.func (p *LogFormatterParams) StatusCodeColor() string					// StatusCodeColor 是用于将 http 状态代码正确记录到终端的 ANSI 颜色。




type LoggerConfig struct {		// LoggerConfig 定义 Logger 中间件的配置。
	// Optional. Default value is gin.defaultLogFormatter				// 选修的。默认值为 gin.defaultLogFormatter
	Formatter LogFormatter
	
	// Output is a writer where logs are written.						// Output 是写入日志的 writer。
	// Optional. Default value is gin.DefaultWriter.					// 选修的。默认值为 gin.DefaultWriter。
	Output io.Writer
	
	// SkipPaths is an url path array which logs are not written.		// SkipPaths 是不写入日志的url路径数组。
	// Optional.														// 选修的。
	SkipPaths []string
}



type Negotiate struct {		// Negotiate 包含所有谈判数据。
	Offered  []string
	HTMLName string
	HTMLData any
	JSONData any
	XMLData  any
	YAMLData any
	Data     any
	TOMLData any
}



type Param struct {			// Param 是单个 URL 参数，由键和值组成。
	Key   string
	Value string
}


type Params []Param			// Params 是一个参数切片，由路由器返回。切片是有序的，第一个 URL 参数也是第一个切片值。因此通过索引读取值是安全的。
1.func (ps Params) ByName(name string) (va string)		// ByName 返回与给定名称匹配的第一个参数的值。如果没有找到匹配的参数，则返回一个空字符串。
2.func (ps Params) Get(name string) (string, bool)		// Get 返回与给定名称匹配的第一个 Param 的值和布尔值 true。如果找不到匹配的参数，则返回一个空字符串和一个布尔值 false 。


type RecoveryFunc func(c *Context, err any)		// RecoveryFunc 定义可传递给 CustomRecovery 的函数。


type ResponseWriter interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier
	
	// Status returns the HTTP response status code of the current request.				// Status 返回当前请求的 HTTP 响应状态码。
	Status() int
	
	// Size returns the number of bytes already written into the response http body.	// Size 返回已经写入响应 http 主体的字节数。
	// See Written()																	// 参见 Written()
	Size() int
	
	// WriteString writes the string into the response body.							// WriteString 将字符串写入响应主体。
	WriteString(string) (int, error)
	
	// Written returns true if the response body was already written.					// 如果响应正文已经写入，则 Written 返回 true。
	Written() bool
	
	// WriteHeaderNow forces to write the http header (status code + headers).			// WriteHeaderNow 强制写入 http 标头（状态代码 + 标头）。
	WriteHeaderNow()
	
	// Pusher get the http.Pusher for server push										// Pusher 获取服务器推送的http.Pusher
	Pusher() http.Pusher
}



type RouteInfo struct {				// RouteInfo 表示请求路由的规范，其中包含方法和路径及其处理程序。
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}



type RouterGroup struct {							// RouterGroup 在内部用于配置路由器，一个 RouterGroup 与一个前缀和一组处理程序（中间件）相关联。
	Handlers HandlersChain
	// contains filtered or unexported fields
}
1.func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes		// Any 注册一个匹配所有 HTTP 方法的路由。 GET、POST、PUT、PATCH、HEAD、OPTIONS、DELETE、CONNECT、TRACE。
2.func (group *RouterGroup) BasePath() string												// BasePath 返回路由器组的基本路径。例如，如果 v := router.Group("/rest/n/v1/api")，则 v.BasePath() 为“/rest/n/v1/api”。
3.func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes	// DELETE 是 router.Handle("DELETE", path, handlers) 的快捷方式。
4.func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes		// GET 是 router.Handle("GET", path, handlers) 的快捷方式。
5.func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup// Group 创建一个新的路由器组。您应该添加所有具有公共中间件或相同路径前缀的路由。例如，可以对所有使用通用中间件进行授权的路由进行分组。
6.func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes		// HEAD 是 router.Handle("HEAD", path, handlers) 的快捷方式。
7.func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes	// Handle 使用给定的路径和方法注册一个新的请求句柄和中间件。最后一个处理程序应该是真正的处理程序，其他的应该是可以并且应该在不同路由之间共享的中间件。请参阅 GitHub 中的示例代码。
																										// 对于 GET、POST、PUT、PATCH 和 DELETE 请求，可以使用各自的快捷方式功能。
																										// 此功能旨在用于批量加载，并允许使用不常用的、非标准化或自定义方法（例如，用于与代理的内部通信）。
8.func (group *RouterGroup) Match(methods []string, relativePath string, handlers ...HandlerFunc) IRoutes// Match 注册一个与您声明的指定方法相匹配的路由。
9.func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes	// OPTIONS 是 router.Handle("OPTIONS", path, handlers) 的快捷方式。
10.func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes	// PATCH 是 router.Handle("PATCH", path, handlers) 的快捷方式。
11.func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes		// POST 是 router.Handle("POST", path, handlers) 的快捷方式。
12.func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes		// PUT 是 router.Handle("PUT", path, handlers) 的快捷方式。
13.func (group *RouterGroup) Static(relativePath, root string) IRoutes						// Static 服务来自给定文件系统根目录的文件。在内部使用 http.FileServer，因此使用 http.NotFound 代替路由器的 NotFound 处理程序。要使用操作系统的文件系统实现，请使用：
																							// router.Static("/static", "/var/www")
14.func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes		// StaticFS 的工作方式与 `Static()` 类似，但可以改用自定义的 `http.FileSystem`。 Gin 默认使用：gin.Dir()
15.func (group *RouterGroup) StaticFile(relativePath, filepath string) IRoutes				// StaticFile 注册一个路由，以便为本地文件系统的单个文件提供服务。 router.StaticFile("favicon.ico", "./resources/favicon.ico")
16.func (group *RouterGroup) StaticFileFS(relativePath, filepath string, fs http.FileSystem) IRoutes		// StaticFileFS 的工作方式与 `StaticFile` 类似，但可以使用自定义的 `http.FileSystem` 代替.. router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false}) Gin 默认使用：gin.Dir()
17.func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes							// Use 将中间件添加到组中，请参阅 GitHub 中的示例代码。



type RoutesInfo []RouteInfo		// RoutesInfo 定义了一个 RouteInfo 切片。
