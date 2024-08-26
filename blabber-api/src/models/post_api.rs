use serde::{Deserialize, Serialize};
use uuid::Uuid;

/// Wrapper object for requests to the main /posts endpoints
#[derive(Serialize, Deserialize)]
pub struct PostRequestBody<T> {
    post: T,
}