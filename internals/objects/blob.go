package objects;

import "fmt"


func CreateBlobObject(content []byte) []byte{

	header := fmt.Sprintf("blob %d \x00",len(content));

	return append([]byte(header),content...);
}


