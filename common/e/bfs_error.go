package e


type BfsError interface {
	Error() string
	GetSCode() int
	GetHttpStatus() int
	GetNote() string
	GetRaw() interface{}
}


func New(httpStatus int,scode int,note string,raw interface{}) BfsError{

	return &bfsError{HttpStatus:httpStatus,SCode:scode,Note:note,Raw:raw}
}
func NewSimple(httpStatus int,scode int,note string) BfsError{

	return &bfsError{HttpStatus:httpStatus,SCode:scode,Note:note}
}
/**
 * 自定义异常实现
 */
type bfsError struct {
	HttpStatus int //自定义响应的 HttpStatusCode
	SCode int //自定义错误码
	Note string //自定义错误数据
	Raw interface{} //原始错误数据
}

 func (e *bfsError) Error() string {
	return e.Note
}
func (e *bfsError) GetSCode() int {
	return e.SCode
}
func (e *bfsError) GetHttpStatus() int {
	return e.HttpStatus
}
func (e *bfsError) GetNote() string {
	return e.Note
}
func (e *bfsError) GetRaw() interface{} {
	return e.Raw
}