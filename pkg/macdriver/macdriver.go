package macdriver

import (
	"github.com/ebitengine/purego/objc"
)

var (
	SEL_alloc            = objc.RegisterName("alloc")
	SEL_new              = objc.RegisterName("new")
	SEL_init             = objc.RegisterName("init")
	SEL_initWithCapacity = objc.RegisterName("initWithCapacity:")
	SEL_release          = objc.RegisterName("release")
	SEL_length           = objc.RegisterName("length")
	SEL_defaultCenter    = objc.RegisterName("defaultCenter")
)
