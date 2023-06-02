package main

import (
	"context"
	"fmt"
	"github.com/Code-Hex/pget"
	chainstoragesdk "github.com/paradeum-team/chainstorage-sdk/sdk"
	sdkcode "github.com/paradeum-team/chainstorage-sdk/sdk/code"
	"github.com/paradeum-team/chainstorage-sdk/sdk/model"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/ulule/deepcopier"
	"net/http"
	"os"
	"text/template"
	"time"
)

func init() {
	objectListCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	objectListCmd.Flags().StringP("Object", "r", "", "对象名称")
	objectListCmd.Flags().StringP("Cid", "c", "", "Cid")
	objectListCmd.Flags().IntP("Offset", "o", 10, "查询偏移量")

	objectRenameCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	objectRenameCmd.Flags().StringP("Object", "o", "", "对象名称")
	objectRenameCmd.Flags().StringP("Cid", "c", "", "Cid")
	objectRenameCmd.Flags().StringP("Rename", "r", "", "重命名")
	objectRenameCmd.Flags().BoolP("Force", "f", false, "有冲突的时候强制覆盖")

	objectRemoveCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	objectRemoveCmd.Flags().StringP("Object", "o", "", "对象名称")
	objectRemoveCmd.Flags().StringP("Cid", "c", "", "Cid")
	objectRemoveCmd.Flags().BoolP("Force", "f", false, "有冲突的时候强制覆盖")

	objectDownloadCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	objectDownloadCmd.Flags().StringP("Object", "o", "", "对象名称")
	objectDownloadCmd.Flags().StringP("Cid", "c", "", "Cid")
	objectDownloadCmd.Flags().BoolP("Target", "t", false, "输出路径")
}

// region Object List

var objectListCmd = &cobra.Command{
	Use:     "lso",
	Short:   "lso",
	Long:    "List object",
	Example: "gcscmd ls cs://BUCKET [--name=<name>] [--cid=<cid>] [--Offset=<Offset>]",

	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Help()
		//fmt.Printf("%s %s\n", cmd.Name(), strconv.Itoa(offset))
		objectListRun(cmd, args)
	},
}

func objectListRun(cmd *cobra.Command, args []string) {
	itemName := ""
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

	// 对象名称
	objectName, err := cmd.Flags().GetString("Object")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	// 对象CID
	cid, err := cmd.Flags().GetString("Cid")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	// 设置参数
	if len(objectName) > 0 {
		itemName = objectName
	} else if len(cid) > 0 {
		itemName = cid
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

	sdk, err := chainstoragesdk.New(&applicationConfig)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)

	}

	// 确认桶数据有效性
	respBucket, err := sdk.Bucket.GetBucketByName(bucketName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rm", err, args)
	}

	// 桶ID
	bucketId := respBucket.Data.Id

	// 列出桶对象
	response, err := sdk.Object.GetObjectList(bucketId, itemName, pageSize, pageIndex)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)

	}

	objectListRunOutput(cmd, args, response)
	//fmt.Printf("response:%+v\n", response)

}

