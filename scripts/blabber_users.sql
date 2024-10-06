-- Clear the tables
delete from blabber.user;
delete from blabber.user_follow;
delete from blabber.user_block;
delete from blabber.user_mute;

-- Insert user account records
insert into blabber.user
    (id, username, user_handle, user_bio, email, password_hash, created_at, updated_at, followers, follows)
values
    (
        '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid,
        'testuser01', 'testuser01',
        'this is a test user',
        'testuser01@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        24, 24
    ),
    (
        '01925f6f-b9d1-7d1c-b89b-9791083af810'::uuid,
        'testuser02', 'testuser02',
        'this is a test user',
        'testuser02@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7486-a4b6-5c7be7d07696'::uuid,
        'testuser03', 'testuser03',
        'this is a test user',
        'testuser03@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7ae1-8a34-eefc726633bb'::uuid,
        'testuser04', 'testuser04',
        'this is a test user',
        'testuser04@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-79c5-a81d-0f03b2f5fb2e'::uuid,
        'testuser05', 'testuser05',
        'this is a test user',
        'testuser05@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-786e-b3dc-8fabf8cfae15'::uuid,
        'testuser06', 'testuser06',
        'this is a test user',
        'testuser06@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-74a5-bca8-9b5dc7953e09'::uuid,
        'testuser07', 'testuser07',
        'this is a test user',
        'testuser07@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-70f3-88a8-1d1f09c7d06b'::uuid,
        'testuser08', 'testuser08',
        'this is a test user',
        'testuser08@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7660-8d04-4258fda596dd'::uuid,
        'testuser09', 'testuser09',
        'this is a test user',
        'testuser09@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-70ab-923b-70ca533c9a04'::uuid,
        'testuser10', 'testuser10',
        'this is a test user',
        'testuser10@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-78d2-a667-ca72e0df19ae'::uuid,
        'testuser11', 'testuser11',
        'this is a test user',
        'testuser11@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-75f2-bbf9-cda51ac07651'::uuid,
        'testuser12', 'testuser12',
        'this is a test user',
        'testuser12@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7a62-bf11-4ea07f543b78'::uuid,
        'testuser13', 'testuser13',
        'this is a test user',
        'testuser13@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7936-aea5-f1ef9d9cfe60'::uuid,
        'testuser14', 'testuser14',
        'this is a test user',
        'testuser14@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7608-b4c1-978f71f706d6'::uuid,
        'testuser15', 'testuser15',
        'this is a test user',
        'testuser15@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-78ba-b874-b1620271cd2b'::uuid,
        'testuser16', 'testuser16',
        'this is a test user',
        'testuser16@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7ee2-910f-bfc06b10c955'::uuid,
        'testuser17', 'testuser17',
        'this is a test user',
        'testuser17@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-776b-a49b-259f0bdc2d78'::uuid,
        'testuser18', 'testuser18',
        'this is a test user',
        'testuser18@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7481-ae6d-53e24abda1a7'::uuid,
        'testuser19', 'testuser19',
        'this is a test user',
        'testuser19@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7eb3-a795-3174135fe8a2'::uuid,
        'testuser20', 'testuser20',
        'this is a test user',
        'testuser20@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7ce9-a57a-08468da08e7e'::uuid,
        'testuser21', 'testuser21',
        'this is a test user',
        'testuser21@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-79b7-a4da-843bd928986b'::uuid,
        'testuser22', 'testuser22',
        'this is a test user',
        'testuser22@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-785a-af81-2cd2dd032963'::uuid,
        'testuser23', 'testuser23',
        'this is a test user',
        'testuser23@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-7fd4-bf0e-53b8d1abfa96'::uuid,
        'testuser24', 'testuser24',
        'this is a test user',
        'testuser24@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925f6f-b9d1-78f4-83aa-9b6b9f70d892'::uuid,
        'testuser25', 'testuser25',
        'this is a test user',
        'testuser25@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        1, 1
    ),
    (
        '01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid,
        'testuser26', 'testuser26',
        'this is a test user',
        'testuser26@gmail.com',
        '$argon2id$v=19$m=65536,t=3,p=2$TjhISDFpZXBMVUw1bmkxUw$h8BqJV5UyjVtX0QxzDAdJed0ftGQl4cw/n/CiQkWybk',
        current_timestamp, current_timestamp,
        0, 0
    );


