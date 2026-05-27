package repository

import (
	"os"
	"fmt"
	"path/filePath"
	"TimeScape/internals/hashing"
	"TimeScape/internals/storage"
	"TimeScape/internals/objects"
)



func CreateTreeFromDirectory(path string)(string,error){

	entries ,err := os.ReadDir(path);
	if(err!=nil){
		return "",err;
	}


	var treeContent string;
	
	for _ ,entry :=range entries{
		name :=entry.Name();

		if (name==".TimeScape" || name==".git"){
			continue;
		}

		fullPath := filepath.Join(path,name);

		if(entry.IsDir()){
			subtreeHash,err :=CreateTreeFromDirectory(fullPath);

			if(err!=nil){
				return "",err;
			}


			treeContent +=fmt.Sprintf("tree %s %s\n",name,subtreeHash);


		}else{
			data,err :=os.ReadFile(fullPath);

			if (err!=nil){

				return "",err;
			}

			blob :=objects.CreateBlobObject(data);
			blobHash :=hashing.GenerateSHA1(blob);
			err=storage.StoreObject(blobHash,blob)

			if(err!=nil){
				return "",err;
			}

			treeContent +=fmt.Sprintf("blob %s %s\n",name,blobHash);
		}

		
		
	}
	tree :=objects.CreateTreeObject(treeContent);
	treeHash :=hashing.GenerateSHA1(tree);

	err =storage.StoreObject(treeHash,tree);
	if(err!=nil){
		return "",err;
	}


	return treeHash,nil
}