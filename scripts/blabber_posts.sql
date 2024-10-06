delete from blabber.post;
delete from blabber.post_tag;
delete from blabber.post_like;
delete from blabber.post_repost;


-- Insert post records
insert into blabber.post
    (id, user_id, parent_id, contents, created_at, likes, reposts)
values
    (
        '01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid,
        '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, null,
        'This is a test post', current_timestamp, 25, 0
    ),
    (
        '01925fff-fd06-75aa-8fc3-2814af59bfb1'::uuid,
        '01925f6f-b9d1-7d1c-b89b-9791083af810'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-77e4-b72c-fc23c4231069'::uuid,
        '01925f6f-b9d1-7486-a4b6-5c7be7d07696'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7288-9b7b-963b7858092f'::uuid,
        '01925f6f-b9d1-7ae1-8a34-eefc726633bb'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-70bf-b9b6-743eee142ca1'::uuid,
        '01925f6f-b9d1-79c5-a81d-0f03b2f5fb2e'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7b1a-b185-bf2c406ee201'::uuid,
        '01925f6f-b9d1-786e-b3dc-8fabf8cfae15'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7435-892d-49ed8a142afa'::uuid,
        '01925f6f-b9d1-74a5-bca8-9b5dc7953e09'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-742f-ad45-f3206a7ae19b'::uuid,
        '01925f6f-b9d1-70f3-88a8-1d1f09c7d06b'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7483-9ef5-7ac80538518a'::uuid,
        '01925f6f-b9d1-7660-8d04-4258fda596dd'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-736e-97b6-0064e63f20a4'::uuid,
        '01925f6f-b9d1-70ab-923b-70ca533c9a04'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7e7b-b8cc-55eb2fa7c699'::uuid,
        '01925f6f-b9d1-78d2-a667-ca72e0df19ae'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7952-b621-999d7c671f2d'::uuid,
        '01925f6f-b9d1-75f2-bbf9-cda51ac07651'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-799e-8768-e3db0e1f22af'::uuid,
        '01925f6f-b9d1-7a62-bf11-4ea07f543b78'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-76b7-be15-df67b7674ca5'::uuid,
        '01925f6f-b9d1-7936-aea5-f1ef9d9cfe60'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7bd2-9bbf-7c7b42d154ec'::uuid,
        '01925f6f-b9d1-7608-b4c1-978f71f706d6'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7a61-a371-4fc28a6009e6'::uuid,
        '01925f6f-b9d1-78ba-b874-b1620271cd2b'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7f3b-ba38-1963d19a0548'::uuid,
        '01925f6f-b9d1-7ee2-910f-bfc06b10c955'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7ab5-b2f0-fcb2dae91430'::uuid,
        '01925f6f-b9d1-776b-a49b-259f0bdc2d78'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7821-aa21-c9896a8155d2'::uuid,
        '01925f6f-b9d1-7481-ae6d-53e24abda1a7'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-71f6-826f-a68c2156687f'::uuid,
        '01925f6f-b9d1-7eb3-a795-3174135fe8a2'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-715e-ab7e-054bea3cd171'::uuid,
        '01925f6f-b9d1-7ce9-a57a-08468da08e7e'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-7f69-b360-7780747f71dc'::uuid,
        '01925f6f-b9d1-79b7-a4da-843bd928986b'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-70f2-901e-d1905ab8bd09'::uuid,
        '01925f6f-b9d1-785a-af81-2cd2dd032963'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-79d1-abad-ad5eb02d9177'::uuid,
        '01925f6f-b9d1-7fd4-bf0e-53b8d1abfa96'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    ),
    (
        '01925fff-fd06-725b-95d5-6ead354f90e5'::uuid,
        '01925f6f-b9d1-78f4-83aa-9b6b9f70d892'::uuid, null,
        'This is a test post', current_timestamp, 1, 0
    );


-- Insert post tag records
insert into blabber.post_tag
    (post_id, tag)
