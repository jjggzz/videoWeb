create database video;

use video;

create table video (
    id bigint primary key auto_increment comment '主键ID',
    access_key varchar(64) not null comment '业务key',
    create_time datetime not null comment '创建时间',
    update_time datetime not null comment '修改时间',
    delete_status int not null default 0 comment '状态 0:未删除 1:已删除',
    customer_id bigint not null comment '客户主键id',
    video_name varchar(64) not null comment '视频名称',
    video_introduction varchar(255) not null comment '视频简介',
    video_cover_path varchar(255) not null comment '视频封面路径',
    video_source_path varchar(255) not null comment '视频资源路径',
    video_size bigint not null comment '视频大小',
    video_thumbs_count bigint not null default 0 comment '视频点赞数',
    video_play_count bigint not null default 0 comment '视频播放数',
    video_comment_count bigint not null default 0 comment '视频评论数'
);

create index idx_video_create_time on video(create_time);
create unique index idx_video_access_key on video(access_key);
create index idx_video_customer_id on video(customer_id);
