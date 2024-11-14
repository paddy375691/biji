package service

import (
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	nwv1 "k8s.io/api/networking/v1"
	"sort"
	"strings"
	"time"
)

//dataSelector 用于封装排序、过滤、分页的数据类型
type dataSelector struct {
	GenericDataList []DataCell
	DataSelect      *DataSelectQuery
}
//DataCell接口，用于各种资源List的类型转换，转换后可以使用dataSelector的排序、过滤、分页方法
type DataCell interface {
	GetCreation() time.Time
	GetName()     string
}

//定义过滤和分页的结构体,过滤：Name，分页：Limit和Page
//Limit是单页的数据条数
//Page是第几页
type DataSelectQuery struct{
	Filter    *FilterQuery
	Paginate  *PaginateQuery
}

type FilterQuery struct {
	Name string
}

type PaginateQuery struct {
	Limit int
	Page  int
}

//实现自定义结构的排序，需要重写Len、Swap、Less方法
//Len方法用于获取数组的长度
func(d *dataSelector) Len() int {
	return len(d.GenericDataList)
}

//Swap方法用于数据比较大小之后的位置变更
func(d *dataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
	//temp := d.GenericDataList[i]
	//d.GenericDataList[i] = d.GenericDataList[j]
	//d.GenericDataList[j] = temp
}

//Less方法用于比较大小
func(d *dataSelector) Less(i, j int) bool {
	a := d.GenericDataList[i].GetCreation()
	b := d.GenericDataList[j].GetCreation()

	return b.Before(a)
}

//重写以上三个方法，使用sort.Sort进行排序
func(d *dataSelector) Sort() *dataSelector {
	sort.Sort(d)
	return d
}

//Filter方法用于过滤数据，比较数据的Name属性，若包含，则返回
func(d *dataSelector) Filter() *dataSelector {
	//判断入参是否为空，若为空，则返回所有数据
	if d.DataSelect.Filter.Name == "" {
		return d
	}
	//若不为空，则按照入参Name进行过滤
	//声明一个新的数组，若Name包含，则把数据放进数组，返回出去
	filtered := []DataCell{}
	for _,value := range d.GenericDataList {
		//定义是否匹配的标签变量,默认是匹配
		matches := true
		objName := value.GetName()
		if !strings.Contains(objName, d.DataSelect.Filter.Name) {
			matches = false
			continue
		}
		if matches {
			filtered = append(filtered, value)
		}
	}
	d.GenericDataList = filtered
	return d
}

//Paginate方法用于数组的分页，根据Limit和Page的传参，取一定范围内的数据，返回
func(d *dataSelector) Paginate() *dataSelector {
	//根据Limit和Page的入参，定义快捷变量
	limit := d.DataSelect.Paginate.Limit
	page := d.DataSelect.Paginate.Page
	//检验参数的合法性
	if limit <= 0 || page <= 0 {
		return d
	}
	//定义取数范围需要的startIndex和endIndex
	//举例，有个25个元素的数组，limit是10，page是3，startIndex是20，endIndex是29（endIndex是24）
	startIndex := limit * ( page -1 )
	endIndex := limit * page -1

	//处理endIndex
	if endIndex > len(d.GenericDataList) {
		endIndex = len(d.GenericDataList)
	}

	d.GenericDataList = d.GenericDataList[startIndex:endIndex]

	return d
}

//定义podCell，重写GetCreation和GetName方法后，可以进行数据转换
//corev1.Pod -> podCell -> DataCell
//appsv1.Deployment -> deployCell -> DataCell
type podCell corev1.Pod

//重写DataCell接口的两个方法
func(p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func(p podCell) GetName() string {
	return p.Name
}

type deploymentCell appsv1.Deployment

func(d deploymentCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func(d deploymentCell) GetName() string {
	return d.Name
}

type serviceCell corev1.Service

func(s serviceCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func(s serviceCell) GetName() string {
	return s.Name
}

type ingressCell nwv1.Ingress

func(i ingressCell) GetCreation() time.Time {
	return i.CreationTimestamp.Time
}

func(i ingressCell) GetName() string {
	return i.Name
}

type namespaceCell corev1.Namespace

func(n namespaceCell) GetCreation() time.Time {
	return n.CreationTimestamp.Time
}

func(n namespaceCell) GetName() string {
	return n.Name
}