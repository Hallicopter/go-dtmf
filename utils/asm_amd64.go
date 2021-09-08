// +build amd64

package utils

//defined in asm_amd64.s
func hasSSE4() bool

var useSSE4 = hasSSE4()
