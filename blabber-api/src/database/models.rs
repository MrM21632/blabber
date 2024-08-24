use chrono::{DateTime, Utc};
use sqlx::prelude::FromRow;
use uuid::Uuid;

#[derive(FromRow, Debug, Clone)]
pub struct User {
    pub id: Uuid,
    pub username: String,
    pub user_handle: String,
    pub user_bio: String,
    pub email: String,
    pub password_hash: String,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
    pub followers: i32,
    pub follows: i32,
}

#[derive(FromRow, Debug, Clone)]
pub struct UserFollowing {
    pub follower_id: Uuid,
    pub followed_id: Uuid,
    pub created_at: DateTime<Utc>,
}

#[derive(FromRow, Debug, Clone)]
pub struct UserBlock {
    pub blocker_id: Uuid,
    pub blocked_id: Uuid,
    pub created_at: DateTime<Utc>,
}

#[derive(FromRow, Debug, Clone)]
pub struct UserMute {
    pub muter_id: Uuid,
    pub muted_id: Uuid,
    pub created_at: DateTime<Utc>,
}

#[derive(FromRow, Debug, Clone)]
pub struct Post {
    pub id: Uuid,
    pub user_id: Uuid,
    pub parent_id: Uuid,
    pub contents: String,
    pub created_at: DateTime<Utc>,
    pub likes: i32,
    pub reposts: i32,
}

#[derive(FromRow, Debug, Clone)]
pub struct Tag {
    pub post_id: Uuid,
    pub tag: String,
}

#[derive(FromRow, Debug, Clone)]
pub struct Like {
    pub post_id: Uuid,
    pub user_id: Uuid,
    pub created_at: DateTime<Utc>,
}

#[derive(FromRow, Debug, Clone)]
pub struct Repost {
    pub post_id: Uuid,
    pub user_id: Uuid,
    pub created_at: DateTime<Utc>,
}
