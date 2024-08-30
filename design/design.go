package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for multiplying numbers, a Goa teaser")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

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

	Method("divide", func() {
		Description("Divide returns the integral division of two integers.")
		Payload(func() {
			// Attribute("c", Int, "Left operand")
			// Attribute("d", Int, "Right operand")

			Field(1, "c", Int, "Left operand")
			Field(2, "d", Int, "Right operand")
			Required("c", "d")
		})
		Result(Int)

		// Error defines an error result.
		Error("DivByZero")

		HTTP(func() {
			GET("/div/{c}/{d}")
			// The HTTP status code for responses corresponding to
			// the "DivByZero" error is 400 Bad Request.
			// The default response for successful requests is
			// 200 OK.
			Response("DivByZero", StatusBadRequest)
		})

		GRPC(func() {
			// The gRPC code for results corresponding to the
			// "DivByZero" error is 3 (INVALID_ARGUMENT).
			// The default response for successful requests is
			// 0 OK.
			Response("DivByZero", CodeInvalidArgument)
		})
	})
})