func objectListRunOutput(cmd *cobra.Command, args []string, resp model.ObjectPageResponse) {
	code := resp.Code
	if code != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("ls", err, args)
	}

	respData := resp.Data
	objectListOutput := ObjectListOutput{
		RequestId: resp.RequestId,
		Code:      resp.Code,
		Msg:       resp.Msg,
		Status:    resp.Status,
		Count:     respData.Count,
		PageIndex: respData.PageIndex,
		PageSize:  respData.PageSize,
		List:      []ObjectOutput{},
	}

	if len(respData.List) > 0 {
		for i, _ := range respData.List {
			objectOutput := ObjectOutput{}
			deepcopier.Copy(respData.List[i]).To(&objectOutput)

			// 创建时间
			objectOutput.CreatedDate = objectOutput.CreatedAt.Format("2006-01-02")
			objectListOutput.List = append(objectListOutput.List, objectOutput)
		}
	}

	//	查看对象
	//	通过命令查询固定桶内对象
	//
	//	模版
	//
	//	gcscmd ls cs://BUCKET [--name=<name>] [--cid=<cid>] [--Offset=<Offset>]
	//	BUCKET
	//
	//	桶名称
	//
	//	cid
	//
	//	对象 CID
	//
	//	name
	//
	//	对象名
	//
	//	Offset
	//
	//	查询偏移量
	//
	//	命令行例子
	//
	//	桶内所有文件查询
	//
	//	gcscmd ls cs://bbb
	//	桶内对应 cid 查询
	//
	//	gcscmd ls cs://bbb  --cid QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo
	//	桶内对象名查询
	//
	//	gcscmd ls cs://bbb  --name Tarkov.mp4
	//	响应
	//
	//	有
	//
	//	total 1
	//	QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo  666993487  2023-04-03 Tarkov.mp4
	//	无
	//
	//	total 0
	//Status: 200

	templateContent := `
total {{len .List}}
{{- if eq (len .List) 0}}
Status: {{.Code}}
{{- else}}
{{- range .List}}
{{.ObjectCid}} {{.ObjectSize}} {{.CreatedDate}} {{.ObjectName}}
{{- end}}
{{- end}}
`
	//Id             int       `json:"id" comment:"对象ID"`
	//BucketId       int       `json:"bucketId" comment:"桶主键"`
	//ObjectName     string    `json:"objectName" comment:"对象名称（255字限制）"`
	//ObjectTypeCode int       `json:"objectTypeCode" comment:"对象类型编码"`
	//ObjectSize     int64     `json:"objectSize" comment:"对象大小（字节）"`
	//IsMarked       int       `json:"isMarked" comment:"星标（1-已标记，0-未标记）"`
	//ObjectCid      string    `json:"objectCid" comment:"对象CID"`
	//CreatedAt      time.Time `json:"createdAt" comment:"创建时间"`
	//UpdatedAt      time.Time `json:"updatedAt" comment:"最后更新时间"`
	//CreatedDate    string    `json:"createdDate" comment:"创建日期"`

	//	templateContent := `
	//total {{len .List}}{{if eq (len .List) 0}}
	//Status: {{.Code}}{{else}}{{range .List}}
	//{{.ObjectAmount}} {{.StorageNetwork}} {{.ObjectPrinciple}} {{.UsedSpace}} {{.CreatedAt.Format "20060102"}} {{.ObjectName}}{{end}}
	//{{end}}
	//`

	t, err := template.New("objectListTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("ls", err, args)

	}

	err = t.Execute(os.Stdout, objectListOutput)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("ls", err, args)

	}
}

type ObjectListOutput struct {
	RequestId string         `json:"requestId,omitempty"`
	Code      int32          `json:"code,omitempty"`
	Msg       string         `json:"msg,omitempty"`
	Status    string         `json:"status,omitempty"`
	Count     int            `json:"count,omitempty"`
	PageIndex int            `json:"pageIndex,omitempty"`
	PageSize  int            `json:"pageSize,omitempty"`
	List      []ObjectOutput `json:"list,omitempty"`
}

type ObjectOutput struct {
	Id             int       `json:"id" comment:"对象ID"`
	BucketId       int       `json:"bucketId" comment:"桶主键"`
	ObjectName     string    `json:"objectName" comment:"对象名称（255字限制）"`
	ObjectTypeCode int       `json:"objectTypeCode" comment:"对象类型编码"`
	ObjectSize     int64     `json:"objectSize" comment:"对象大小（字节）"`
	IsMarked       int       `json:"isMarked" comment:"星标（1-已标记，0-未标记）"`
	ObjectCid      string    `json:"objectCid" comment:"对象CID"`
	CreatedAt      time.Time `json:"createdAt" comment:"创建时间"`
	UpdatedAt      time.Time `json:"updatedAt" comment:"最后更新时间"`
	CreatedDate    string    `json:"createdDate" comment:"创建日期"`
}

// endregion Object List

// region Object Rename

var objectRenameCmd = &cobra.Command{
	Use:     "rn",
	Short:   "rn",
	Long:    "rename object",
	Example: "gcscmd rn cs://BUCKET] [--name=<name>] [--cid=<cid>] [--rename=<rename>] [--force]",

	Run: func(cmd *cobra.Command, args []string) {
		objectRenameRun(cmd, args)
	},
}

