package index

import (
	"os"
	
	"strings"
)


type IndexEntry struct {
	FileName string
	Hash string
}


func ReadIndex()([]IndexEntry,error){

	indexPath :=".TimeScape/index";

	data,err :=os.ReadFile(indexPath);

	if(err!=nil){
		return nil,err;
	}


	lines :=strings.Split(string(data),"\n");
	var entries []IndexEntry;


	for _,line := range lines{
		if(line==""){
			continue
		}

		parts :=strings.Split(line," ");
		entry:=IndexEntry{ FileName:parts[0],Hash:parts[1],}
		entries=append(entries, entry);
		}	
		return entries,nil;

}