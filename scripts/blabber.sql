-- Cleanup if necessary prior to running the rest of the script
drop schema if exists blabber cascade;

-- (Re)create the schema
create schema blabber;

-- User accounts
create table blabber.user (
    id uuid not null primary key,  -- default gen_random_uuid()
    username varchar(80) not null,
    user_handle varchar(20) not null unique,
    user_bio varchar(200),
    email text not null unique,
    password_hash text not null,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    followers int not null default 0,
    follows int not null default 0
);

-- Followings between users
create table blabber.user_follow (
    follower_id uuid not null,
    followed_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (follower_id, followed_id),
    foreign key (follower_id) references blabber.user(id) on delete cascade,
    foreign key (followed_id) references blabber.user(id) on delete cascade
);

-- Blocks between users
create table blabber.user_block (
    blocker_id uuid not null,
    blocked_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (blocker_id, blocked_id),
    foreign key (blocker_id) references blabber.user(id) on delete cascade,
    foreign key (blocked_id) references blabber.user(id) on delete cascade
);

-- Mutes between users
create table blabber.user_mute (
    muter_id uuid not null,
    muted_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (muter_id, muted_id),
    foreign key (muter_id) references blabber.user(id) on delete cascade,
    foreign key (muted_id) references blabber.user(id) on delete cascade
);

-- Posts
create table blabber.post (
    id uuid not null primary key,  -- default gen_random_uuid()
    user_id uuid not null references blabber.user(id) on delete cascade,
    parent_id uuid references blabber.post(id) on delete set null,
    contents varchar(200) not null,
    created_at timestamp with time zone not null,
    likes int not null default 0,
    reposts int not null default 0
);

-- Tags on posts
create table blabber.post_tag (
    post_id uuid not null,
    tag varchar(64) not null,

    primary key (post_id, tag),
    foreign key (post_id) references blabber.post(id) on delete cascade
);

-- Likes on posts
create table blabber.post_like (
    post_id uuid not null,
    user_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (post_id, user_id),
    foreign key (post_id) references blabber.post(id) on delete cascade,
    foreign key (user_id) references blabber.user(id) on delete cascade
);

-- Reposts
create table blabber.repost (
    post_id uuid not null,
    user_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (post_id, user_id),
    foreign key (post_id) references blabber.post(id) on delete cascade,
    foreign key (user_id) references blabber.user(id) on delete cascade
);
