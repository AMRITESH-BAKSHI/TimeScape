package objects

import("bytes")

func ExtractBlobContent(data []byte) []byte{
	index :=bytes.IndexByte(data,0);

	return data[index+1:];
}