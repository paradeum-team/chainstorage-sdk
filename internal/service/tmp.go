package service

import (
	"bytes"
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-unixfsnode"
	"github.com/ipld/go-car"
	"github.com/ipld/go-car/v2/blockstore"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/traversal"
	"github.com/ipld/go-ipld-prime/traversal/selector"
	"io"
	"os"

	"context"

	ipldfmt "github.com/ipfs/go-ipld-format"
	_ "github.com/ipld/go-ipld-prime/codec/cbor"
	_ "github.com/ipld/go-ipld-prime/codec/dagcbor"
	_ "github.com/ipld/go-ipld-prime/codec/dagjson"
	_ "github.com/ipld/go-ipld-prime/codec/json"
	_ "github.com/ipld/go-ipld-prime/codec/raw"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	selectorParser "github.com/ipld/go-ipld-prime/traversal/selector/parse"
)

//	func Test() {
//		// Open the CAR file
//		file, err := os.Open("example.car")
//		if err != nil {
//			fmt.Println("Failed to open CAR file:", err)
//			return
//		}
//		defer file.Close()
//
//		// Create a new CAR reader
//
//		reader, err := car.NewReader(file)
//
//		red := reader.DataReader()
//		for {
//			red.Read()
//			car.ErrOffsetImpossible
//		}
//		if err != nil {
//			fmt.Println("Failed to create CAR reader:", err)
//			return
//		}
//
//		// Create a new selector for UnixFS files and directories
//		unixFsSelector := fluent.MustBuildMap(func(m fluent.MapAssembler) {
//			m.AssembleEntry("Links", fluent.MustBuildList(func(l fluent.ListAssembler) {
//				l.AssembleValue(fluent.MustBuildMap(func(m fluent.MapAssembler) {
//					m.AssembleEntry("Name", fluent.BuildKey())
//					m.AssembleEntry("Size", fluent.BuildKey())
//					m.AssembleEntry("Type", fluent.MustBuildMap(func(m fluent.MapAssembler) {
//						m.AssembleEntry("Code", fluent.BuildKey())
//					}).Node())
//				}).Node())
//			}).Node())
//		}).Node()
//
//		// Iterate over all DAGs in the CAR file
//		for {
//			// Read the next DAG
//			dag, err := reader.Next()
//			if err == car.ErrNotFound {
//				// End of file
//				break
//			} else if err != nil {
//				fmt.Println("Failed to read DAG:", err)
//				return
//			}
//
//			// Get the UnixFS root node for the DAG
//			rootNode, err := dag.RootNode()
//			if err != nil {
//				fmt.Println("Failed to get root node:", err)
//				continue
//			}
//
//			// Traverse the UnixFS tree using the selector
//			err = rootNode.Traverse(unixFsSelector, func(prog ipld.TraversalProgress, node ipld.Node, err error) error {
//				if err != nil {
//					return err
//				}
//
//				// Extract the file or directory name and size
//				name, err := node.LookupByString("Name")
//				if err != nil {
//					return err
//				}
//				size, err := node.LookupByString("Size")
//				if err != nil {
//					return err
//				}
//
//				fmt.Println("Name:", name.AsString())
//				fmt.Println("Size:", size.AsInt())
//
//				return nil
//			}, nil)
//			if err != nil {
//				fmt.Println("Failed to traverse UnixFS tree:", err)
//				continue
//			}
//		}
//	}

// GetCarDag is a command to get a dag out of a car
func GetCarDag(carFilePath, output string) error {
	//if c.Args().Len() < 2 {
	//	return fmt.Errorf("usage: car get-dag [-s selector] <file.car> [root cid] <output file>")
	//}
	//
	//// if root cid is emitted we'll read it from the root of file.car.
	//output := c.Args().Get(1)
	var rootCid cid.Cid

	bs, err := blockstore.OpenReadOnly(carFilePath)
	if err != nil {
		return err
	}

	//if c.Args().Len() == 2 {
	roots, err := bs.Roots()
	if err != nil {
		return err
	}
	//if len(roots) != 1 {
	//	return fmt.Errorf("car file has does not have exactly one root, dag root must be specified explicitly")
	//}
	rootCid = roots[0]
	//} else {
	//	rootCid, err = cid.Parse(output)
	//	if err != nil {
	//		return err
	//	}
	//	output = c.Args().Get(2)
	//}

	strict := true //c.Bool("strict")

	// selector traversal, default to ExploreAllRecursively which only explores the DAG blocks
	// because we only care about the blocks loaded during the walk, not the nodes matched
	sel := selectorParser.CommonSelector_MatchAllRecursively
	//if c.IsSet("selector") {
	//	sel, err = selectorParser.ParseJSONSelector(c.String("selector"))
	//	if err != nil {
	//		return err
	//	}
	//}
	linkVisitOnlyOnce := true //!c.IsSet("selector") // if using a custom selector, this isn't as safe

	block, err := bs.Get(context.Background(), rootCid)
	if err != nil {
		fmt.Printf("blockstore.get(), Error:%+v\n", err)
		return err
	}

	fmt.Printf("rootCid:%s\n", rootCid.String())
	node, err := ipldfmt.Decode(block)
	if err != nil {
		fmt.Printf("blockstore.get(), Error:%+v\n", err)
		return err
	}

	fmt.Printf("\tnode:%+v\n", node)

	for _, l := range node.Links() {
		fmt.Printf("\t\tcid:%s\n", l.Cid.String())
	}

	return temp(rootCid, output, bs, strict, sel, linkVisitOnlyOnce)
	//version := 1
	//switch version {
	//case 2:
	//	return writeCarV2(context.Background(), rootCid, output, bs, strict, sel, linkVisitOnlyOnce)
	//case 1:
	//	return writeCarV1(rootCid, output, bs, strict, sel, linkVisitOnlyOnce)
	//default:
	//	return fmt.Errorf("invalid CAR version %d", version)
	//}
}

