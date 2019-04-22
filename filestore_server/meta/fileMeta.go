package meta

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
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetaMap[fileSha1]
}

//根据sha1删除文件元信息
func RemoveFileMeta(fileSha1 string) {
	//fix me,是否线程安全？
	delete(fileMetaMap, fileSha1)
}
