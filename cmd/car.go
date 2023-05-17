package main

import (
	"chainstoragesdk"
	"chainstoragesdk/code"
	"chainstoragesdk/consts"
	"chainstoragesdk/model"
	"chainstoragesdk/utils"
	"fmt"
	ipldfmt "github.com/ipfs/go-ipld-format"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/ulule/deepcopier"
	"net/http"
	"os"
	"text/template"
	"time"
)

func init() {
	carUploadCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	carUploadCmd.Flags().StringP("Object", "o", "", "上传对象路径")

	carImportCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	carImportCmd.Flags().StringP("Carfile", "c", "", "car文件标识")

	//objectRenameCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	//objectRenameCmd.Flags().StringP("Object", "o", "", "对象名称")
	//objectRenameCmd.Flags().StringP("Cid", "c", "", "Cid")
	//objectRenameCmd.Flags().StringP("Rename", "r", "", "重命名")
	//objectRenameCmd.Flags().BoolP("Force", "f", false, "有冲突的时候强制覆盖")
	//
	//objectRemoveCmd.Flags().StringP("Bucket", "b", "", "桶名称")
	//objectRemoveCmd.Flags().StringP("Object", "o", "", "对象名称")
	//objectRemoveCmd.Flags().StringP("Cid", "c", "", "Cid")
	//objectRemoveCmd.Flags().BoolP("Force", "f", false, "有冲突的时候强制覆盖")
}

// region CAR Upload

var carUploadCmd = &cobra.Command{
	Use:     "put",
	Short:   "put",
	Long:    "upload object",
	Example: "gcscmd put FILE[/DIR...] cs://BUCKET",

	Run: func(cmd *cobra.Command, args []string) {
		carUploadRun(cmd, args)
	},
}

func carUploadRun(cmd *cobra.Command, args []string) {

	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	// 上传对象路径
	objectName, err := cmd.Flags().GetString("Object")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	sdk, err := chainstoragesdk.New()
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("put", err, args)
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

	// 对象上传
	response, err := UploadData(sdk, bucketId, objectName)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("put", err, args)

	}

	carUploadRunOutput(cmd, args, response)
}

func carUploadRunOutput(cmd *cobra.Command, args []string, resp model.ObjectCreateResponse) {
	respCode := int(resp.Code)

	if respCode != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("put", err, args)
	}

	carUploadOutput := CarUploadOutput{
		RequestId: resp.RequestId,
		Code:      resp.Code,
		Msg:       resp.Msg,
		Status:    resp.Status,
	}

	//对象上传
	//通过命令向固定桶内上传对象，包括文件、目录
	//
	//模版
	//
	//gcscmd put FILE[/DIR...] cs://BUCKET
	//BUCKET
	//
	//桶名称
	//
	//命令行例子
	//
	//上传文件
	//
	//当前目录
	//
	//gcscmd put ./aaa.mp4 cs://bbb
	//绝对路径
	//
	//gcscmd put /home/pz/aaa.mp4 cs://bbb
	//相对路径
	//
	//gcscmd put ../pz/aaa.mp4 cs://bbb
	//上传目录
	//
	//gcscmd put ./aaaa cs://bbb
	//上传 carfile
	//
	//gcscmd put ./aaa.car cs://bbb --carfile
	//响应
	//
	//过程
	//
	//################                                                                15%
	//Tarkov.mp4
	//完成
	//
	//CID:    QmWgnG7pPjG31w328hZyALQ2BgW5aQrZyKpT47jVpn8CNo
	//Name:Tarkov.mp4
	//报错
	//
	//Error: This file is a car file, add --carfile to confirm uploading car

	templateContent := `
CID: {{.ObjectCid}}
Name: {{.ObjectName}}
`

	t, err := template.New("carUploadTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("put", err, args)

	}

	err = t.Execute(os.Stdout, carUploadOutput)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("put", err, args)

	}
}

type CarUploadOutput struct {
	RequestId string       `json:"requestId,omitempty"`
	Code      int32        `json:"code,omitempty"`
	Msg       string       `json:"msg,omitempty"`
	Status    string       `json:"status,omitempty"`
	Data      ObjectOutput `json:"objectOutput,omitempty"`
}

