package projectOperateModel

import (
	"database/sql"
)

//有关项目的人将插入到teammate表中 （1 承包者 2 发布者 3普通成员 4路人）
func InsertIntoTeammate(ProjectId, UserId, identity int, DB *sql.DB) error {
	SqlInsert := `INSERT INTO teammate(project_id, user_id, identity_id) VALUES ($1,$2,$3)`
	stmt, err := DB.Prepare(SqlInsert)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ProjectId, UserId, identity)
	if err != nil {
		return err
	}
	return nil
}
