package repository

import ( 
	"os"
	"TimeScape/internals/objects"
	"TimeScape/internals/storage"

)



func Checkout(commitHash string)error{

	commitData,err :=storage.ReadObject(commitHash);

	if(err!=nil){
		return err;
	}

	commit :=objects.ParseCommitData(commitData);
	TreeContent ,err:=storage.ReadObject(commit.TreeHash);
	
	if(err!=nil){
		return err;
	}

	entries :=objects.ParserTreeObject(TreeContent);

	for _,entry :=range entries {
		blobdata,err :=storage.ReadObject(entry.Hash);
		
	if(err!=nil){
		return err;
	}
	content :=objects.ExtractBlobContent(blobdata);

	err =os.WriteFile(entry.FileName,content,0644)
	if(err!=nil){
		return err;
	}
	

	}
return nil;

}

