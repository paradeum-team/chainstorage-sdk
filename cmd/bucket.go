package main

import (
	chainstoragesdk "github.com/paradeum-team/chainstorage-sdk/sdk"
	sdkcode "github.com/paradeum-team/chainstorage-sdk/sdk/code"
	"github.com/paradeum-team/chainstorage-sdk/sdk/consts"
	"github.com/paradeum-team/chainstorage-sdk/sdk/model"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/ulule/deepcopier"
	"net/http"
	"os"
	"regexp"
	"text/template"
	"time"
)

// var (
//
//	offset int
//
// )

func init() {
	bucketListCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	bucketListCmd.Flags().IntP("Offset", "o", 10, "查询偏移量")

	bucketCreateCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	bucketCreateCmd.Flags().IntP("Storage", "s", 10001, "存储网络编码")
	bucketCreateCmd.Flags().IntP("Principle", "p", 10001, "桶策略编码")

	bucketRemoveCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	bucketRemoveCmd.Flags().BoolP("Force", "f", false, "如果有数据，先清空再删除桶")

	bucketEmptyCmd.Flags().StringP("Bucket", "b", "", "桶名称")
}

// region Bucket List

var bucketListCmd = &cobra.Command{
	Use:     "ls",
	Short:   "ls",
	Long:    "List links from object or bucket",
	Example: "gcscmd ls [--Offset=<Offset>]",

	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Help()
		//fmt.Printf("%s %s\n", cmd.Name(), strconv.Itoa(offset))
		bucketListRun(cmd, args)
	},
}

func bucketListRun(cmd *cobra.Command, args []string) {
	bucketName := ""
	pageSize := 10
	pageIndex := 1

	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)

	}

	// 查询偏移量
	offset, err := cmd.Flags().GetInt("Offset")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)

	}

	if offset > 0 || offset < 1000 {
		pageSize = offset
	}

	sdk, err := chainstoragesdk.New(sdkCfgFile)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)

	}
	//fmt.Printf("sdk.myConfig:%+v\n", sdk.Config)

	// 列出桶对象
	response, err := sdk.Bucket.GetBucketList(bucketName, pageSize, pageIndex)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)

	}

	bucketListRunOutput(cmd, args, response)
	//fmt.Printf("response:%+v\n", response)

}

func bucketListRunOutput(cmd *cobra.Command, args []string, resp model.BucketPageResponse) {
	code := resp.Code
	if code != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("ls", err, args)
	}

	respData := resp.Data
	bucketListOutput := BucketListOutput{
		RequestId: resp.RequestId,
		Code:      resp.Code,
		Msg:       resp.Msg,
		Status:    resp.Status,
		Count:     respData.Count,
		PageIndex: respData.PageIndex,
		PageSize:  respData.PageSize,
		List:      []BucketOutput{},
	}

	if len(respData.List) > 0 {
		for i, _ := range respData.List {
			bucketOutput := BucketOutput{}
			deepcopier.Copy(respData.List[i]).To(&bucketOutput)

			// 存储网络
			bucketOutput.StorageNetwork = consts.StorageNetworkCodeMapping[bucketOutput.StorageNetworkCode]

			// 桶策略
			bucketOutput.BucketPrinciple = consts.BucketPrincipleCodeMapping[bucketOutput.BucketPrincipleCode]

			// 创建时间
			bucketOutput.CreatedDate = bucketOutput.CreatedAt.Format("2006-01-02")
			bucketListOutput.List = append(bucketListOutput.List, bucketOutput)
		}
	}

	//Id                  int       `json:"id" comment:"桶ID"`
	//UserId              int       `json:"userId" comment:"用户ID"`
	//BucketName          string    `json:"bucketName" comment:"桶名称（3-63字长度限制）"`
	//StorageNetworkCode  int       `json:"storageNetworkCode" comment:"存储网络编码（10001-IPFS）"`
	//BucketPrincipleCode int       `json:"bucketPrincipleCode" comment:"桶策略编码（10001-公开，10000-私有）"`
	//UsedSpace           int64     `json:"usedSpace" comment:"已使用空间（字节）"`
	//ObjectAmount        int       `json:"objectAmount" comment:"对象数量"`
	//Status              int       `json:"status" comment:"记录状态（0-有效，1-删除）"`
	//CreatedAt           time.Time `json:"createdAt" comment:"创建时间"`
	//UpdatedAt           time.Time `json:"updatedAt" comment:"最后更新时间"`
	//
	//99 ipfs public 666993487 2023-04-03  aaa
	//total
	//桶数量
	//99
	//桶内对象数量
	//ipfs
	//存储网络
	//public
	//桶策略
	//666993487
	//大小单位字节
	//2023-04-03
	//创建时间
	//aaa
	//桶名称
	//
	//	templateContent := `
	//total {{len .Data.List}}{{if eq (len .Data.List) 0}}
	//Status: {{.Code}}{{else}}{{range .Data.List}}
	//{{.ObjectAmount}} {{.StorageNetworkCode}} {{.BucketPrincipleCode}} {{.UsedSpace}} {{.CreatedAt}} {{.BucketName}}{{end}}
	//{{end}}
	//`

	templateContent := `
total {{len .List}}
{{- if eq (len .List) 0}}
Status: {{.Code}}
{{- else}}
{{- range .List}}
{{.ObjectAmount}} {{.StorageNetwork}} {{.BucketPrinciple}} {{.UsedSpace}} {{.CreatedDate}} {{.BucketName}}
{{- end}}
{{- end}}
	`

	//	templateContent := `
	//total {{len .List}}{{if eq (len .List) 0}}
	//Status: {{.Code}}{{else}}{{range .List}}
	//{{.ObjectAmount}} {{.StorageNetwork}} {{.BucketPrinciple}} {{.UsedSpace}} {{.CreatedAt.Format "20060102"}} {{.BucketName}}{{end}}
	//{{end}}
	//`

	t, err := template.New("bucketListTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("ls", err, args)

	}

	err = t.Execute(os.Stdout, bucketListOutput)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("ls", err, args)

	}
}