type CarUploadResponse struct {
	RequestId string      `json:"requestId,omitempty"`
	Code      int32       `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Status    string      `json:"status,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

// 上传数据
// func UploadData(bucketId int, dataPath string) (model.CarResponse, error) {
func UploadData(sdk *chainstoragesdk.CssClient, bucketId int, dataPath string) (model.ObjectCreateResponse, error) {
	//response := model.CarResponse{}
	response := model.ObjectCreateResponse{}

	// 数据路径为空
	if len(dataPath) == 0 {
		return response, code.ErrCarUploadFileInvalidDataPath
	}

	// 数据路径无效
	fileInfo, err := os.Stat(dataPath)
	if os.IsNotExist(err) {
		return response, code.ErrCarUploadFileInvalidDataPath
	} else if err != nil {
		return response, err
	}

	// add constant
	carVersion := 1
	fileDestination := sdk.Car.GenerateTempFileName(utils.CurrentDate()+"_", ".tmp")
	//fileDestination := GenerateTempFileName("", ".tmp")
	fmt.Printf("UploadData carVersion:%d, fileDestination:%s, dataPath:%s\n", carVersion, fileDestination, dataPath)

	// 创建Car文件
	err = sdk.Car.CreateCarFile(dataPath, fileDestination)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, code.ErrCarUploadFileCreateCarFileFail
	}
	// todo: 清除CAR文件，添加utils
	//defer func(objectSize string) {
	//	err := os.Remove(objectSize)
	//	if err != nil {
	//		fmt.Printf("Error:%+v\n", err)
	//		//logger.Errorf("file.Delete %s err: %v", objectSize, err)
	//	}
	//}(fileDestination)

	// 解析CAR文件，获取DAG信息，获取文件或目录的CID
	linkContent := ipldfmt.Link{}
	err = sdk.Car.ParseCarFile(fileDestination, &linkContent)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, code.ErrCarUploadFileParseCarFileFail
	}

	objectCid := linkContent.Cid.String()
	objectSize := int64(linkContent.Size)
	objectName := linkContent.Name

	// 设置请求参数
	carFileUploadReq := model.CarFileUploadReq{}
	carFileUploadReq.BucketId = bucketId
	carFileUploadReq.ObjectCid = objectCid
	carFileUploadReq.ObjectSize = objectSize
	carFileUploadReq.ObjectName = objectName
	carFileUploadReq.FileDestination = dataPath

	// 上传为目录的情况
	if fileInfo.IsDir() {
		// todo: add constant
		// const (
		//	ObjectTypeCodeDir   = 20000
		// )
		carFileUploadReq.ObjectTypeCode = consts.ObjectTypeCodeDir
	}

	// 计算文件sha256
	sha256, err := utils.GetFileSha256ByPath(dataPath)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, code.ErrCarUploadFileComputeCarFileHashFail
	}
	carFileUploadReq.RawSha256 = sha256

	// 使用Root CID秒传检查
	objectExistResponse, err := sdk.Object.IsExistObjectByCid(objectCid)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, code.ErrCarUploadFileReferenceObjcetFail
	}

	// CID存在，执行秒传操作
	objectExistCheck := objectExistResponse.Data
	if objectExistCheck.IsExist {
		response, err := sdk.Car.ReferenceObject(&carFileUploadReq)
		if err != nil {
			fmt.Printf("Error:%+v\n", err)
			return response, code.ErrCarUploadFileReferenceObjcetFail
		}

		return response, err
	}

	// CAR文件大小，超过分片阈值
	carFileSize := fileInfo.Size()
	carFileShardingThreshold := sdk.Config.CarFileShardingThreshold

	// 生成CAR分片文件上传
	if carFileSize > int64(carFileShardingThreshold) {
		//todo:分片上传
		response, err = UploadBigCarFile(sdk, &carFileUploadReq)
		if err != nil {
			fmt.Printf("Error:%+v\n", err)
			return response, code.ErrCarUploadFileFail
		}
	}

	// 普通上传
	response, err = sdk.Car.UploadCarFile(&carFileUploadReq)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, code.ErrCarUploadFileFail
	}

	return response, err
}

// 上传大CAR文件
func UploadBigCarFile(sdk *chainstoragesdk.CssClient, req *model.CarFileUploadReq) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 生成CAR分片文件
	shardingCarFileUploadReqs := []model.CarFileUploadReq{}
	err := sdk.Car.GenerateShardingCarFiles(req, &shardingCarFileUploadReqs)
	if err != nil {
		return response, err
	}

	// 上传CAR文件分片
	uploadingReqs := []model.CarFileUploadReq{}
	deepcopier.Copy(&shardingCarFileUploadReqs).To(&uploadingReqs)
	// 重试3次，每次间隔3秒
	maxRetries := 3
	retryDelay := time.Duration(3) * time.Second

	uploadRespList := []model.ShardingCarFileUploadResponse{}
	for i, _ := range shardingCarFileUploadReqs {
		for j := 0; j < maxRetries; j++ {
			uploadResp, err := sdk.Car.UploadShardingCarFile(&shardingCarFileUploadReqs[i])
			if err == nil {
				uploadRespList = append(uploadRespList, uploadResp)
				break
			}

			time.Sleep(retryDelay)
		}

		// 尝试maxRetries次失败
		if err != nil {
			return response, err
		}
	}

	// 确认分片上传成功
	response, err = sdk.Car.ConfirmShardingCarFiles(req)
	if err != nil {
		return response, err
	}

	return response, nil
}

