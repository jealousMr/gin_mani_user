create table user_info(
    id bigint(20) auto_increment,
    user_id varchar(32) not null,
    nick_name varchar(128) default '',
    avatar_url varchar(256) not null,
    gender tinyint(4) unsigned not null,
    country varchar(32) default null,
    user_state tinyint(4) unsigned not null,
    create_at timestamp default CURRENT_TIMESTAMP,
    update_at timestamp default CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table user_record(
    id bigint(20) auto_increment,
    user_id varchar(32) not null,
    rule_id varchar(32) not null,
    extra text null,
    create_at timestamp default CURRENT_TIMESTAMP,
    update_at timestamp default CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE (`user_id`,`rule_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;