type BucketListOutput struct {
	RequestId string         `json:"requestId,omitempty"`
	Code      int32          `json:"code,omitempty"`
	Msg       string         `json:"msg,omitempty"`
	Status    string         `json:"status,omitempty"`
	Count     int            `json:"count,omitempty"`
	PageIndex int            `json:"pageIndex,omitempty"`
	PageSize  int            `json:"pageSize,omitempty"`
	List      []BucketOutput `json:"list,omitempty"`
}

type BucketOutput struct {
	Id                  int       `json:"id" comment:"桶ID"`
	BucketName          string    `json:"bucketName" comment:"桶名称（3-63字长度限制）"`
	StorageNetworkCode  int       `json:"storageNetworkCode" comment:"存储网络编码（10001-IPFS）"`
	BucketPrincipleCode int       `json:"bucketPrincipleCode" comment:"桶策略编码（10001-公开，10000-私有）"`
	UsedSpace           int64     `json:"usedSpace" comment:"已使用空间（字节）"`
	ObjectAmount        int       `json:"objectAmount" comment:"对象数量"`
	CreatedAt           time.Time `json:"createdAt" comment:"创建时间"`
	StorageNetwork      string    `json:"storageNetwork" comment:"存储网络（10001-IPFS）"`
	BucketPrinciple     string    `json:"bucketPrinciple" comment:"桶策略（10001-公开，10000-私有）"`
	CreatedDate         string    `json:"createdDate" comment:"创建日期"`
}

// endregion Bucket List

// region Bucket Create

var bucketCreateCmd = &cobra.Command{
	Use:     "mb",
	Short:   "mb",
	Long:    "create bucket",
	Example: "gcscmd mb cs://[BUCKET] [--storageNetworkCode=<storageNetworkCode>] [--bucketPrincipleCode=<bucketPrincipleCode>]",

	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Help()
		//fmt.Printf("%s %s\n", cmd.Name(), strconv.Itoa(offset))
		bucketCreateRun(cmd, args)
	},
}

func bucketCreateRun(cmd *cobra.Command, args []string) {
	bucketName := ""
	//pageSize := 10
	//pageIndex := 1

	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}
	//bucketCreateCmd.Flags().IntP("Storage", "s", 10001, "存储网络编码")
	//bucketCreateCmd.Flags().IntP("Principle", "p", 10001, "桶策略编码")
	// 存储网络编码
	storage, err := cmd.Flags().GetInt("Storage")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("mb", err, args)

	}

	if storage > 0 {
		_, exist := consts.StorageNetworkCodeMapping[storage]
		if !exist {
			//todo: log detail error?
			//panic(err)
			//fmt.Printf("error:%+v\n", err)
			err := errors.Errorf("invalid storage network code, %d", storage)
			processError("mb", err, args)

		}
	}

	// 桶策略编码
	principle, err := cmd.Flags().GetInt("Principle")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("mb", err, args)

	}

	if principle > 0 {
		_, exist := consts.BucketPrincipleCodeMapping[principle]
		if !exist {
			//todo: log detail error?
			//panic(err)
			//fmt.Printf("error:%+v\n", err)
			err := errors.Errorf("invalid bucket principle code, %d", principle)
			processError("mb", err, args)

		}
	}

	sdk, err := chainstoragesdk.New(sdkCfgFile)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("mb", err, args)

	}
	//fmt.Printf("sdk.myConfig:%+v\n", sdk.Config)

	// 创建桶
	response, err := sdk.Bucket.CreateBucket(bucketName, storage, principle)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("mb", err, args)

	}

	bucketCreateRunOutput(cmd, args, response)
	//fmt.Printf("response:%+v\n", response)

}