func writeCarV2(ctx context.Context, rootCid cid.Cid, output string, bs *blockstore.ReadOnly, strict bool, sel datamodel.Node, linkVisitOnlyOnce bool) error {
	_ = os.Remove(output)

	outStore, err := blockstore.OpenReadWrite(output, []cid.Cid{rootCid}, blockstore.AllowDuplicatePuts(false))
	if err != nil {
		return err
	}

	ls := cidlink.DefaultLinkSystem()
	ls.KnownReifiers = map[string]linking.NodeReifier{"unixfs": unixfsnode.Reify}
	ls.TrustedStorage = true
	ls.StorageReadOpener = func(_ linking.LinkContext, l datamodel.Link) (io.Reader, error) {
		if cl, ok := l.(cidlink.Link); ok {
			blk, err := bs.Get(ctx, cl.Cid)
			if err != nil {
				if ipldfmt.IsNotFound(err) {
					if strict {
						return nil, err
					}
					return nil, traversal.SkipMe{}
				}
				return nil, err
			}
			if err := outStore.Put(ctx, blk); err != nil {
				return nil, err
			}
			return bytes.NewBuffer(blk.RawData()), nil
		}
		return nil, fmt.Errorf("unknown link type: %T", l)
	}

	nsc := func(lnk datamodel.Link, lctx ipld.LinkContext) (datamodel.NodePrototype, error) {
		if lnk, ok := lnk.(cidlink.Link); ok && lnk.Cid.Prefix().Codec == 0x70 {
			return dagpb.Type.PBNode, nil
		}
		return basicnode.Prototype.Any, nil
	}

	rootLink := cidlink.Link{Cid: rootCid}
	ns, _ := nsc(rootLink, ipld.LinkContext{})
	rootNode, err := ls.Load(ipld.LinkContext{}, rootLink, ns)
	if err != nil {
		return err
	}

	traversalProgress := traversal.Progress{
		Cfg: &traversal.Config{
			LinkSystem:                     ls,
			LinkTargetNodePrototypeChooser: nsc,
			LinkVisitOnlyOnce:              linkVisitOnlyOnce,
		},
	}

	s, err := selector.CompileSelector(sel)
	if err != nil {
		return err
	}

	err = traversalProgress.WalkMatching(rootNode, s, func(p traversal.Progress, n datamodel.Node) error {
		lb, ok := n.(datamodel.LargeBytesNode)
		if ok {
			rs, err := lb.AsLargeBytes()
			if err == nil {
				_, err := io.Copy(io.Discard, rs)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return outStore.Finalize()
}

func writeCarV1(rootCid cid.Cid, output string, bs *blockstore.ReadOnly, strict bool, sel datamodel.Node, linkVisitOnlyOnce bool) error {
	opts := make([]car.Option, 0)
	if linkVisitOnlyOnce {
		opts = append(opts, car.TraverseLinksOnlyOnce())
	}
	sc := car.NewSelectiveCar(context.Background(), bs, []car.Dag{{Root: rootCid, Selector: sel}}, opts...)
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	return sc.Write(f)
}

func temp(rootCid cid.Cid, output string, bs *blockstore.ReadOnly, strict bool, sel datamodel.Node, linkVisitOnlyOnce bool) error {
	//opts := make([]car.Option, 0)
	//if linkVisitOnlyOnce {
	//	opts = append(opts, car.TraverseLinksOnlyOnce())
	//}
	//sc := car.NewSelectiveCar(context.Background(), bs, []car.Dag{{Root: rootCid, Selector: sel}}, opts...)
	//f, err := os.Create(output)
	//if err != nil {
	//	return err
	//}
	//defer f.Close()
	//
	//return sc.Write(f)
	return nil
}
