package service

import (
	"bytes"
	client "chainstorage-sdk/base/tokenclient"
	"chainstorage-sdk/conf"
	"chainstorage-sdk/model"
	"chainstorage-sdk/utils"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alanshaw/go-carbites"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	ipldfmt "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-unixfsnode/data/builder"
	"github.com/ipld/go-car/v2"
	"github.com/ipld/go-car/v2/blockstore"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// 上传数据
func UploadData(bucketId int, dataPath string) (model.CarResponse, error) {
	response := model.CarResponse{}

	// 数据路径为空
	if len(dataPath) == 0 {
		return response, errors.New("数据路径为空")
	}

	// 数据路径无效
	fileInfo, err := os.Stat(dataPath)
	if os.IsNotExist(err) {
		return response, errors.New("数据路径无效")
	} else if err != nil {
		return response, err
	}

	// add constant
	carVersion := 1
	fileDestination := generateTempFileName(utils.CurrentDate()+"_", ".tmp")
	//fileDestination := generateTempFileName("", ".tmp")
	fmt.Printf("UploadData carVersion:%d, fileDestination:%s, dataPath:%s\n", carVersion, fileDestination, dataPath)

	// 创建Car文件
	ctx := context.Background()
	err = createCar(ctx, carVersion, fileDestination, dataPath)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, errors.New("创建Car文件失败")
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
	err = parseCarDag(fileDestination, &linkContent)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, errors.New("解析CAR文件失败")
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
		carFileUploadReq.ObjectTypeCode = 20000
	}

	// 计算文件sha256
	sha256, err := utils.GetFileSha256ByPath(dataPath)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, errors.New("hash计算失败")
	}
	carFileUploadReq.RawSha256 = sha256

	// 使用Root CID秒传检查
	objectExistResponse, err := IsExistObjectByCid(objectCid)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, errors.New("CID秒传检查失败")
	}

	// CID存在，执行秒传操作
	objectExistCheck := objectExistResponse.Data
	if objectExistCheck.IsExist {
		//todo:
		ReferenceObject(&carFileUploadReq)
	}

	// CAR文件大小，超过分片阈值
	carFileSize := fileInfo.Size()
	carFileShardingThreshold := conf.Config.CarFileShardingThreshold

	// 生成CAR分片文件上传
	if carFileSize > int64(carFileShardingThreshold) {
		//todo:分片上传
	}

	// 普通上传
	UploadCarFile(&carFileUploadReq)

	return response, nil
}

// CreateCar creates a car
func createCar(ctx context.Context, carVersion int, fileDestination, dataPath string) error {
	var err error

	// make a cid with the right length that we eventually will patch with the root.
	hasher, err := multihash.GetHasher(multihash.SHA2_256)
	if err != nil {
		return err
	}
	digest := hasher.Sum([]byte{})
	hash, err := multihash.Encode(digest, multihash.SHA2_256)
	if err != nil {
		return err
	}
	proxyRoot := cid.NewCidV1(uint64(multicodec.DagPb), hash)

	options := []car.Option{}
	switch carVersion {
	case 1:
		options = []car.Option{blockstore.WriteAsCarV1(true)}
	case 2:
		// already the default
	default:
		return fmt.Errorf("invalid CAR version %d", carVersion)
	}

	cdest, err := blockstore.OpenReadWrite(fileDestination, []cid.Cid{proxyRoot}, options...)
	if err != nil {
		return err
	}

	// Write the unixfs blocks into the store.
	root, err := writeFiles(context.Background(), cdest, dataPath)
	if err != nil {
		return err
	}

	if err := cdest.Finalize(); err != nil {
		return err
	}
	// re-open/finalize with the final root.
	return car.ReplaceRootsInFile(fileDestination, []cid.Cid{root})
}

