package uri

var UriParamsMap=make(map[string]string)
var WhiteUriListMap=make(map[string]string)
func init(){
	UriParamsMap["dgst"]=":dgst"
	UriParamsMap["name"]=":name"
	UriParamsMap["afid"]=":afid"
	UriParamsMap["algotype"]=":algotype"
	UriParamsMap["filename"]=":filename"
	UriParamsMap["fullname"]=":fullname"
	UriParamsMap["agfid"]=":agfid"
	UriParamsMap["param"]=":param"

	WhiteUriListMap["/qn/version"]="/qn/version"
	WhiteUriListMap["/qn/sys/time"]="/qn/sys/time"
	WhiteUriListMap["/qn/sys/healthy"]="/qn/sys/healthy"
	WhiteUriListMap["/metrics"]="/metrics"
	WhiteUriListMap["/pn/addresses"]="/pn/addresses"

	WhiteUriListMap["/un/file"]="/un/file"
	WhiteUriListMap["/gn/gfid"]="/gn/gfid"


}
