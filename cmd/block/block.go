package main

func testBlockBlob() error {
	blob, err := PrepareBlob()
	if err != nil {
		return err
	}

	chunks := map[string][]byte{
		"AAAA": []byte("name: hello\n"),
		"BBBB": []byte("  - foo: world\n"),
		"CCCC": []byte("  - bar: world\n"),
	}

	for id, content := range chunks {
		if err := blob.PutBlock(id, content, nil); err != nil {
			return err
		}
	}

	if err := PutAndPrintBlockList(blob, "AAAA", "BBBB", "CCCC"); err != nil {
		return err
	}
	if err := PutAndPrintBlockList(blob, "AAAA", "CCCC", "BBBB"); err != nil {
		return err
	}
	if err := PutAndPrintBlockList(blob, "AAAA", "CCCC"); err != nil {
		return err
	}

	IDs, err := GetBlockList(blob)
	if err != nil {
		return err
	}
	if err := blob.PutBlock("DDDD", []byte("version: 0.1.0\n"), nil); err != nil {
		return err
	}
	newIDs := append([]string(nil), IDs[0], "DDDD")
	newIDs = append(newIDs, IDs[1:]...)
	if err := PutAndPrintBlockList(blob, newIDs...); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := testBlockBlob(); err != nil {
		panic(err)
	}
}