func writeFiles(ctx context.Context, bs *blockstore.ReadWrite, paths ...string) (cid.Cid, error) {
	ls := cidlink.DefaultLinkSystem()
	ls.TrustedStorage = true
	ls.StorageReadOpener = func(_ ipld.LinkContext, l ipld.Link) (io.Reader, error) {
		cl, ok := l.(cidlink.Link)
		if !ok {
			return nil, fmt.Errorf("not a cidlink")
		}
		blk, err := bs.Get(ctx, cl.Cid)
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(blk.RawData()), nil
	}
	ls.StorageWriteOpener = func(_ ipld.LinkContext) (io.Writer, ipld.BlockWriteCommitter, error) {
		buf := bytes.NewBuffer(nil)
		return buf, func(l ipld.Link) error {
			cl, ok := l.(cidlink.Link)
			if !ok {
				return fmt.Errorf("not a cidlink")
			}
			blk, err := blocks.NewBlockWithCid(buf.Bytes(), cl.Cid)
			if err != nil {
				return err
			}
			bs.Put(ctx, blk)
			return nil
		}, nil
	}

	topLevel := make([]dagpb.PBLink, 0, len(paths))
	for _, p := range paths {
		l, size, err := builder.BuildUnixFSRecursive(p, &ls)
		if err != nil {
			return cid.Undef, err
		}
		name := path.Base(p)
		entry, err := builder.BuildUnixFSDirectoryEntry(name, int64(size), l)
		if err != nil {
			return cid.Undef, err
		}
		topLevel = append(topLevel, entry)
	}

	// make a directory for the file(s).

	root, _, err := builder.BuildUnixFSDirectory(topLevel, &ls)
	if err != nil {
		return cid.Undef, nil
	}
	rcl, ok := root.(cidlink.Link)
	if !ok {
		return cid.Undef, fmt.Errorf("could not interpret %s", root)
	}

	return rcl.Cid, nil
}

// TempFileName generates a temporary filename for use in testing or whatever
func generateTempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)

	carFileGenerationPath := conf.Config.CarFileGenerationPath
	if _, err := os.Stat(carFileGenerationPath); os.IsNotExist(err) {
		_ = os.MkdirAll(carFileGenerationPath, os.ModePerm)
	}

	//return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
	return filepath.Join(carFileGenerationPath, prefix+hex.EncodeToString(randBytes)+suffix)
}

func generateFileName(prefix, suffix string) string {
	carFileGenerationPath := conf.Config.CarFileGenerationPath
	if _, err := os.Stat(carFileGenerationPath); os.IsNotExist(err) {
		_ = os.MkdirAll(carFileGenerationPath, os.ModePerm)
	}

	//return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
	return filepath.Join(carFileGenerationPath, prefix+suffix)
}

// parse a dag from a car file
func parseCarDag(carFilePath string, linkContent *ipldfmt.Link) error {
	bs, err := blockstore.OpenReadOnly(carFilePath)
	if err != nil {
		return err
	}
	defer bs.Close()

	roots, err := bs.Roots()
	if err != nil {
		return err
	}

	if len(roots) != 1 {
		//return fmt.Errorf("car file has does not have exactly one root, dag root must be specified explicitly")
		return fmt.Errorf("car文件根级别仅支持单个文件或者目录")
	}

	rootCid := roots[0]
	block, err := bs.Get(context.Background(), rootCid)
	if err != nil {
		fmt.Printf("parseCarDag:blockstore.get(), Error:%+v\n", err)
		return err
	}

	node, err := ipldfmt.Decode(block)
	if err != nil {
		fmt.Printf("parseCarDag:blockstore.get(), Error:%+v\n", err)
		return err
	}

	links := node.Links()
	if len(links) == 0 {
		return fmt.Errorf("there aren't any IPFS Merkle DAG Link between Nodes")
	}

	link := links[0]
	if link == nil {
		return fmt.Errorf("there aren't any IPFS Merkle DAG Link between Nodes")
	}

	linkContent.Cid = link.Cid
	linkContent.Name = link.Name
	linkContent.Size = link.Size

	return nil
}

func TempParseCarDag(carFilePath string, linkContent *ipldfmt.Link) error {
	return parseCarDag(carFilePath, linkContent)
}

func SliceBigCarFile(carFilePath string) error {
	bigCarFile, err := os.Open(carFilePath)
	if err != nil {
		return err
	}
	defer bigCarFile.Close()

	targetSize := conf.Config.CarFileShardingThreshold * 10 //1024 * 1024     // 1MiB chunks
	strategy := carbites.Treewalk                           // also carbites.Treewalk
	spltr, _ := carbites.Split(bigCarFile, targetSize, strategy)

	var i int
	for {
		car, err := spltr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		b, _ := io.ReadAll(car)
		filename := fmt.Sprintf("chunk-%d.car", i)
		fileDestination := generateFileName(utils.CurrentDate()+"_", filename)
		//ioutil.WriteFile(fmt.Sprintf("chunk-%d.car", i), b, 0644)
		os.WriteFile(fileDestination, b, 0644)
		i++
	}

	return nil
}

