package EditorFileModel

import (
	"sync"
	"OutsourcingPlatform/ospf/infrastructure"
)

var onceGetUploadFile sync.Once

type UploadFileModel struct {
	*infrastructure.Infrastructure
	FilePath string
}

var uploadFileModel *UploadFileModel

func GetUploadFileModel (i *infrastructure.Infrastructure,filepath string) *UploadFileModel{
	onceGetUploadFile.Do(func() {
		uploadFileModel = &UploadFileModel{
			Infrastructure:i,
			FilePath:filepath,
		}
	})
	return uploadFileModel
}

