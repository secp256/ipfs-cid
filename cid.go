package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"

	// core "github.com/ipfs/go-ipfs/core"
	balanced "github.com/ipfs/go-ipfs/importer/balanced"
	ihelper "github.com/ipfs/go-ipfs/importer/helpers"
	chunker "gx/ipfs/QmWo8jYc19ppG7YoTsrr2kEtLRbARTJho5oNXFTR6B7Peq/go-ipfs-chunker"
)

func get_cid(data []byte) {

	// ipfsnode
	// ctx := context.TODO()
	// node, err := core.NewNode(ctx, &core.BuildCfg{
	// 	NilRepo: true,
	// })
	// if err != nil {
	// 	os.Exit(2)
	// 	return
	// }
	// defer node.Close()

	chnk, err := chunker.FromString(bytes.NewReader(data), "")
	if err != nil {
		os.Exit(3)
		return
	}

	params := ihelper.DagBuilderParams{}
	root, err := balanced.Layout(params.New(chnk))
	if err != nil {
		os.Exit(4)
		return
	}

	fmt.Printf("{\"cid\":\"%s\"}", root.Cid().String())
}

func str_to_bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func main() {
	var fname, raw_string string
	flag.StringVar(&fname, "f", "", "file to be added")
	flag.StringVar(&raw_string, "s", "", "string to be added")
	flag.Parse()

	if raw_string == "" && fname == "" {
		get_cid(str_to_bytes(""))
		return
	}

	if len(raw_string) > 0 {
		data := str_to_bytes(raw_string)
		// data := []byte(raw_string)
		get_cid(data)
	} else if len(fname) > 0 {
		data, _ := ioutil.ReadFile(fname)
		// fmt.Println(len(data))
		get_cid(data)
	}
}