func generateShardingCarFiles(carFilePath string) error {
	bigCarFile, err := os.Open(carFilePath)
	if err != nil {
		return err
	}
	defer bigCarFile.Close()

	targetSize := conf.Config.CarFileShardingThreshold * 10 //1024 * 1024     // 1MiB chunks
	strategy := carbites.Treewalk                           // also carbites.Treewalk
	spltr, _ := carbites.Split(bigCarFile, targetSize, strategy)

	shardingCarFileDestinationList := []string{}
	i := 1
	for {
		car, err := spltr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		b, _ := io.ReadAll(car)
		filename := fmt.Sprintf("_chunk.c%d", i)
		//fileDestination := generateFileName(utils.CurrentDate()+"_", filename)
		fileDestination := strings.Replace(carFilePath, filepath.Ext(carFilePath), filename, 1)
		shardingCarFileDestinationList = append(shardingCarFileDestinationList, fileDestination)

		//ioutil.WriteFile(fmt.Sprintf("chunk-%d.car", i), b, 0644)
		err = os.WriteFile(fileDestination, b, 0644)
		if err != nil {
			panic(err)
		}
		i++
	}

	return nil
}

// 引用对象
func ReferenceObject(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 请求Url
	apiBaseAddress := conf.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/upload/car/reference"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	bucketId := req.BucketId
	rawSha256 := req.RawSha256
	objectCid := req.ObjectCid
	objectName := req.ObjectName
	objectTypeCode := req.ObjectTypeCode
	//fileDestination := req.FileDestination

	//params := map[string]string{
	//	"bucketId":  strconv.Itoa(bucketId),
	//	"rawSha256": rawSha256,
	//	"objectCid": objectCid,
	//}
	params := map[string]interface{}{
		"bucketId":       bucketId,
		"rawSha256":      rawSha256,
		"objectCid":      objectCid,
		"objectName":     objectName,
		"objectTypeCode": objectTypeCode,
	}

	// API调用
	httpStatus, body, err := client.RestyPost(apiUrl, params)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:ReferenceObject:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:ReferenceObject:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:ReferenceObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	fmt.Printf("response:%+v", response)
	return response, nil
}

// 上传CAR文件
func UploadCarFile(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 请求Url
	apiBaseAddress := conf.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/upload/car/file"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	bucketId := req.BucketId
	rawSha256 := req.RawSha256
	objectCid := req.ObjectCid
	objectName := req.ObjectName
	fileDestination := req.FileDestination

	params := map[string]string{
		"bucketId":  strconv.Itoa(bucketId),
		"rawSha256": rawSha256,
		"objectCid": objectCid,
	}
	//params := map[string]interface{}{
	//	"bucketId":  bucketId,
	//	"rawSha256": rawSha256,
	//	"objectCid": objectCid,
	//}

	// API调用
	httpStatus, body, err := client.RestyPostForm(objectName, fileDestination, params, apiUrl)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:UploadCarFile:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:UploadCarFile:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:UploadCarFile:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	fmt.Printf("response:%+v", response)
	return response, nil
}

// 上传CAR文件
func UploadShardingCarFile(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 请求Url
	apiBaseAddress := conf.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/upload/car/shard"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	bucketId := req.BucketId
	rawSha256 := req.RawSha256
	objectCid := req.ObjectCid
	objectName := req.ObjectName
	fileDestination := req.FileDestination
	shardingSha256 := req.ShardingSha256
	shardingNo := req.ShardingNo

	params := map[string]string{
		"bucketId":       strconv.Itoa(bucketId),
		"rawSha256":      rawSha256,
		"objectCid":      objectCid,
		"shardingSha256": shardingSha256,
		"shardingNo":     strconv.Itoa(shardingNo),
	}
	//params := map[string]interface{}{
	//	"bucketId":  bucketId,
	//	"rawSha256": rawSha256,
	//	"objectCid": objectCid,
	//}

	// API调用
	httpStatus, body, err := client.RestyPostForm(objectName, fileDestination, params, apiUrl)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:UploadShardingCarFile:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:UploadShardingCarFile:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:UploadShardingCarFile:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	fmt.Printf("response:%+v", response)
	return response, nil
}
