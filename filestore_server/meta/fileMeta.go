package meta

import "go_study/filestore_server/db"

//FileMeta:文件元信息
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetaMap map[string]FileMeta

func init() {
	fileMetaMap = make(map[string]FileMeta)
}

//新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetaMap[fmeta.FileSha1] = fmeta
}

//获取文件元信息
func GetFileMeta(fileSha1 string) (FileMeta, bool) {
	fMeta, ok := fileMetaMap[fileSha1]
	return fMeta, ok
}

//根据sha1删除文件元信息
func RemoveFileMeta(fileSha1 string) {
	//fix me,是否线程安全？
	delete(fileMetaMap, fileSha1)
}

////////////////////////////////在数据库中对文件进行增删改查操作///////////////////////////////////////////////
//UpdateFileMeta2DB:更新文件元信息到mysql中
func UpdateFileMeta2DB(fmeta FileMeta) bool {
	return db.InnerFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

//从数据库获取文件元信息
func GetFileMetaFromDb(fileSha1 string) (FileMeta, bool) {
	tbFile, err := db.InnerGetFileMetaFromDB(fileSha1)
	fmeta := FileMeta{}

	if err != nil {
		return fmeta, false
	}

	fmeta.FileSha1 = tbFile.FileHash.String
	fmeta.FileName = tbFile.FileName.String
	fmeta.Location = tbFile.FileAddr.String
	fmeta.FileSize = tbFile.FileSize.Int64

	return fmeta, true
}
