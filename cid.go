package main

import (
	"bytes"
	"context"
	"log"

	core "github.com/ipfs/go-ipfs/core"
	coreunix "github.com/ipfs/go-ipfs/core/coreunix"

	dagtest "github.com/ipfs/go-ipfs/merkledag/test"
	mfs "github.com/ipfs/go-ipfs/mfs"
	ft "github.com/ipfs/go-ipfs/unixfs"
)

func main() {
	// ipfsnode
	ctx := context.TODO()
	node, err := core.NewNode(ctx, &core.BuildCfg{})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer node.Close()

	data := []byte("hello")
	// coreunix.Add
	ipfs_cid, err := coreunix.Add(node, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(ipfs_cid)

	// fileAdder
	fileAdder, err := coreunix.NewAdder(ctx, node.Pinning, node.Blockstore, node.DAG)
	if err != nil {
		return
	}

	md := dagtest.Mock()
	mr, err := mfs.NewRoot(ctx, md, ft.EmptyDirNode(), nil)
	if err != nil {
		return
	}
	fileAdder.SetMfsRoot(mr)

	root, err := fileAdder.Myadd(bytes.NewReader(data))
	if err != nil {
		return
	}
	log.Println(root.Cid().String())

	log.Println("finished")
}
