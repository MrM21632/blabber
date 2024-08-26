use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use sqlx::prelude::FromRow;
use uuid::Uuid;

#[derive(FromRow, Debug, Clone, Serialize, Deserialize)]
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

#[derive(FromRow, Debug, Clone, Serialize, Deserialize)]
pub struct UserFollowing {
    pub follower_id: Uuid,
    pub followed_id: Uuid,
    pub created_at: DateTime<Utc>,
}

#[derive(FromRow, Debug, Clone, Serialize, Deserialize)]
pub struct UserBlock {
    pub blocker_id: Uuid,
    pub blocked_id: Uuid,
    pub created_at: DateTime<Utc>,
}

#[derive(FromRow, Debug, Clone, Serialize, Deserialize)]
pub struct UserMute {
    pub muter_id: Uuid,
    pub muted_id: Uuid,
    pub created_at: DateTime<Utc>,
}
