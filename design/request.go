package design

import (
	. "goa.design/goa/v3/dsl"
)

var UpdateAccount = Type("UpdateAccount", func() {
	// Attribute("name", String, func() {
	// 	MinLength(3)
	// 	MaxLength(20)
	// 	// Format(FormatJSON)
	// })
	// Attribute("email", String, func() {
	// 	Format(FormatEmail)
	// })
	// Attributes(func() {
	Field(1, "name", String, func() {
		MinLength(3)
		MaxLength(20)
	})
	Field(2, "email", String, func() {
		Format(FormatEmail)
	})
	// })
	Required("name", "email")
})

var UpdateAccountResult = ResultType("application/vnd.create", func() {
	Attributes(func() {
		Field(1, "name", String, "Name of the created resource")
		Field(2, "href", String, "Href of the created resource")
	})
})