func bucketCreateRunOutput(cmd *cobra.Command, args []string, resp model.BucketCreateResponse) {
	code := resp.Code
	if code != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("mb", err, args)
	}

	respData := resp.Data
	bucketCreateOutput := BucketCreateOutput{
		RequestId: resp.RequestId,
		Code:      resp.Code,
		Msg:       resp.Msg,
		Status:    resp.Status,
	}

	bucketOutput := BucketOutput{}
	deepcopier.Copy(respData).To(&bucketOutput)

	// 存储网络
	bucketOutput.StorageNetwork = consts.StorageNetworkCodeMapping[bucketOutput.StorageNetworkCode]

	// 桶策略
	bucketOutput.BucketPrinciple = consts.BucketPrincipleCodeMapping[bucketOutput.BucketPrincipleCode]

	// 创建时间 todo: timezone
	//bucketOutput.CreatedDate = bucketOutput.CreatedAt.Format("2006-01-02")
	bucketOutput.CreatedDate = bucketOutput.CreatedAt.Format("2006-01-02T15:04:05-07:00")
	bucketCreateOutput.Data = bucketOutput
	//bucketOutput.CreatedAt.Format(time.RFC3339)

	//	创建桶
	//	通过命令创建桶操作
	//
	//	模版
	//
	//	gcscmd mb cs://[BUCKET] [--storageNetworkCode=<storageNetworkCode>] [--bucketPrincipleCode=<bucketPrincipleCode>]
	//	BUCKET
	//
	//	桶名称
	//
	//	storageNetworkCode
	//
	//	存储网络代码
	//
	//	bucketPrincipleCode
	//
	//	同策略代码
	//
	//	命令行例子
	//
	//	gcscmd mb cs://bbb
	//	响应
	//
	//BUCKET: bbb
	//storageNetwork: ipfs
	//bucketPrinciple: public
	//createdAt: 2023-04-03T18:52:11.312+08:00

	templateContent := `
BUCKET: {{.BucketName}}
storageNetwork: {{.StorageNetwork}}
bucketPrinciple: {{.BucketPrinciple}}
createdAt: {{.CreatedDate}}
`

	t, err := template.New("bucketCreateTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("mb", err, args)

	}

	err = t.Execute(os.Stdout, bucketOutput)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("mb", err, args)

	}
}

type BucketCreateOutput struct {
	RequestId string       `json:"requestId,omitempty"`
	Code      int32        `json:"code,omitempty"`
	Msg       string       `json:"msg,omitempty"`
	Status    string       `json:"status,omitempty"`
	Data      BucketOutput `json:"bucketOutput,omitempty"`
}

// endregion Bucket Create

// region Bucket Remove

var bucketRemoveCmd = &cobra.Command{
	Use:     "rb",
	Short:   "rb",
	Long:    "remove bucket",
	Example: "gcscmd rb cs://[BUCKET] [--force]",

	Run: func(cmd *cobra.Command, args []string) {
		bucketRemoveRun(cmd, args)
	},
}

func bucketRemoveRun(cmd *cobra.Command, args []string) {
	bucketName := ""

	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)

	}

	// 强制移除桶
	force, err := cmd.Flags().GetBool("Force")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("rb", err, args)

	}

	sdk, err := chainstoragesdk.New(sdkCfgFile)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rb", err, args)

	}
	//fmt.Printf("sdk.myConfig:%+v\n", sdk.Config)

	// 确认桶数据有效性
	respBucket, err := sdk.Bucket.GetBucketByName(bucketName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rb", err, args)
	}

	// 桶ID
	bucketId := respBucket.Data.Id

	// 移除桶
	response, err := sdk.Bucket.RemoveBucket(bucketId, force)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rb", err, args)

	}

	bucketRemoveRunOutput(cmd, args, response)
}

