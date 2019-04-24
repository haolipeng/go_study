create database fileserver;
use fileserver;
CREATE TABLE tbl_file (
	id int(11) not null AUTO_INCREMENT,
	file_sha1 char(40) not null default '' comment 'file hash',
    file_name varchar(256) not null default '' comment 'file name',
    file_size bigint(20) default '0' comment 'file size',
    file_addr varchar(1024) not null default '' comment 'file location',
    create_at datetime default now() comment 'create time',
    update_at datetime default now() on update current_timestamp() comment 'update time',
    status int(11) not null default '0' comment 'status',
    ext1 int(11) default '0' comment 'ext1',
    ext2 text comment 'ext2',
    primary key (id),
    unique key idx_file_hash (file_sha1),
    key idx_status (status)
)engine=innodb default charset=utf8;