func objectRenameRun(cmd *cobra.Command, args []string) {

	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	// 对象名称
	objectName, err := cmd.Flags().GetString("Object")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	//// 对象CID
	//cid, err := cmd.Flags().GetString("Cid")
	//if err != nil {
	//	//todo: log detail error?
	//	//panic(err)
	//	//fmt.Printf("error:%+v\n", err)
	//	processError("ls", err, args)
	//}

	// 对象名称
	newName, err := cmd.Flags().GetString("Rename")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	// todo: return successed
	if newName == objectName {

	}

	// 强制覆盖
	force, err := cmd.Flags().GetBool("Force")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("rn", err, args)

	}

	sdk, err := chainstoragesdk.New(&applicationConfig)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rn", err, args)

	}

	// 确认桶数据有效性
	respBucket, err := sdk.Bucket.GetBucketByName(bucketName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rm", err, args)
	}

	// 桶ID
	bucketId := respBucket.Data.Id

	// 确认对象数据有效性
	respObject, err := sdk.Object.GetObjectByName(bucketId, objectName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rn", err, args)
	}

	// 对象ID
	objectId := respObject.Data.Id

	// 重命名对象
	response, err := sdk.Object.RenameObject(objectId, newName, force)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rn", err, args)

	}

	objectRenameRunOutput(cmd, args, response)
}

func objectRenameRunOutput(cmd *cobra.Command, args []string, resp model.ObjectRenameResponse) {
	respCode := int(resp.Code)

	if int(respCode) == sdkcode.ErrObjectNameConflictInBucket.Code() {
		err := errors.New("Error: conflicting rename filename, add --force to confirm overwrite\n")
		processError("rn", err, args)
	} else if respCode != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("rn", err, args)
	}

	objectRenameOutput := ObjectRenameOutput{
		RequestId: resp.RequestId,
		Code:      resp.Code,
		Msg:       resp.Msg,
		Status:    resp.Status,
	}

	//	删除对象
	//	通过命令删除固定桶内对象
	//
	//	模版
	//
	//	gcscmd rm cs://[BUCKET] [--name=<name>] [--cid=<cid>] [--force]
	//	BUCKET
	//
	//	桶名称
	//
	//	cid
	//
	//	添加对应的 CID
	//
	//	name
	//
	//	对象名
	//
	//	force
	//
	//	无添加筛选条件或命中多的对象时需要添加
	//
	//	命令行例子
	//
	//	清空桶
	//
	//	gcscmd rm cs://bbb --force
	//	使用对象名删除单文件
	//
	//	gcscmd rm  cs://bbb --name Tarkov.mp4
	//	使用模糊查询删除对象
	//
	//	gcscmd rm  cs://bbb --name .mp4 --force
	//	使用对象名删除单目录
	//
	//	gcscmd rm  cs://bbb --name aaa
	//	使用CID删除单对象
	//
	//	gcscmd rm  cs://bbb --cid QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo
	//	使用 CID 删除多个对象(命中多个对象时加)
	//
	//	gcscmd rm  cs://bbb --cid QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo --force
	//	响应
	//
	//	成功
	//
	//Status: 200
	//	多对象没有添加 force
	//
	//Error: multiple object  are matching this query, add --force to confirm the bulk removal

	templateContent := `
成功
Status: {{.Code}}
`

	t, err := template.New("objectRenameTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("rn", err, args)

	}

	err = t.Execute(os.Stdout, objectRenameOutput)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("rn", err, args)

	}
}

type ObjectRenameOutput struct {
	RequestId string       `json:"requestId,omitempty"`
	Code      int32        `json:"code,omitempty"`
	Msg       string       `json:"msg,omitempty"`
	Status    string       `json:"status,omitempty"`
	Data      ObjectOutput `json:"objectOutput,omitempty"`
}

// endregion Object Rename

// region Object Remove

