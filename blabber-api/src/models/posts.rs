use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use sqlx::prelude::FromRow;
use uuid::Uuid;

#[derive(FromRow, Debug, Clone, Serialize, Deserialize)]
pub struct Post {
    pub id: Uuid,
    pub user_id: Uuid,
    pub parent_id: Uuid,
    pub contents: String,
    pub created_at: DateTime<Utc>,
    pub likes: i32,
    pub reposts: i32,
}

#[derive(FromRow, Debug, Clone, Serialize, Deserialize)]
pub struct Tag {
    pub post_id: Uuid,
    pub tag: String,
}

#[derive(FromRow, Debug, Clone, Serialize, Deserialize)]
pub struct Like {
    pub post_id: Uuid,
    pub user_id: Uuid,
    pub created_at: DateTime<Utc>,
}

#[derive(FromRow, Debug, Clone, Serialize, Deserialize)]
pub struct Repost {
    pub post_id: Uuid,
    pub user_id: Uuid,
    pub created_at: DateTime<Utc>,
}
