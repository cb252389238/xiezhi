package gojieba

import (
	"github.com/cb252389238/xiezhi/gojieba/deps/cppjieba"
	"github.com/cb252389238/xiezhi/gojieba/deps/limonp"
	"github.com/cb252389238/xiezhi/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
