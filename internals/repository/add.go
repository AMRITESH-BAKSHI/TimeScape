package repository

import (
	"TimeScape/internals/hashing"
	"TimeScape/internals/index"
	"fmt"
	"os"
	"TimeScape/internals/objects"
	

	"TimeScape/internals/storage"
	"TimeScape/utils"
)


func AddFile(fileName string) error{

_,err :=os.Stat(fileName)
	if(err!=nil){
		fmt.Println("file does not exists")
		return  err
	}
	
	
	content,err :=utils.ReadFile(fileName);
	
	if(err!=nil){
		fmt.Println("Error reading file")
		return err
	}
	
	blob := objects.CreateBlobObject(content);
	
	hash :=hashing.GenerateSHA1(blob);
	err=storage.StoreObject(hash,blob)
	if(err!=nil){
		return err;
	}
	
	err= index.AddToIndex(fileName,hash);
	if(err!=nil){
		return err;
	}
	
return nil

}