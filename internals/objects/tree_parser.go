package objects

import (
	
	"strings"
)


type TreeEntry struct {
	Type string
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
			Type: parts[0],
			FileName :parts[1],
			Hash: parts[2],
		} 

		entries =append(entries,entry);
		}

		return entries;
}