values
    -- Every post has the 'test' tag
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, 'test'),
    ('01925fff-fd06-75aa-8fc3-2814af59bfb1'::uuid, 'test'),
    ('01925fff-fd06-77e4-b72c-fc23c4231069'::uuid, 'test'),
    ('01925fff-fd06-7288-9b7b-963b7858092f'::uuid, 'test'),
    ('01925fff-fd06-70bf-b9b6-743eee142ca1'::uuid, 'test'),
    ('01925fff-fd06-7b1a-b185-bf2c406ee201'::uuid, 'test'),
    ('01925fff-fd06-7435-892d-49ed8a142afa'::uuid, 'test'),
    ('01925fff-fd06-742f-ad45-f3206a7ae19b'::uuid, 'test'),
    ('01925fff-fd06-7483-9ef5-7ac80538518a'::uuid, 'test'),
    ('01925fff-fd06-736e-97b6-0064e63f20a4'::uuid, 'test'),
    ('01925fff-fd06-7e7b-b8cc-55eb2fa7c699'::uuid, 'test'),
    ('01925fff-fd06-7952-b621-999d7c671f2d'::uuid, 'test'),
    ('01925fff-fd06-799e-8768-e3db0e1f22af'::uuid, 'test'),
    ('01925fff-fd06-76b7-be15-df67b7674ca5'::uuid, 'test'),
    ('01925fff-fd06-7bd2-9bbf-7c7b42d154ec'::uuid, 'test'),
    ('01925fff-fd06-7a61-a371-4fc28a6009e6'::uuid, 'test'),
    ('01925fff-fd06-7f3b-ba38-1963d19a0548'::uuid, 'test'),
    ('01925fff-fd06-7ab5-b2f0-fcb2dae91430'::uuid, 'test'),
    ('01925fff-fd06-7821-aa21-c9896a8155d2'::uuid, 'test'),
    ('01925fff-fd06-71f6-826f-a68c2156687f'::uuid, 'test'),
    ('01925fff-fd06-715e-ab7e-054bea3cd171'::uuid, 'test'),
    ('01925fff-fd06-7f69-b360-7780747f71dc'::uuid, 'test'),
    ('01925fff-fd06-70f2-901e-d1905ab8bd09'::uuid, 'test'),
    ('01925fff-fd06-79d1-abad-ad5eb02d9177'::uuid, 'test'),
    ('01925fff-fd06-725b-95d5-6ead354f90e5'::uuid, 'test'),
    -- First ten posts have the 'extra' tag
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, 'extra'),
    ('01925fff-fd06-75aa-8fc3-2814af59bfb1'::uuid, 'extra'),
    ('01925fff-fd06-77e4-b72c-fc23c4231069'::uuid, 'extra'),
    ('01925fff-fd06-7288-9b7b-963b7858092f'::uuid, 'extra'),
    ('01925fff-fd06-70bf-b9b6-743eee142ca1'::uuid, 'extra'),
    ('01925fff-fd06-7b1a-b185-bf2c406ee201'::uuid, 'extra'),
    ('01925fff-fd06-7435-892d-49ed8a142afa'::uuid, 'extra'),
    ('01925fff-fd06-742f-ad45-f3206a7ae19b'::uuid, 'extra'),
    ('01925fff-fd06-7483-9ef5-7ac80538518a'::uuid, 'extra'),
    ('01925fff-fd06-736e-97b6-0064e63f20a4'::uuid, 'extra'),
    -- Next ten posts have the 'blabber' tag
    ('01925fff-fd06-7e7b-b8cc-55eb2fa7c699'::uuid, 'blabber'),
    ('01925fff-fd06-7952-b621-999d7c671f2d'::uuid, 'blabber'),
    ('01925fff-fd06-799e-8768-e3db0e1f22af'::uuid, 'blabber'),
    ('01925fff-fd06-76b7-be15-df67b7674ca5'::uuid, 'blabber'),
    ('01925fff-fd06-7bd2-9bbf-7c7b42d154ec'::uuid, 'blabber'),
    ('01925fff-fd06-7a61-a371-4fc28a6009e6'::uuid, 'blabber'),
    ('01925fff-fd06-7f3b-ba38-1963d19a0548'::uuid, 'blabber'),
    ('01925fff-fd06-7ab5-b2f0-fcb2dae91430'::uuid, 'blabber'),
    ('01925fff-fd06-7821-aa21-c9896a8155d2'::uuid, 'blabber'),
    ('01925fff-fd06-71f6-826f-a68c2156687f'::uuid, 'blabber'),
    -- Next five posts have the 'random' tag
    ('01925fff-fd06-715e-ab7e-054bea3cd171'::uuid, 'random'),
    ('01925fff-fd06-7f69-b360-7780747f71dc'::uuid, 'random'),
    ('01925fff-fd06-70f2-901e-d1905ab8bd09'::uuid, 'random'),
    ('01925fff-fd06-79d1-abad-ad5eb02d9177'::uuid, 'random'),
    ('01925fff-fd06-725b-95d5-6ead354f90e5'::uuid, 'random');


