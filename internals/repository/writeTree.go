package repository

import (
	"TimeScape/internals/hashing"
	"TimeScape/internals/index"
	"TimeScape/internals/objects"
	"TimeScape/internals/storage"
	
)




func WriteTree()(string,error){

	entries ,err  :=index.ReadIndex();
	if(err!=nil){
		return "",err;
	}

	tree :=objects.CreateTreeObject(entries);

	hash :=hashing.GenerateSHA1(tree);

	err =storage.StoreObject(hash,tree)

	if err!=nil{
		return "",err	
	}


	return hash,nil;
}