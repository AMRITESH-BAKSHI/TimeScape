package objects

import (
	
	"strings"
)


type TreeEntry struct {
	FileName string
	Hash string



}


func ParserTreeObject(data[]byte)[]TreeEntry{

	content := ExtractBlobContent(data);

	lines := strings.Split(string(content),"\n");
	var entries []TreeEntry;

	for _,line :=range lines {
		if(line==""){
			continue
		}

		parts :=strings.Split(line," ");

		entry :=TreeEntry{
			FileName :parts[0],
			Hash: parts[1],
		} 

		entries =append(entries,entry);
		}

		return entries;
}