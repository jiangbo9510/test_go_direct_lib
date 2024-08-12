package test_go_direct_lib

import "github.com/jiangbo9510/test_base_lib"

func SetCommon(v string) {
	test_base_lib.Common = v
}

func GetCommon() string {
	return test_base_lib.GetCommon()
}
