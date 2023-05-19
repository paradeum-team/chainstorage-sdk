package chainstoragesdk

import (
	"bytes"
	"chainstoragesdk/code"
	"chainstoragesdk/consts"
	"chainstoragesdk/model"
	"chainstoragesdk/utils"
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
	"github.com/kataras/golog"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	"github.com/ulule/deepcopier"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Car struct {
	Config *Configuration
	Client *RestyClient
	logger *golog.Logger
}

// 创建CAR文件
func (c *Car) CreateCarFile(dataPath string, fileDestination string) error {
	ctx := context.Background()
	carVersion := 1
	return createCar(ctx, carVersion, fileDestination, dataPath)
}

// 创建CAR文件分片
func (c *Car) SplitCarFile(carFilePath string, chunkedFileDestinations *[]string) error {
	// CAR file chunking setting
	// todo:
	targetSize := c.Config.CarFileShardingThreshold * 10 //1024 * 1024     // 1MiB chunks
	strategy := carbites.Treewalk

	return chunkCarFile(carFilePath, targetSize, strategy, chunkedFileDestinations)
}

// 引用对象
func (c *Car) ReferenceObject(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 请求Url
	apiBaseAddress := c.Config.ChainStorageApiBaseAddress
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
	httpStatus, body, err := c.Client.RestyPost(apiUrl, params)
	if err != nil {
		//utils.LogError(fmt.Sprintf("API:ReferenceObject:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))
		c.logger.Errorf(fmt.Sprintf("API:ReferenceObject:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		c.logger.Errorf(fmt.Sprintf("API:ReferenceObject:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.logger.Errorf(fmt.Sprintf("API:ReferenceObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	//fmt.Printf("response:%+v", response)
	return response, nil
}

// 上传CAR文件
func (c *Car) UploadCarFile(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 请求Url
	apiBaseAddress := c.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/upload/car/file"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	bucketId := req.BucketId
	rawSha256 := req.RawSha256
	objectCid := req.ObjectCid
	objectName := req.ObjectName
	fileDestination := req.FileDestination
	carFileCid := req.CarFileCid
	objectSize := strconv.FormatInt(req.ObjectSize, 10)

	params := map[string]string{
		"bucketId":   strconv.Itoa(bucketId),
		"rawSha256":  rawSha256,
		"objectCid":  objectCid,
		"carFileCid": carFileCid,
		"objectName": objectName,
		"objectSize": objectSize,
	}
	//params := map[string]interface{}{
	//	"bucketId":  bucketId,
	//	"rawSha256": rawSha256,
	//	"objectCid": objectCid,
	//}

	// API调用
	httpStatus, body, err := c.Client.RestyPostForm(objectName, fileDestination, params, apiUrl)
	if err != nil {
		c.logger.Errorf(fmt.Sprintf("API:UploadCarFile:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		c.logger.Errorf(fmt.Sprintf("API:UploadCarFile:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.logger.Errorf(fmt.Sprintf("API:UploadCarFile:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 上传CAR文件分片
func (c *Car) UploadShardingCarFile(req *model.CarFileUploadReq) (model.ShardingCarFileUploadResponse, error) {
	response := model.ShardingCarFileUploadResponse{}

	// 请求Url
	apiBaseAddress := c.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/upload/car/shard"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	bucketId := req.BucketId
	rawSha256 := req.RawSha256
	objectCid := req.ObjectCid
	objectName := req.ObjectName
	fileDestination := req.FileDestination
	shardingSha256 := req.ShardingSha256
	shardingNo := req.ShardingNo
	carFileCid := req.CarFileCid
	objectSize := strconv.FormatInt(req.ObjectSize, 10)

	params := map[string]string{
		"bucketId":       strconv.Itoa(bucketId),
		"rawSha256":      rawSha256,
		"objectCid":      objectCid,
		"shardingSha256": shardingSha256,
		"shardingNo":     strconv.Itoa(shardingNo),
		"carFileCid":     carFileCid,
		"objectSize":     objectSize,
	}
	//params := map[string]interface{}{
	//	"bucketId":  bucketId,
	//	"rawSha256": rawSha256,
	//	"objectCid": objectCid,
	//}

	// API调用
	httpStatus, body, err := c.Client.RestyPostForm(objectName, fileDestination, params, apiUrl)
	if err != nil {
		c.logger.Errorf(fmt.Sprintf("API:UploadShardingCarFile:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		c.logger.Errorf(fmt.Sprintf("API:UploadShardingCarFile:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.logger.Errorf(fmt.Sprintf("API:UploadShardingCarFile:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 校验上传CAR文件分片结果
func (c *Car) VerifyShardingCarFiles(req *model.CarFileUploadReq) (model.ShardingCarFilesVerifyResponse, error) {
	response := model.ShardingCarFilesVerifyResponse{}

	// 请求Url
	apiBaseAddress := c.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/upload/car/verify"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	bucketId := req.BucketId
	rawSha256 := req.RawSha256
	objectCid := req.ObjectCid
	objectName := req.ObjectName
	objectTypeCode := req.ObjectTypeCode
	shardingAmount := req.ShardingAmount
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
		"shardingAmount": shardingAmount,
	}

	// API调用
	httpStatus, body, err := c.Client.RestyPost(apiUrl, params)
	if err != nil {
		c.logger.Errorf(fmt.Sprintf("API:VerifyShardingCarFiles:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		c.logger.Errorf(fmt.Sprintf("API:VerifyShardingCarFiles:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.logger.Errorf(fmt.Sprintf("API:VerifyShardingCarFiles:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	//fmt.Printf("response:%+v", response)
	return response, nil
}

// 确认上传CAR文件分片结果
func (c *Car) ConfirmShardingCarFiles(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 请求Url
	apiBaseAddress := c.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/upload/car/confirm"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	bucketId := req.BucketId
	rawSha256 := req.RawSha256
	objectCid := req.ObjectCid
	objectName := req.ObjectName
	objectTypeCode := req.ObjectTypeCode
	shardingAmount := req.ShardingAmount
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
		"shardingAmount": shardingAmount,
	}

	// API调用
	httpStatus, body, err := c.Client.RestyPost(apiUrl, params)
	if err != nil {
		c.logger.Errorf(fmt.Sprintf("API:ConfirmShardingCarFiles:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		c.logger.Errorf(fmt.Sprintf("API:ConfirmShardingCarFiles:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.logger.Errorf(fmt.Sprintf("API:ConfirmShardingCarFiles:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// region CAR file

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

// 生成CAR分片文件
func chunkCarFile(carFilePath string, targetSize int, strategy carbites.Strategy, chunkedFileDestinations *[]string) error {
	carFile, err := os.Open(carFilePath)
	if err != nil {
		return err
	}
	defer carFile.Close()

	// create CAR splitter
	splitter, err := carbites.Split(carFile, targetSize, strategy)
	if err != nil {
		return err
	}

	index := 1
	for {
		carPart, err := splitter.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			//fmt.Printf("Error:%+v\n", err)
			return err
		}

		bytes, err := io.ReadAll(carPart)
		if err != nil {
			//fmt.Printf("Error:%+v\n", err)
			return err
		}

		// set chunking file destination
		filename := fmt.Sprintf("_chunk.c%d", index)
		chunkedFileDestination := strings.Replace(carFilePath, filepath.Ext(carFilePath), filename, 1)
		*chunkedFileDestinations = append(*chunkedFileDestinations, chunkedFileDestination)

		// create chunking file
		err = os.WriteFile(chunkedFileDestination, bytes, 0644)
		if err != nil {
			//fmt.Printf("Error:%+v\n", err)
			return err
		}

		index++
	}

	return nil
}

// parse a dag from a car file
func parseCarDag(carFilePath string, rootLink *model.RootLink) error {
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
	//fmt.Printf("rootCid:%s\n", rootCid.String())

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

	rootLink.RootCid = rootCid
	rootLink.Cid = link.Cid
	rootLink.Name = link.Name
	rootLink.Size = link.Size
	//fmt.Printf("linkCid:%s\n", link.Cid.String())

	return nil
}

// endregion CAR file

// TempFileName generates a temporary filename for use in testing or whatever
func (c *Car) GenerateTempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)

	carFileGenerationPath := c.Config.CarFileGenerationPath
	if _, err := os.Stat(carFileGenerationPath); os.IsNotExist(err) {
		_ = os.MkdirAll(carFileGenerationPath, os.ModePerm)
	}

	//return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
	return filepath.Join(carFileGenerationPath, prefix+hex.EncodeToString(randBytes)+suffix)
}

func (c *Car) generateFileName(prefix, suffix string) string {
	carFileGenerationPath := c.Config.CarFileGenerationPath
	if _, err := os.Stat(carFileGenerationPath); os.IsNotExist(err) {
		_ = os.MkdirAll(carFileGenerationPath, os.ModePerm)
	}

	//return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
	return filepath.Join(carFileGenerationPath, prefix+suffix)
}

func (c *Car) ParseCarFile(carFilePath string, rootLink *model.RootLink) error {
	return parseCarDag(carFilePath, rootLink)
}

func (c *Car) SliceBigCarFile(carFilePath string) error {
	bigCarFile, err := os.Open(carFilePath)
	if err != nil {
		return err
	}
	defer bigCarFile.Close()

	targetSize := c.Config.CarFileShardingThreshold * 10 //1024 * 1024     // 1MiB chunks
	strategy := carbites.Treewalk                        // also carbites.Treewalk
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
		fileDestination := c.generateFileName(utils.CurrentDate()+"_", filename)
		//ioutil.WriteFile(fmt.Sprintf("chunk-%d.car", i), b, 0644)
		os.WriteFile(fileDestination, b, 0644)
		i++
	}

	return nil
}

// 生成CAR分片文件
func (c *Car) GenerateShardingCarFiles(req *model.CarFileUploadReq, shardingCarFileUploadReqs *[]model.CarFileUploadReq) error {
	fileDestination := req.FileDestination

	bigCarFile, err := os.Open(fileDestination)
	if err != nil {
		return err
	}
	defer bigCarFile.Close()

	// CAR文件分片设置
	targetSize := c.Config.CarFileShardingThreshold //1024 * 1024     // 1MiB chunks
	strategy := carbites.Treewalk                   // also carbites.Treewalk
	spltr, _ := carbites.Split(bigCarFile, targetSize, strategy)

	//shardingCarFileDestinationList := []string{}
	shardingNo := 1

	for {
		car, err := spltr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			//panic(err)

			fmt.Printf("Error:%+v\n", err)
			return err
		}

		bytes, err := io.ReadAll(car)
		if err != nil {
			//panic(err)

			fmt.Printf("Error:%+v\n", err)
			return err
		}

		// 设置文件名称
		filename := fmt.Sprintf("_chunk.c%d", shardingNo)
		//shardingFileDestination := generateFileName(utils.CurrentDate()+"_", filename)
		shardingFileDestination := strings.Replace(fileDestination, filepath.Ext(fileDestination), filename, 1)
		//shardingCarFileDestinationList = append(shardingCarFileDestinationList, shardingFileDestination)

		chunkSize := int64(len(bytes))

		// 生成分片文件
		//ioutil.WriteFile(fmt.Sprintf("chunk-%d.car", shardingNo), bytes, 0644)
		err = os.WriteFile(shardingFileDestination, bytes, 0644)
		if err != nil {
			//panic(err)

			fmt.Printf("Error:%+v\n", err)
			return err
		}

		// 计算分片文件sha256
		shardingSha256, err := utils.GetFileSha256ByPath(shardingFileDestination)
		if err != nil {
			//panic(err)

			fmt.Printf("Error:%+v\n", err)
			return err
		}
		//carFileUploadReq.RawSha256 = shardingSha256

		shardingNo++

		// 设置分片请求对象
		shardingCarFileUploadReq := model.CarFileUploadReq{}
		deepcopier.Copy(req).To(&shardingCarFileUploadReq)
		shardingCarFileUploadReq.FileDestination = shardingFileDestination
		shardingCarFileUploadReq.ShardingSha256 = shardingSha256
		shardingCarFileUploadReq.ShardingNo = shardingNo

		//// todo: remove it
		//rootLink := model.RootLink{}
		//parseCarDag(shardingFileDestination, &rootLink)
		//rootCid := rootLink.RootCid.String()
		//size := int64(rootLink.Size)
		//shardingCarFileUploadReq.CarFileCid = rootCid
		shardingCarFileUploadReq.ObjectSize = chunkSize

		*shardingCarFileUploadReqs = append(*shardingCarFileUploadReqs, shardingCarFileUploadReq)
	}

	// 分片失败
	shardingAmount := len(*shardingCarFileUploadReqs)
	if shardingAmount == 0 {
		// todo: add constant
		fmt.Printf("Error:%+v\n", err)
		//return err
		return code.ErrCarUploadFileChunkCarFileFail
	}

	req.ShardingAmount = shardingAmount

	return nil
}

// 上传数据
// func UploadData(bucketId int, dataPath string) (model.CarResponse, error) {
func (c *Car) UploadData(bucketId int, dataPath string) (model.ObjectCreateResponse, error) {
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
	fileDestination := c.GenerateTempFileName(utils.CurrentDate()+"_", ".tmp")
	//fileDestination := GenerateTempFileName("", ".tmp")
	fmt.Printf("UploadData carVersion:%d, fileDestination:%s, dataPath:%s\n", carVersion, fileDestination, dataPath)

	// 创建Car文件
	ctx := context.Background()
	err = createCar(ctx, carVersion, fileDestination, dataPath)
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
	rootLink := model.RootLink{}
	err = parseCarDag(fileDestination, &rootLink)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, code.ErrCarUploadFileParseCarFileFail
	}

	rootCid := rootLink.RootCid.String()
	objectCid := rootLink.Cid.String()
	objectSize := int64(rootLink.Size)
	objectName := rootLink.Name

	// 设置请求参数
	carFileUploadReq := model.CarFileUploadReq{}
	carFileUploadReq.BucketId = bucketId
	carFileUploadReq.ObjectCid = objectCid
	carFileUploadReq.ObjectSize = objectSize
	carFileUploadReq.ObjectName = objectName
	carFileUploadReq.FileDestination = dataPath
	carFileUploadReq.CarFileCid = rootCid

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
	objectService := Object{}
	objectExistResponse, err := objectService.IsExistObjectByCid(objectCid)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, code.ErrCarUploadFileReferenceObjcetFail
	}

	// CID存在，执行秒传操作
	objectExistCheck := objectExistResponse.Data
	if objectExistCheck.IsExist {
		response, err := c.ReferenceObject(&carFileUploadReq)
		if err != nil {
			fmt.Printf("Error:%+v\n", err)
			return response, code.ErrCarUploadFileReferenceObjcetFail
		}

		return response, err
	}

	// CAR文件大小，超过分片阈值
	carFileSize := fileInfo.Size()
	carFileShardingThreshold := c.Config.CarFileShardingThreshold

	// 生成CAR分片文件上传
	if carFileSize > int64(carFileShardingThreshold) {
		//todo:分片上传
		response, err = c.UploadBigCarFile(&carFileUploadReq)
		if err != nil {
			fmt.Printf("Error:%+v\n", err)
			return response, code.ErrCarUploadFileFail
		}
	}

	// 普通上传
	response, err = c.UploadCarFile(&carFileUploadReq)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, code.ErrCarUploadFileFail
	}

	return response, err
}

// 上传大CAR文件
func (c *Car) UploadBigCarFile(req *model.CarFileUploadReq) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 生成CAR分片文件
	shardingCarFileUploadReqs := []model.CarFileUploadReq{}
	err := c.GenerateShardingCarFiles(req, &shardingCarFileUploadReqs)
	if err != nil {
		return response, err
	}

	// 上传CAR文件分片
	uploadingReqs := []model.CarFileUploadReq{}
	deepcopier.Copy(&shardingCarFileUploadReqs).To(&uploadingReqs)

	//for {
	//
	//}
	for i, _ := range shardingCarFileUploadReqs {
		c.UploadShardingCarFile(&shardingCarFileUploadReqs[i])
	}

	return response, nil
}
