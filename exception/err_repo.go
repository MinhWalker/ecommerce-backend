package exception

import "errors"

var (
	//repo
	RepoNotUpdate = errors.New("Fail to update infomation")
	RepoNotFound  = errors.New("Repo not exist")
	RepoConlict   = errors.New("Repo has exist")
	RepoInsertFail = errors.New("Fail to add Repo")

	//bookmark
	BookmarkNotFound = errors.New("Bookmark not exist")
	BookmarkFail     = errors.New("Bookmark Fail")
	DelBookmarkFail  = errors.New("DelBookmark Fail")
	BookmarkConflic  = errors.New("Bookmark has exist")

	//genneral
	ErrorSql = errors.New("SQL Error")
)


