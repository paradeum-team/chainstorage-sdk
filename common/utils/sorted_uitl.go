package bfsutils

import (
	"reflect"
	"sort"
)

//通用排序
//结构体排序，必须重写数组Len() Swap() Less()函数
type body_wrapper struct {
	Bodys []interface{}
	by func(p,q*interface{}) bool //内部Less()函数会用到
}


type IntSlice []int64

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


type SortBodyBy func(p, q* interface{}) bool //定义一个函数类型

//数组长度Len()
func (acw body_wrapper) Len() int  {
	return len(acw.Bodys)
}
//元素交换
func (acw body_wrapper) Swap(i,j int){
	acw.Bodys[i],acw.Bodys[j] = acw.Bodys[j],acw.Bodys[i]
}
//比较函数，使用外部传入的by比较函数
func (acw body_wrapper) Less(i,j int) bool {
	return acw.by(&acw.Bodys[i],&acw.Bodys[j])
}
//自定义排序字段，参考SortBodyByCreateTime中的传入函数
func SortBody(bodys [] interface{}, by SortBodyBy){
	sort.Sort(body_wrapper{bodys,by})
}


//降序排列
func DescSortBodyByFieldName(bodys [] interface{},fieldName string){
	sort.Sort(body_wrapper{bodys,func(p,q * interface{}) bool{
		v :=reflect.ValueOf(*p).Elem()
		i := v.FieldByName(fieldName)
		v =reflect.ValueOf(*q).Elem()
		j := v.FieldByName(fieldName)

		return  i.String() > j.String()
	}})
}

func DescSortBodyAudit(bodys [] interface{},fieldName string){
	sort.Sort(body_wrapper{bodys,func(p,q * interface{}) bool{
		v :=reflect.ValueOf(*p)
		i := v.FieldByName(fieldName)
		v =reflect.ValueOf(*q)
		j := v.FieldByName(fieldName)

		return  i.String() > j.String()
	}})
}

//升序排列
func AscendSortBodyFieldName(bodys [] interface{},fieldName string){
	sort.Sort(body_wrapper{bodys,func(p,q * interface{}) bool{
		v :=reflect.ValueOf(*p).Elem()
		i := v.FieldByName(fieldName)
		v =reflect.ValueOf(*q).Elem()
		j := v.FieldByName(fieldName)
		return  i.String() < j.String()
	}})
}