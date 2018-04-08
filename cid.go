package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	core "github.com/ipfs/go-ipfs/core"
	coreunix "github.com/ipfs/go-ipfs/core/coreunix"
	dagtest "github.com/ipfs/go-ipfs/merkledag/test"
	mfs "github.com/ipfs/go-ipfs/mfs"
	ft "github.com/ipfs/go-ipfs/unixfs"
)

func get_cid(data []byte) {

	// ipfsnode
	ctx := context.TODO()
	node, err := core.NewNode(ctx, &core.BuildCfg{
		NilRepo: true,
	})
	if err != nil {
		os.Exit(2)
		return
	}
	// defer node.Close()

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

func main() {
	var fname, raw_string string
	flag.StringVar(&fname, "f", "", "file to be added")
	flag.StringVar(&raw_string, "s", "", "string to be added")
	flag.Parse()

	// fmt.Println("fname:", fname)
	// fmt.Println("raw_string:", raw_string)

	if raw_string == "" && fname == "" {
		os.Exit(1)
		return
	}

	if len(raw_string) > 0 {
		data := []byte(raw_string)
		// fmt.Println("data:", data)
		get_cid(data)
	} else if len(fname) > 0 {
		data, _ := ioutil.ReadFile(fname)
		// fmt.Println("data:", data)
		get_cid(data)
	}
}