var objectRemoveCmd = &cobra.Command{
	Use:     "rmo",
	Short:   "rmo",
	Long:    "remove object",
	Example: "gcscmd rmo cs://BUCKET] [--name=<name>] [--cid=<cid>] [--remove=<remove>] [--force]",

	Run: func(cmd *cobra.Command, args []string) {
		objectRemoveRun(cmd, args)
	},
}

func objectRemoveRun(cmd *cobra.Command, args []string) {

	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	// 对象名称
	objectName, err := cmd.Flags().GetString("Object")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	//// 对象CID
	//cid, err := cmd.Flags().GetString("Cid")
	//if err != nil {
	//	//todo: log detail error?
	//	//panic(err)
	//	//fmt.Printf("error:%+v\n", err)
	//	processError("ls", err, args)
	//}

	//// 强制覆盖
	//force, err := cmd.Flags().GetBool("Force")
	//if err != nil {
	//	//todo: log detail error?
	//	//panic(err)
	//	//fmt.Printf("error:%+v\n", err)
	//	processError("rmo", err, args)
	//
	//}

	sdk, err := chainstoragesdk.New(&applicationConfig)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rmo", err, args)

	}

	// 确认桶数据有效性
	respBucket, err := sdk.Bucket.GetBucketByName(bucketName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rm", err, args)
	}

	// 桶ID
	bucketId := respBucket.Data.Id

	// 确认对象数据有效性
	respObject, err := sdk.Object.GetObjectByName(bucketId, objectName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rmo", err, args)
	}

	// 对象ID
	objectId := respObject.Data.Id

	// 重命名对象
	response, err := sdk.Object.RemoveObject([]int{objectId})
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rmo", err, args)

	}

	objectRemoveRunOutput(cmd, args, response)
}

func objectRemoveRunOutput(cmd *cobra.Command, args []string, resp model.ObjectRemoveResponse) {
	respCode := int(resp.Code)

	if respCode != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("rmo", err, args)
	}

	objectRemoveOutput := ObjectRemoveOutput{
		RequestId: resp.RequestId,
		Code:      resp.Code,
		Msg:       resp.Msg,
		Status:    resp.Status,
	}

	//	删除对象
	//	通过命令删除固定桶内对象
	//
	//	模版
	//
	//	gcscmd rm cs://[BUCKET] [--name=<name>] [--cid=<cid>] [--force]
	//	BUCKET
	//
	//	桶名称
	//
	//	cid
	//
	//	添加对应的 CID
	//
	//	name
	//
	//	对象名
	//
	//	force
	//
	//	无添加筛选条件或命中多的对象时需要添加
	//
	//	命令行例子
	//
	//	清空桶
	//
	//	gcscmd rm cs://bbb --force
	//	使用对象名删除单文件
	//
	//	gcscmd rm  cs://bbb --name Tarkov.mp4
	//	使用模糊查询删除对象
	//
	//	gcscmd rm  cs://bbb --name .mp4 --force
	//	使用对象名删除单目录
	//
	//	gcscmd rm  cs://bbb --name aaa
	//	使用CID删除单对象
	//
	//	gcscmd rm  cs://bbb --cid QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo
	//	使用 CID 删除多个对象(命中多个对象时加)
	//
	//	gcscmd rm  cs://bbb --cid QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo --force
	//	响应
	//
	//	成功
	//
	//Status: 200
	//	多对象没有添加 force
	//
	//Error: multiple object  are matching this query, add --force to confirm the bulk removal

	templateContent := `
成功
Status: {{.Code}}
`

	t, err := template.New("objectRemoveTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("rmo", err, args)

	}

	err = t.Execute(os.Stdout, objectRemoveOutput)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("rmo", err, args)

	}
}

type ObjectRemoveOutput struct {
	RequestId string       `json:"requestId,omitempty"`
	Code      int32        `json:"code,omitempty"`
	Msg       string       `json:"msg,omitempty"`
	Status    string       `json:"status,omitempty"`
	Data      ObjectOutput `json:"objectOutput,omitempty"`
}

// endregion Object Remove

// region Object Download

