package objects

import (
	"fmt"
	
	
)



func CreateTreeObject(content string) []byte{
	
	
	header := fmt.Sprintf("tree %d\x00",len(content));


	return append([]byte(header),[]byte(content)...);

}