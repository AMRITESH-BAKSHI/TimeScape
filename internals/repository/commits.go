package repository

import (
	"TimeScape/internals/hashing"
	"TimeScape/internals/objects"

	"TimeScape/internals/storage"

)



func Commit(msg string)(string,error){

	treehash,err :=WriteTree();

	if(err!=nil){
		return "",err
	}
	parentHash :=CurrentCommit();

	Commit :=objects.CreateCommitObject(treehash,parentHash,msg);

	commitHash :=hashing.GenerateSHA1(Commit);


	err=storage.StoreObject(commitHash,Commit)

	if(err!=nil){
		return "",err;
	}

	err=UpdateCurrentCommit(commitHash);

	if(err!=nil){
		return "",err;
	}

	return commitHash,nil;




}