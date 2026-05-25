package objects

import (
	"fmt"
	
	"time"
	
	
)



func CreateCommitObject(treehash string,parenthash string,msg string)[]byte{

	content :=fmt.Sprintf("tree %s\nparent %s\nmessage %s\ntimestamp %d \n ",treehash,parenthash,msg,time.Now().Unix());

	
	
	header := fmt.Sprintf("commit %d\x00",len(content));


	return append([]byte(header),[]byte(content)...);



}