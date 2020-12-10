create database customer;

use customer;

create table customer (
  id bigint primary key auto_increment comment '主键ID',
  access_key varchar(64) not null comment '业务key',
  create_time datetime not null comment '创建时间',
  update_time datetime not null comment '修改时间',
  delete_status int not null comment '状态 0:已删除 1:未删除',
  phone char(11) not null comment '电话号码',
  username varchar(32) not null comment '用户名',
  password varchar(64) not null comment '密码',
  email varchar(64) not null comment '邮箱',
  nickname varchar(32) not null comment '昵称',
  status int not null  comment '账号状态 0:冻结 1:启动'
)comment '用户表';

create index idx_customer_create_time on customer(create_time);
create unique index idx_customer_access_key on customer(access_key);
create unique index idx_customer_phone on customer(phone);
create unique index idx_customer_username on customer(username);