-- Insert user following records
insert into blabber.user_follow
    (follower_id, followed_id, created_at)
values
    -- testuser01 follows every other account; tests building feeds from many different users
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7d1c-b89b-9791083af810'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7486-a4b6-5c7be7d07696'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7ae1-8a34-eefc726633bb'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-79c5-a81d-0f03b2f5fb2e'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-786e-b3dc-8fabf8cfae15'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-74a5-bca8-9b5dc7953e09'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-70f3-88a8-1d1f09c7d06b'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7660-8d04-4258fda596dd'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-70ab-923b-70ca533c9a04'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-78d2-a667-ca72e0df19ae'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-75f2-bbf9-cda51ac07651'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7a62-bf11-4ea07f543b78'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7936-aea5-f1ef9d9cfe60'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7608-b4c1-978f71f706d6'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-78ba-b874-b1620271cd2b'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7ee2-910f-bfc06b10c955'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-776b-a49b-259f0bdc2d78'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7481-ae6d-53e24abda1a7'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7eb3-a795-3174135fe8a2'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7ce9-a57a-08468da08e7e'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-79b7-a4da-843bd928986b'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-785a-af81-2cd2dd032963'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-7fd4-bf0e-53b8d1abfa96'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, '01925f6f-b9d1-78f4-83aa-9b6b9f70d892'::uuid, current_timestamp),
    -- Every other user follows testuser01
    ('01925f6f-b9d1-7d1c-b89b-9791083af810'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7486-a4b6-5c7be7d07696'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7ae1-8a34-eefc726633bb'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-79c5-a81d-0f03b2f5fb2e'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-786e-b3dc-8fabf8cfae15'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-74a5-bca8-9b5dc7953e09'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-70f3-88a8-1d1f09c7d06b'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7660-8d04-4258fda596dd'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-70ab-923b-70ca533c9a04'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-78d2-a667-ca72e0df19ae'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-75f2-bbf9-cda51ac07651'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7a62-bf11-4ea07f543b78'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7936-aea5-f1ef9d9cfe60'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7608-b4c1-978f71f706d6'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-78ba-b874-b1620271cd2b'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7ee2-910f-bfc06b10c955'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-776b-a49b-259f0bdc2d78'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7481-ae6d-53e24abda1a7'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7eb3-a795-3174135fe8a2'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7ce9-a57a-08468da08e7e'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-79b7-a4da-843bd928986b'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-785a-af81-2cd2dd032963'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-7fd4-bf0e-53b8d1abfa96'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925f6f-b9d1-78f4-83aa-9b6b9f70d892'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp);


-- Insert user block records
insert into blabber.user_block
    (blocker_id, blocked_id, created_at)
values
    -- testuser26 has blocked testuser01
    -- Blocking: User's messages are not visible in feeds. However, they can still appear in threads; in this case,
    -- they will initially be hidden.
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp);


-- Insert user mute records
insert into blabber.user_mute
    (muter_id, muted_id, created_at)
values
    -- testuser26 has muted testuser11 through testuser20
    -- Muting: User's messages are not visible across the site.
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-78d2-a667-ca72e0df19ae'::uuid, current_timestamp),
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-75f2-bbf9-cda51ac07651'::uuid, current_timestamp),
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-7a62-bf11-4ea07f543b78'::uuid, current_timestamp),
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-7936-aea5-f1ef9d9cfe60'::uuid, current_timestamp),
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-7608-b4c1-978f71f706d6'::uuid, current_timestamp),
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-78ba-b874-b1620271cd2b'::uuid, current_timestamp),
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-7ee2-910f-bfc06b10c955'::uuid, current_timestamp),
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-776b-a49b-259f0bdc2d78'::uuid, current_timestamp),
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-7481-ae6d-53e24abda1a7'::uuid, current_timestamp),
    ('01925fd7-8179-7ed2-bfb9-3a7479b9b6ba'::uuid, '01925f6f-b9d1-7eb3-a795-3174135fe8a2'::uuid, current_timestamp);


select id, username, user_handle, followers, follows from blabber.user;

select
    u.id,
    u.username,
    count(distinct uf1.follower_id) as follower_count,
    count(distinct uf2.followed_id) as followed_count
from blabber.user u
left join blabber.user_follow uf1
    on u.id = uf1.followed_id
left join blabber.user_follow uf2
    on u.id = uf2.follower_id
group by u.id;