func bucketRemoveRunOutput(cmd *cobra.Command, args []string, resp model.BucketRemoveResponse) {
	respCode := int(resp.Code)

	if int(respCode) == sdkcode.ErrBucketMustBeEmpty.Code() {
		err := errors.New("Error: Bucket contains objects, add --force to confirm deletion\n")
		processError("rb", err, args)
	} else if respCode != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("rb", err, args)
	}

	bucketRemoveOutput := BucketRemoveOutput{
		RequestId: resp.RequestId,
		Code:      resp.Code,
		Msg:       resp.Msg,
		Status:    resp.Status,
	}

	//	通过命令移除桶操作
	//
	//	模版
	//
	//	gcscmd rb cs://[BUCKET] [--force]
	//	BUCKET
	//
	//	桶名称
	//
	//	force
	//
	//	如果有数据，先清空再删除桶
	//
	//	命令行例子
	//
	//	gcscmd rb cs://bbb
	//	响应
	//
	//	成功
	//
	//Status: 200
	//	错误
	//
	//Error: Bucket contains objects, add --force to confirm deletion

	templateContent := `
成功
Status: {{.Code}}
`

	t, err := template.New("bucketRemoveTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("rb", err, args)

	}

	err = t.Execute(os.Stdout, bucketRemoveOutput)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("rb", err, args)

	}
}

type BucketRemoveOutput struct {
	RequestId string       `json:"requestId,omitempty"`
	Code      int32        `json:"code,omitempty"`
	Msg       string       `json:"msg,omitempty"`
	Status    string       `json:"status,omitempty"`
	Data      BucketOutput `json:"bucketOutput,omitempty"`
}

// endregion Bucket Remove

// region Bucket Empty

var bucketEmptyCmd = &cobra.Command{
	Use:     "rm",
	Short:   "rm",
	Long:    "empty bucket",
	Example: "gcscmd rm cs://[BUCKET]",

	Run: func(cmd *cobra.Command, args []string) {
		bucketEmptyRun(cmd, args)
	},
}

func bucketEmptyRun(cmd *cobra.Command, args []string) {
	bucketName := ""

	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)

	}

	//// 强制移除桶
	//force, err := cmd.Flags().GetBool("Force")
	//if err != nil {
	//	//todo: log detail error?
	//	//panic(err)
	//	//fmt.Printf("error:%+v\n", err)
	//	processError("rm", err, args)
	//
	//}

	sdk, err := chainstoragesdk.New(sdkCfgFile)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rm", err, args)

	}
	//fmt.Printf("sdk.myConfig:%+v\n", sdk.Config)

	// 确认桶数据有效性
	respBucket, err := sdk.Bucket.GetBucketByName(bucketName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rm", err, args)
	}

	// 桶ID
	bucketId := respBucket.Data.Id

	// 清空桶
	response, err := sdk.Bucket.EmptyBucket(bucketId)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rm", err, args)

	}

	bucketEmptyRunOutput(cmd, args, response)
}

func bucketEmptyRunOutput(cmd *cobra.Command, args []string, resp model.BucketEmptyResponse) {
	code := resp.Code

	if code != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("rm", err, args)
	}

	bucketEmptyOutput := BucketEmptyOutput{
		RequestId: resp.RequestId,
		Code:      resp.Code,
		Msg:       resp.Msg,
		Status:    resp.Status,
	}

	templateContent := `
成功
Status: {{.Code}}
`
	//	清空桶
	//	通过命令清空桶操作
	//
	//	模版
	//
	//	gcscmd rm cs://[BUCKET] [--force]
	//	BUCKET
	//
	//	桶名称
	//
	//	命令行例子
	//
	//	gcscmd rm cs://bbb [--force]
	//	响应
	//
	//	正确
	//
	//Status: 200
	//	错误
	//
	//	empty bucket operation, add --force to confirm emptying

	t, err := template.New("bucketEmptyTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("rm", err, args)

	}

	err = t.Execute(os.Stdout, bucketEmptyOutput)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("rm", err, args)

	}
}

type BucketEmptyOutput struct {
	RequestId string       `json:"requestId,omitempty"`
	Code      int32        `json:"code,omitempty"`
	Msg       string       `json:"msg,omitempty"`
	Status    string       `json:"status,omitempty"`
	Data      BucketOutput `json:"bucketOutput,omitempty"`
}

// endregion Bucket Empty

// 检查桶名称
func checkBucketName(bucketName string) error {
	if len(bucketName) < 3 || len(bucketName) > 63 {
		return sdkcode.ErrInvalidBucketName
	}

	// 桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试
	isMatch := regexp.MustCompile(`^[a-z0-9-]*$`).MatchString(bucketName)
	if !isMatch {
		return sdkcode.ErrInvalidBucketName
	}

	return nil
}
