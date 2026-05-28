package repository

import (
	"TimeScape/internals/objects"
	"TimeScape/internals/storage"
	
	"os"
	"path/filepath"
)



func Checkout(commitHash string)error{

	commitData,err :=storage.ReadObject(commitHash);

	if(err!=nil){
		return err;
	}

	commit :=objects.ParseCommitData(commitData);
	// TreeContent ,err:=storage.ReadObject(commit.TreeHash);
	err =RestoreTree(commit.TreeHash,".");
	
	if(err!=nil){
		return err;
	}
	// entries :=objects.ParserTreeObject(TreeContent);

	// for _,entry :=range entries {
	// 	blobdata,err :=storage.ReadObject(entry.Hash);
		
	
	// content :=objects.ExtractBlobContent(blobdata);

	// err =os.WriteFile(entry.FileName,content,0644)
	
	// }
	err =UpdateHead(commitHash);
	
	if(err!=nil){
		return err;
	}
	

return nil;

}



func RestoreTree(treeHash string,currentPath string)error{

		treeData, err :=storage.ReadObject(treeHash);
		if(err!=nil){
			return err;
		}

		entries :=objects.ParserTreeObject(treeData);

		for _,entry :=range entries{
			fullPath :=filepath.Join(currentPath,entry.FileName);

			if(entry.Type == "blob"){
				blobData,err :=storage.ReadObject(entry.Hash);
				if(err!=nil){
					return err;
				}
				content :=objects.ExtractBlobContent(blobData);
				
				err =os.WriteFile(fullPath,content,0644);

				if(err!=nil){
					return err;
				}



			}else if(entry.Type=="tree"){
				err :=os.MkdirAll(fullPath,0755);
				if(err!=nil){
					return err;
				}
				err =RestoreTree(entry.Hash,fullPath);
				if(err!=nil){
					return err;
				}	
				
			}

			}
			return nil;
			
}