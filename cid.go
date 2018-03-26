package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	core "github.com/ipfs/go-ipfs/core"
	coreunix "github.com/ipfs/go-ipfs/core/coreunix"
	dagtest "github.com/ipfs/go-ipfs/merkledag/test"
	mfs "github.com/ipfs/go-ipfs/mfs"
	ft "github.com/ipfs/go-ipfs/unixfs"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(1)
		return
	}
	data := []byte(os.Args[1])

	// ipfsnode
	ctx := context.TODO()
	node, err := core.NewNode(ctx, &core.BuildCfg{})
	if err != nil {
		os.Exit(2)
		return
	}
	defer node.Close()

	// fileAdder
	fileAdder, err := coreunix.NewAdder(ctx, node.Pinning, node.Blockstore, node.DAG)
	if err != nil {
		os.Exit(3)
		return
	}

	md := dagtest.Mock()
	mr, err := mfs.NewRoot(ctx, md, ft.EmptyDirNode(), nil)
	if err != nil {
		os.Exit(4)
		return
	}
	fileAdder.SetMfsRoot(mr)

	root, err := fileAdder.Myadd(bytes.NewReader(data))
	if err != nil {
		os.Exit(5)
		return
	}

	ipfs_cid := root.Cid().String()
	fmt.Printf("{\"cid\":\"%s\"}", ipfs_cid)
}
