package model

import (
	"OutsourcingPlatform/ospf/infrastructure"
	"log"
	"sync"
	"fmt"
	"strconv"
	"math/rand"
)

type Model struct {
	*infrastructure.Infrastructure
}

var onceModel sync.Once
var model *Model

func GetModel(i *infrastructure.Infrastructure) *Model {
	onceModel.Do(func() {
		model = &Model{
			Infrastructure: i,
		}
	})
	return model
}

func (m *Model) InitAllTable() {

	// 创建 syslog 表
	_, err := m.DB.Exec(`create table if not exists syslog(
		id SERIAL NOT NULL,
		log jsonb NOT NULL,
		PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Panicln("create table OutsourcingPlatform.syslog failed: " + err.Error())
	}

	 	//创建   user	表
	 	_, err = m.DB.Exec(`
create table if not exists competitors (
   project_id           INT8                 not null,
   user_id              INT8                 not null,
   compete_id           TIMESTAMP            null,
   constraint PK_COMPETITORS primary key (project_id, user_id)
);

create table if not exists evaluate (
   evaluate_id          SERIAL not null,
   estimator_id         INT8                 null,
   estimatored_id       INT8                 null,
   evaluate_time        TIMESTAMP            null,
   evaluate_detail      TEXT                 null,
   project_ssociated_id INT8                 null,
   is_privacy           BOOL                 null,
   constraint PK_EVALUATE primary key (evaluate_id)
);


create table if not exists field_labels (
   label_id             SERIAL not null,
   creator_id           INT8                 null,
   label_name           VARCHAR(32)          null,
   label_type           INT4                 null,
   constraint PK_FIELD_LABELS primary key (label_id)
);


create table if not exists project (
   project_id           SERIAL not null,
   publisher_id         INT8                 not null,
   contractor_id        INT8                 null,
   project_name         VARCHAR(128)         not null,
   project_publish_time TIMESTAMP            null,
   project_end_time     TIMESTAMP            null,
   project_alter_time   TIMESTAMP            null,
   project_circle       INT8                 null,
   project_money        FLOAT8               null,
   project_detail       TEXT                 null,
   project_status       INT4                 null,
   constraint PK_PROJECT primary key (project_id)
);


create table if not exists teammate (
   project_id           INT8                 not null,
   user_id              INT8                 not null,
   identity_id          INT4                 null,
   constraint PK_TEAMMATE primary key (project_id, user_id)
);


create table if not exists users (
   user_id              SERIAL not null,
   invited_id           INT8                 null,
   user_name            VARCHAR(32)          not null,
   openid               VARCHAR(128)         null,
   head_img             TEXT                 null,
   user_password        TEXT                 null,
   user_tel             VARCHAR(32)          null,
   user_sex             VARCHAR(16)          null,
   user_idcard          VARCHAR(32)          null,
   user_bank            VARCHAR(32)          null,
   permissions          INT4                 null,
   regist_time          TIMESTAMP            null,
   personal_update_time TIMESTAMP            null,
   permissiongs_update_time TIMESTAMP            null,
   prove_data           TEXT                 null,
   expertise_field      TEXT                 null,
   wechat               TEXT                 null,
   address              TEXT                 null,
   level                INT4                 null,
   user_state           INT4                 null,
   invitation_qrcode_token_normal TEXT                 default '',
invitation_qrcode_token_witness text default '',
   constraint PK_USERS primary key (user_id)
);
	 `)
	 	if err != nil {
	 		log.Panicln("create table OutsourcingPlatform.syslog failed: " + err.Error())
	 	}

}

func (m *Model) InsertDataToAllTable() {

/*
	insert into users (user_name,openid)
	values( '上帝','mockopenid3');
*/



	insert_users_sql1:=`
insert into users (user_name,openid)
	values( '上帝','mockopenid3');
`
	_,err:= m.DB.Exec(insert_users_sql1)
	if err!=nil {
		fmt.Println("insert_users_sql1:error",err)

	}


	insert_users_sql := `

insert into users(user_name,openid,head_img,user_password,user_tel,user_sex,user_idcard,user_bank,permissions,
				regist_time,personal_update_time,permissiongs_update_time,prove_data,expertise_field,wechat,address,level,user_state)
				values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)`
	sex := [2]string{"男","女"}
	for i:= 50;i<20000;i++{
		_,err:= m.DB.Exec(insert_users_sql,"李雷"+strconv.Itoa(i),rand.Int(),"head_img_name"+strconv.Itoa(i),"1234656abc","15635457843",
			sex[(i%2)],"440402155632012457","6222020903001483077",15,"2018-01-01","2018-01-01","2018-01-01","there are some data who can prove me","angular,golang,java",
				"T1120491948","广东省 番禺区 广州大学",1,0)
		if err!=nil {
			fmt.Println("insert_users_sql:error",err)
			break
		}
	}

}
