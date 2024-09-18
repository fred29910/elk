package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("coderx", func() {
	Title("CoderX Service")
	Description("Service for code generation and management")
	Server("coderx", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})
var ParserResult = Type("ParserResult", func() {
	Field(1, "parsedCode", String, "解析后的代码")
	Field(2, "errors", ArrayOf(String), "解析错误信息")
	Required("parsedCode")
})
var QueryParams = Type("QueryParams", func() {
	Field(1, "fields", ArrayOf(String), "选择的字段")
	Field(2, "filters", MapOf(String, Any), "过滤条件")
	Field(3, "page", Int, "页码")
	Field(4, "pageSize", Int, "每页数量")
})

var _ = Service("coderx", func() {
	Description("The coderx service provides code generation and management functionalities.")

	Method("parseCode", func() {
		Payload(func() {
			Field(1, "code", String, "待解析的代码")
			Field(2, "Authorization", String, "Bearer {token}", func() { // 添加 Authorization 到 Payload
				Description("JWT 令牌")
			})
			Required("code")
		})
		Result(ParserResult) // 使用 ParserResult 作为结果类型
		HTTP(func() {
			POST("/parse")
			Header("Authorization", String, "Bearer {token}", func() { // 添加 Token 到 Header
				Description("JWT 令牌")
			})
		})
		GRPC(func() {
		})
	})

	Method("generateCode", func() {
		Payload(func() {
			Field(1, "language", String, "Programming language")
			Field(2, "specification", String, "Code specification")
			Required("language", "specification")
		})
		Result(ParserResult)
		HTTP(func() {
			POST("/generate")
		})
		GRPC(func() {
		})
	})

	Method("analyzeCode", func() {
		Payload(func() {
			Field(1, "code", String, "Code to analyze")
			Required("code")
		})
		Result(String)
		HTTP(func() {
			POST("/analyze")
		})
		GRPC(func() {
		})
	})

	Method("optimizeCode", func() {
		Payload(func() {
			Field(1, "code", String, "Code to optimize")
			Field(2, "target", String, "Optimization target (e.g., 'performance', 'memory')")
			Required("code", "target")
		})
		Result(String)
		HTTP(func() {
			POST("/optimize")
		})
		GRPC(func() {
		})
	})

	Method("listUsers", func() {
		Description("列出用户")
		Payload(func() {
			// Attribute("queryParams", QueryParams)
			Field(1, "fields", ArrayOf(String), "选择的字段")
			Field(2, "filters", MapOf(String, Any), "过滤条件")
			Field(3, "page", Int, "页码")
			Field(4, "pageSize", Int, "每页数量")
			Required("fields", "page", "pageSize")
		})
		Result(func() {
			Attribute("users", ArrayOf(User))
			Attribute("total", Int64, "总数")
			Required("users", "total")
		})
		HTTP(func() {
			GET("/users")
			Params(func() {
				Param("fields")
				Param("filters")
				Param("page")
				Param("pageSize")
			})
		})
	})
})

// 保留现有的 calc 服务定义
var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("multiply", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/multiply/{a}/{b}")
		})

		GRPC(func() {
		})
	})
})

var JWTAuth = JWTSecurity("jwt", func() {
	Description("JWT认证")
	Scope("api:access", "API访问权限")
})

var _ = Service("login", func() {
	Description("登录服务")

	Security(JWTAuth, func() {
		Scope("api:access")
	})
	Method("login", func() {
		NoSecurity()
		Payload(func() {
			Field(1, "username", String, "用户名")
			Field(2, "password", String, "密码")
			Required("username", "password")
		})
		Result(LoginResult) // 更新结果为 LoginResult
		HTTP(func() {
			POST("/login")
			Response(StatusOK)
		})
	})

	Method("logout", func() {
		Payload(func() {
			Token("token", String, "jwt token info")
			Required("token")
		})
		Result(Empty)
		HTTP(func() {
			POST("/logout")
			Response(StatusOK)
		})
	})

	Method("currentUser", func() {
		Payload(func() {
			Token("token", String, "jwt token info")
			Required("token")
		})
		Result(User)
		HTTP(func() {
			GET("/current-user")
			Response(StatusOK)
		})
	})
})

// 定义 User 类型
var User = Type("User", func() {
	Attribute("id", String, "用户ID")
	Attribute("username", String, "用户名")
	Attribute("email", String, "用户邮箱")
	Required("id", "username", "email")
})

// 定义 LoginResult 类型
var LoginResult = Type("LoginResult", func() {
	Attribute("token", String, "JWT令牌")
	Attribute("refresh_token", String, "刷新令牌")
	Attribute("expires_in", Int, "有效时间（秒）")
	Required("token", "refresh_token", "expires_in")
})
