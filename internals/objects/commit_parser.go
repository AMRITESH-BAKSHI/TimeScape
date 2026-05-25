package objects

import ("strings")


type CommitData struct {
	TreeHash string
	ParentHash string
	MSG string
	TimeStamp string
}



func ParseCommitData(data []byte) CommitData{

	content :=ExtractBlobContent(data);


	lines :=strings.Split(string(content),"\n");

	var commit CommitData;


	for _,line :=range lines{
		if(strings.HasPrefix(line,"tree ")){
			commit.TreeHash=strings.TrimPrefix(line,"tree ");
		}else if(strings.HasPrefix(line,"parent ")){
			commit.ParentHash=strings.TrimPrefix(line,"parent ");
		}else if(strings.HasPrefix(line,"message ")){
			commit.MSG=strings.TrimPrefix(line,"message ");
		}else if(strings.HasPrefix(line,"timestamp ")){
			commit.TimeStamp=strings.TrimPrefix(line,"timestamp ");
		}
		



	}


	return commit;



}
