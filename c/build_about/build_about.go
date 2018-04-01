package build_about

/*
const char* build_time(void)
{
    static const char* psz_build_time = "["__DATE__ "  " __TIME__ "]";
    return psz_build_time;
}
 */

import "C"

func GetBuildTime()string{
	return C.GoString(C.build_time())
}