var objectDownloadCmd = &cobra.Command{
	Use:     "get",
	Short:   "get",
	Long:    "download object",
	Example: "gcscmd get cs://BUCKET [--name=<name>] [--cid=<cid>]",

	Run: func(cmd *cobra.Command, args []string) {
		objectDownloadRun(cmd, args)
	},
}

func objectDownloadRun(cmd *cobra.Command, args []string) {

	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	// 对象名称
	objectName, err := cmd.Flags().GetString("Object")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	//// 对象CID
	//cid, err := cmd.Flags().GetString("Cid")
	//if err != nil {
	//	//todo: log detail error?
	//	//panic(err)
	//	//fmt.Printf("error:%+v\n", err)
	//	processError("ls", err, args)
	//}

	//// 输出路径
	//target, err := cmd.Flags().GetBool("Target")
	//if err != nil {
	//	//todo: log detail error?
	//	//panic(err)
	//	//fmt.Printf("error:%+v\n", err)
	//	processError("get", err, args)
	//
	//}

	sdk, err := chainstoragesdk.New(&applicationConfig)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("get", err, args)
	}

	// 确认桶数据有效性
	respBucket, err := sdk.Bucket.GetBucketByName(bucketName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("rm", err, args)
	}

	// 桶ID
	bucketId := respBucket.Data.Id

	// 确认对象数据有效性
	response, err := sdk.Object.GetObjectByName(bucketId, objectName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("get", err, args)
	}

	// 对象ID
	//objectId := response.Data.Id
	//objectName := response.Data.ObjectName
	objectCid := response.Data.ObjectCid
	downloadEndpoint := "https://test-ipfs-gateway.netwarps.com/ipfs/"
	downloadUrl := fmt.Sprintf("%s%s", downloadEndpoint, objectCid)

	// todo:
	downloadUrl = "https://test-ipfs-gateway.netwarps.com/ipfs/bafybeiguyrqm6z76mrhntk64fiwwdpjqv64ny3ugw64owznlbeotknvypa"
	cli := pget.New()
	cli.URLs = []string{downloadUrl}
	cli.Output = objectName

	version := "CMD"

	if err := cli.Run(context.Background(), version, []string{"-t", "30"}); err != nil {
		if cli.Trace {
			fmt.Fprintf(os.Stderr, "Error:\n%+v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "Error:\n  %v\n", err)
		}
		processError("get", err, args)
	}

	objectDownloadRunOutput(cmd, args, response)
}

func objectDownloadRunOutput(cmd *cobra.Command, args []string, resp model.ObjectCreateResponse) {
	respCode := int(resp.Code)

	if respCode != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("get", err, args)
	}

	//objectDownloadOutput := ObjectDownloadOutput{
	//	RequestId: resp.RequestId,
	//	Code:      resp.Code,
	//	Msg:       resp.Msg,
	//	Status:    resp.Status,
	//}

	//	下载对象
	//	通过命令下载固定桶内的对象数据
	//
	//	模版
	//
	//	gcscmd get cs://BUCKET [--name=<name>] [--cid=<cid>]
	//	BUCKET
	//
	//	桶名称
	//
	//	carfile
	//
	//	car文件标识
	//
	//	命令行例子
	//
	//	桶内对应 cid
	//
	//	gcscmd get cs://bbb --cid QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo
	//	桶内对象名
	//
	//	gcscmd get cs://bbb  --name Tarkov.mp4
	//	响应
	//
	//	过程
	//
	//	################                                                                15%
	//		QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo        Tarkov.mp4
	//	完成
	//
	//CID:    QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo
	//Name:Tarkov.mp4

	templateContent := `
CID: {{.ObjectCid}}
Name: {{.ObjectName}}
`

	t, err := template.New("objectDownloadTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("get", err, args)
	}

	err = t.Execute(os.Stdout, resp)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("get", err, args)
	}
}

type ObjectDownloadOutput struct {
	RequestId string       `json:"requestId,omitempty"`
	Code      int32        `json:"code,omitempty"`
	Msg       string       `json:"msg,omitempty"`
	Status    string       `json:"status,omitempty"`
	Data      ObjectOutput `json:"objectOutput,omitempty"`
}

// endregion Object Download