-- Insert post like records
insert into blabber.post_like
    (post_id, user_id, created_at)
values
    -- testuser01 likes every post (even their own, which is weird :P)
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-75aa-8fc3-2814af59bfb1'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-77e4-b72c-fc23c4231069'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7288-9b7b-963b7858092f'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-70bf-b9b6-743eee142ca1'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7b1a-b185-bf2c406ee201'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7435-892d-49ed8a142afa'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-742f-ad45-f3206a7ae19b'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7483-9ef5-7ac80538518a'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-736e-97b6-0064e63f20a4'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7e7b-b8cc-55eb2fa7c699'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7952-b621-999d7c671f2d'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-799e-8768-e3db0e1f22af'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-76b7-be15-df67b7674ca5'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7bd2-9bbf-7c7b42d154ec'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7a61-a371-4fc28a6009e6'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7f3b-ba38-1963d19a0548'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7ab5-b2f0-fcb2dae91430'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7821-aa21-c9896a8155d2'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-71f6-826f-a68c2156687f'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-715e-ab7e-054bea3cd171'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-7f69-b360-7780747f71dc'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-70f2-901e-d1905ab8bd09'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-79d1-abad-ad5eb02d9177'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    ('01925fff-fd06-725b-95d5-6ead354f90e5'::uuid, '01925f6f-b9d1-7a13-8c33-b23ca1225674'::uuid, current_timestamp),
    -- Every other user likes testuser01's post (quite popular!)
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7d1c-b89b-9791083af810'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7486-a4b6-5c7be7d07696'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7ae1-8a34-eefc726633bb'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-79c5-a81d-0f03b2f5fb2e'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-786e-b3dc-8fabf8cfae15'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-74a5-bca8-9b5dc7953e09'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-70f3-88a8-1d1f09c7d06b'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7660-8d04-4258fda596dd'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-70ab-923b-70ca533c9a04'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-78d2-a667-ca72e0df19ae'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-75f2-bbf9-cda51ac07651'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7a62-bf11-4ea07f543b78'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7936-aea5-f1ef9d9cfe60'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7608-b4c1-978f71f706d6'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-78ba-b874-b1620271cd2b'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7ee2-910f-bfc06b10c955'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-776b-a49b-259f0bdc2d78'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7481-ae6d-53e24abda1a7'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7eb3-a795-3174135fe8a2'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7ce9-a57a-08468da08e7e'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-79b7-a4da-843bd928986b'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-785a-af81-2cd2dd032963'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-7fd4-bf0e-53b8d1abfa96'::uuid, current_timestamp),
    ('01925fff-fd06-7419-bc00-1a57b59ba1ec'::uuid, '01925f6f-b9d1-78f4-83aa-9b6b9f70d892'::uuid, current_timestamp);


select id, user_id, contents, created_at, likes, reposts from blabber.post;

select
    p.id,
    count(distinct pl.user_id) as like_count
from blabber.post p
left join blabber.post_like pl
    on p.id = pl.post_id
group by p.id;