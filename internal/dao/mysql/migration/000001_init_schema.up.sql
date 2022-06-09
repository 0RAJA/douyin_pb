create table if not exists user # 用户信息
(
    id             bigint primary key auto_increment,
    name           varchar(40)  not null unique,
    password       varchar(128) not null,
    follow_count   bigint       not null default 0,
    follower_count bigint       not null default 0,
    constraint user_name unique index (name),
    constraint user_c1 check ( follower_count >= 0 ),
    constraint user_c2 check ( follow_count >= 0)
) default charset = utf8;

create index user_name_password on user (name, password);

create table if not exists user_followers # 用户关注
(
    id      bigint primary key auto_increment,
    from_id bigint not null,
    to_id   bigint not null,
    constraint user_followers_fk_from_id foreign key (from_id) references user (id) on delete cascade,
    constraint user_followers_fk_to_id foreign key (to_id) references user (id) on delete cascade,
    constraint user_from_to_id unique index (from_id, to_id)
) default charset = utf8;

create table if not exists videos # 视频信息
(
    id             bigint primary key auto_increment,
    user_id        bigint       not null,
    title          varchar(40)  not null,
    play_url       varchar(255) not null,
    cover_url      varchar(255) not null,
    favorite_count bigint       not null default 0,
    comment_count  bigint       not null default 0,
    created_at     timestamp             default CURRENT_TIMESTAMP not null,
    constraint videos_fk_user_id foreign key (user_id) references user (id) on delete no action,
    constraint videos_c1 check ( favorite_count >= 0 ),
    constraint videos_c2 check ( comment_count >= 0)
) default charset = utf8;

create index videos_user_id on videos (user_id);
create index videos_created_at on videos (created_at);

create table if not exists user_videos # 用户对视频的点赞关系
(
    id       bigint primary key auto_increment,
    user_id  bigint not null,
    video_id bigint not null,
    constraint user_videos_fk_user_id foreign key (user_id) references user (id) on delete cascade,
    constraint user_videos_fk_video_id foreign key (video_id) references videos (id) on delete cascade,
    constraint user_videos_user_id_video_id unique index (user_id, video_id)
) default charset = utf8;

create table if not exists comment # 用户对视频的评论
(
    id          bigint primary key auto_increment,
    user_id     bigint                              not null,
    video_id    bigint                              not null,
    content     text                                not null,
    create_date timestamp default CURRENT_TIMESTAMP not null,
    constraint comment_fk_user_id foreign key (user_id) references user (id) on delete cascade,
    constraint comment_fk_video_id foreign key (video_id) references videos (id) on delete cascade
) default charset = utf8;

create index comment_video_id on comment (video_id);
