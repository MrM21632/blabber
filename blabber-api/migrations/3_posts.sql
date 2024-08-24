-- Posts
create table if not exists blabber.post (
    id uuid not null primary key,  -- default gen_random_uuid()
    user_id uuid not null references blabber.user(id) on delete cascade,
    parent_id uuid references blabber.post(id) on delete set null,
    contents varchar(200) not null,
    created_at timestamp with time zone not null,
    likes int not null default 0,
    reposts int not null default 0
);

-- Tags on posts
create table if not exists blabber.post_tag (
    post_id uuid not null,
    tag varchar(64) not null,

    primary key (post_id, tag),
    foreign key (post_id) references blabber.post(id) on delete cascade
);

-- Likes on posts
create table if not exists blabber.post_like (
    post_id uuid not null,
    user_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (post_id, user_id),
    foreign key (post_id) references blabber.post(id) on delete cascade,
    foreign key (user_id) references blabber.user(id) on delete cascade
);

-- Reposts
create table if not exists blabber.repost (
    post_id uuid not null,
    user_id uuid not null,
    created_at timestamp with time zone not null,

    primary key (post_id, user_id),
    foreign key (post_id) references blabber.post(id) on delete cascade,
    foreign key (user_id) references blabber.user(id) on delete cascade
);