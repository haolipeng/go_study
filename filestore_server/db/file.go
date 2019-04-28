package db

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go_study/filestore_server/db/mysql"
	"os"
)

func InnerFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	dbConn := mysql.DBConn()
	if dbConn == nil {
		fmt.Println("db connection is nil")
		os.Exit(1)
	}

	stmt, err := dbConn.Prepare("insert ignore tbl_file (file_sha1,file_name,file_size,file_addr,status) " +
		"values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to prepare statement err:", err.Error())
		return false
	}

	res, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	rows, err := res.RowsAffected()
	if err == nil {
		if rows <= 0 {
			fmt.Printf("Exec function not affected,file with hash :%s have been upload before")
		}
		return true
	}

	return false
}

type TableFile struct {
	FileHash sql.NullString
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//GetFileMeta:从数据库中获取文件元信息
func InnerGetFileMetaFromDB(filehash string) (*TableFile, error) {
	dbConn := mysql.DBConn()
	tabFile := TableFile{}
	stmt, err := dbConn.Prepare("select file_sha1,file_name,file_size,file_addr from tbl_file where file_sha1=? and status=1")
	if err != nil {
		return &tabFile, errors.New("prepare failed")
	}

	err = stmt.QueryRow(filehash).Scan(&tabFile.FileHash, &tabFile.FileName, &tabFile.FileSize, &tabFile.FileAddr)
	if err != nil {
		return &tabFile, errors.New("QueryRow failed")
	}

	return &tabFile, nil
}
