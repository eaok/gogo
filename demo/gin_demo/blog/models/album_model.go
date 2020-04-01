package models

import "blog/dao"

type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	CreateTime int64 `db:"create_time"`
}

func AddAlbum(album *Album) (int64, error) {
	sqlStr := "insert into album(filepath,filename,status,create_time)values(?,?,?,?)"
	return dao.ModifyDB(sqlStr, album.Filepath, album.Filename, album.Status, album.CreateTime)
}

func QueryAlbum() (dest []*Album, err error) {
	sqlStr := "select id,filepath,filename,status,create_time from album"
	err = dao.QueryRows(&dest, sqlStr)
	return
}
