use serde::{Deserialize, Serialize};
use uuid::Uuid;

/// Wrapper object for requests to the main /users endpoints
#[derive(Serialize, Deserialize)]
pub struct UserRequestBody<T> {
    user: T,
}

/// Details for creating a new user account
#[derive(Deserialize)]
pub struct NewUserDetails {
    username: String,
    user_handle: String,
    user_bio: Option<String>,
    email: String,
    password: String,
}

/// Details for updating a user account
#[derive(Deserialize)]
pub struct UpdateUserDetails {
    id: Uuid,
    username: Option<String>,
    user_handle: Option<String>,
    user_bio: Option<String>,
    email: Option<String>,
    password: Option<String>,
}

/// Details for specifically updating a user's password
#[derive(Deserialize)]
pub struct UpdatePasswordDetails {
    id: Uuid,
    old_password: String,
    new_password: String,
}

/// Details for retrieving a specific user account's details
#[derive(Deserialize)]
pub struct GetUserDetails {
    id: Uuid,
}

/// Details for creating a new user follow relation
#[derive(Deserialize)]
pub struct FollowUserDetails {
    follower_id: Uuid,
    followed_id: Uuid,
}

/// Details for getting a user's followers
#[derive(Deserialize)]
pub struct GetFollowersDetails {
    followed_id: Uuid,
}

/// Details for getting a user's follows
#[derive(Deserialize)]
pub struct GetFollowsDetails {
    follower_id: Uuid,
}

/// Details for creating a new user mute relation
#[derive(Deserialize)]
pub struct MuteUserDetails {
    muter_id: Uuid,
    muted_id: Uuid,
}

/// Details for getting a user's muters
#[derive(Deserialize)]
pub struct GetMutersDetails {
    muted_id: Uuid,
}

/// Details for getting a user's mutes
#[derive(Deserialize)]
pub struct GetMutesDetails {
    muter_id: Uuid,
}

/// Details for creating a new user block relation
#[derive(Deserialize)]
pub struct BlockUserDetails {
    blocker_id: Uuid,
    blocked_id: Uuid,
}

/// Details for getting a user's blockers
#[derive(Deserialize)]
pub struct GetBlockersDetails {
    blocked_id: Uuid,
}

/// Details for getting a user's blocks
#[derive(Deserialize)]
pub struct GetBlocksDetails {
    blocker_id: Uuid,
}
