-- User accounts
create table if not exists blabber.user (
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
create table if not exists blabber.user_follow (
    follower_id uuid not null,
    followed_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (follower_id, followed_id),
    foreign key (follower_id) references blabber.user(id) on delete cascade,
    foreign key (followed_id) references blabber.user(id) on delete cascade
);

-- Blocks between users
create table if not exists blabber.user_block (
    blocker_id uuid not null,
    blocked_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (blocker_id, blocked_id),
    foreign key (blocker_id) references blabber.user(id) on delete cascade,
    foreign key (blocked_id) references blabber.user(id) on delete cascade
);

-- Mutes between users
create table if not exists blabber.user_mute (
    muter_id uuid not null,
    muted_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (muter_id, muted_id),
    foreign key (muter_id) references blabber.user(id) on delete cascade,
    foreign key (muted_id) references blabber.user(id) on delete cascade
);