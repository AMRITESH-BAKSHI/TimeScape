package objects

import (
	"fmt"
	"strings"
	"TimeScape/internals/index"
)



func CreateTreeObject(entries []index.IndexEntry) []byte{
	
	
	var lines []string

	for _,entry :=range entries{
		line:=fmt.Sprintf( "%s %s",entry.FileName,entry.Hash);

		lines=append(lines, line);
	} 
	treeContent :=strings.Join(lines,"\n");
	
	header := fmt.Sprintf("tree %d\x00",len(treeContent));


	return append([]byte(header),[]byte(treeContent)...);

}