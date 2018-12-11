package main

func main() {
	blob, err := PrepareBlob()
	if err != nil {
		panic(err)
	}

	chunks := map[string][]byte{
		"AAAA": []byte("name: hello\n"),
		"BBBB": []byte("  - foo: world\n"),
		"CCCC": []byte("  - bar: world\n"),
	}

	for id, content := range chunks {
		if err := blob.PutBlock(id, content, nil); err != nil {
			panic(err)
		}
	}

	if err := PutBlockList(blob, "AAAA", "BBBB", "CCCC"); err != nil {
		panic(err)
	}
	if err := PrintBlob(blob); err != nil {
		panic(err)
	}

	if err := PutBlockList(blob, "AAAA", "CCCC"); err != nil {
		panic(err)
	}
	if err := PrintBlob(blob); err != nil {
		panic(err)
	}
}