// endregion CAR Upload

// region CAR Import

var carImportCmd = &cobra.Command{
	Use:     "import",
	Short:   "import",
	Long:    "import car file",
	Example: "gcscmd import  ./aaa.car cs://BUCKET",

	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Help()
		//fmt.Printf("%s %s\n", cmd.Name(), strconv.Itoa(offset))
		carImportRun(cmd, args)
	},
}

func carImportRun(cmd *cobra.Command, args []string) {
	// 桶名称
	bucketName, err := cmd.Flags().GetString("Bucket")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	// car文件标识
	carfile, err := cmd.Flags().GetString("Carfile")
	if err != nil {
		//todo: log detail error?
		//panic(err)
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)
	}

	sdk, err := chainstoragesdk.New()
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
	response, err := sdk.Car.UploadData(bucketId, carfile)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("error:%+v\n", err)
		processError("ls", err, args)

	}

	carImportRunOutput(cmd, args, response)
	//fmt.Printf("response:%+v\n", response)

}

func carImportRunOutput(cmd *cobra.Command, args []string, resp model.ObjectCreateResponse) {
	code := resp.Code
	if code != http.StatusOK {
		err := errors.Errorf("code:%d, message:&s\n", resp.Code, resp.Msg)
		processError("ls", err, args)
	}

	respData := resp.Data
	//carImportOutput := CarImportOutput{
	//	RequestId: resp.RequestId,
	//	Code:      resp.Code,
	//	Msg:       resp.Msg,
	//	Status:    resp.Status,
	//	//Count:     respData.Count,
	//	//PageIndex: respData.PageIndex,
	//	//PageSize:  respData.PageSize,
	//	List: []ObjectOutput{},
	//}

	//if len(respData.List) > 0 {
	//	for i, _ := range respData.List {
	//		objectOutput := ObjectOutput{}
	//		deepcopier.Copy(respData.List[i]).To(&objectOutput)
	//
	//		// 创建时间
	//		objectOutput.CreatedDate = objectOutput.CreatedAt.Format("2006-01-02")
	//		carImportOutput.List = append(carImportOutput.List, objectOutput)
	//	}
	//}

	//	导入 car 文件
	//	通过命令向固定桶内导入 car 文件对象
	//
	//	模版
	//
	//	gcscmd import  ./aaa.car cs://BUCKET
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
	//	当前目录
	//
	//	gcscmd import ./aaa.car cs://bbb
	//	绝对路径
	//
	//	gcscmd import /home/pz/aaa.car cs://bbb
	//	相对路径
	//
	//	gcscmd import ../pz/aaa.car cs://bbb
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
	//	报错
	//
	//Error: This is not a carfile

	templateContent := `
CID: {{.ObjectCid}}
Name: {{.ObjectName}}
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

	t, err := template.New("carImportTemplate").Parse(templateContent)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("ls", err, args)

	}

	err = t.Execute(os.Stdout, respData)
	if err != nil {
		//todo: log detail error?
		//fmt.Printf("err:%+v", resp)
		//panic(err)
		processError("ls", err, args)

	}
}

type CarImportOutput struct {
	RequestId string         `json:"requestId,omitempty"`
	Code      int32          `json:"code,omitempty"`
	Msg       string         `json:"msg,omitempty"`
	Status    string         `json:"status,omitempty"`
	Count     int            `json:"count,omitempty"`
	PageIndex int            `json:"pageIndex,omitempty"`
	PageSize  int            `json:"pageSize,omitempty"`
	List      []ObjectOutput `json:"list,omitempty"`
}

//type ObjectOutput struct {
//	Id             int       `json:"id" comment:"对象ID"`
//	BucketId       int       `json:"bucketId" comment:"桶主键"`
//	ObjectName     string    `json:"objectName" comment:"对象名称（255字限制）"`
//	ObjectTypeCode int       `json:"objectTypeCode" comment:"对象类型编码"`
//	ObjectSize     int64     `json:"objectSize" comment:"对象大小（字节）"`
//	IsMarked       int       `json:"isMarked" comment:"星标（1-已标记，0-未标记）"`
//	ObjectCid      string    `json:"objectCid" comment:"对象CID"`
//	CreatedAt      time.Time `json:"createdAt" comment:"创建时间"`
//	UpdatedAt      time.Time `json:"updatedAt" comment:"最后更新时间"`
//	CreatedDate    string    `json:"createdDate" comment:"创建日期"`
//}

// endregion CAR Import
