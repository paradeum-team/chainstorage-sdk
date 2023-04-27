package service

import (
	"bytes"
	"chainstorage-sdk/conf"
	"chainstorage-sdk/model"
	"chainstorage-sdk/utils"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-unixfsnode/data/builder"
	"github.com/ipld/go-car/v2"
	"github.com/ipld/go-car/v2/blockstore"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-multihash"
	"io"
	"os"
	"path"
	"path/filepath"
)

// 上传数据
func UploadData(dataPath string) (model.CarResponse, error) {
	response := model.CarResponse{}

	// 数据路径为空
	if len(dataPath) == 0 {
		return response, errors.New("数据路径为空")
	}

	// 数据路径无效
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return response, errors.New("数据路径无效")
	}

	// add constant
	carVersion := 1
	fileDestination := generateTempFileName(utils.CurrentDate()+"_", ".tmp")
	//fileDestination := generateTempFileName("", ".tmp")
	fmt.Printf("UploadData carVersion:%d, fileDestination:%s, dataPath:%s\n", carVersion, fileDestination, dataPath)

	// 创建Car文件
	ctx := context.Background()
	err := CreateCar(ctx, carVersion, fileDestination, dataPath)
	if err != nil {
		fmt.Printf("Error:%+v\n", err)
		return response, errors.New("创建Car文件失败")
	}

	return response, nil
}

// CreateCar creates a car
func CreateCar(ctx context.Context, carVersion int, fileDestination, dataPath string) error {
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
