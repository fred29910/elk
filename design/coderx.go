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

var _ = Service("coderx", func() {
	Description("The coderx service provides code generation and management functionalities.")

	Method("generateCode", func() {
		Payload(func() {
			Field(1, "language", String, "Programming language")
			Field(2, "specification", String, "Code specification")
			Required("language", "specification")
		})
		Result(String)
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
