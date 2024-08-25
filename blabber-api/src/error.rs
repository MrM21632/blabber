use std::{borrow::Cow, collections::HashMap};

use axum::{http::{header::WWW_AUTHENTICATE, StatusCode}, response::IntoResponse, Json};

/// Common error type which can be used across the API for error handling needs.
#[derive(thiserror::Error, Debug)]
pub enum Error {
    /// HTTP 401 Unauthorized errors
    #[error("Authentication Required")]
    Unauthorized,

    /// HTTP 403 Forbidden errors
    #[error("User Prohibited from Performing Action")]
    Forbidden,

    /// HTTP 404 Not Found errors
    #[error("Requested Path Not Found")]
    NotFound,

    /// HTTP 422 Unprocessable Entity errors
    #[error("Data Received is Unprocessable")]
    UnprocessableEntity {
        errors: HashMap<Cow<'static, str>, Vec<Cow<'static, str>>>,
    },

    /// Errors with the database; treated as 500 Internal Server Error
    #[error("Database Error")]
    SqlxError(#[from] sqlx::Error),

    /// Generic handler for internal errors; treated as 500 Internal Server Error
    #[error("Internal Server Error")]
    InternalError(#[from] anyhow::Error),
}

impl Error {
    /// Constructor method for UnprocessableEntity errors
    pub fn unprocessable_entity<K, V>(errors: impl IntoIterator<Item = (K, V)>) -> Self
    where
        K: Into<Cow<'static, str>>,
        V: Into<Cow<'static, str>>,
    {
        let mut error_map = HashMap::new();
        for (key, val) in errors {
            error_map
                .entry(key.into())
                .or_insert_with(Vec::new)
                .push(val.into());
        }

        Self::UnprocessableEntity { errors: error_map }
    }

    fn status_code(&self) -> StatusCode {
        match self {
            Self::Unauthorized => StatusCode::UNAUTHORIZED,
            Self::Forbidden => StatusCode::FORBIDDEN,
            Self::NotFound => StatusCode::NOT_FOUND,
            Self::UnprocessableEntity { .. } => StatusCode::UNPROCESSABLE_ENTITY,
            Self::SqlxError(_) | Self::InternalError(_) => StatusCode::INTERNAL_SERVER_ERROR,
        }
    }
}

impl IntoResponse for Error {
    fn into_response(self) -> axum::response::Response {
        match self {
            Self::UnprocessableEntity { errors } => {
                #[derive(serde::Serialize)]
                struct Errors {
                    errors: HashMap<Cow<'static, str>, Vec<Cow<'static, str>>>,
                }

                return (StatusCode::UNPROCESSABLE_ENTITY, Json(Errors { errors })).into_response();
            }
            Self::Unauthorized => {
                return (
                    self.status_code(),
                    [(WWW_AUTHENTICATE, "Token")],
                    self.to_string(),
                ).into_response();
            }
            Self::SqlxError(ref e) => {
                // TODO: Add tracing as a dependency
                // tracing::span!(format!("Database error occurred: {:?}", e));
            }
            Self::InternalError(ref e) => {
                // TODO: Add tracing as a dependency
                // tracing::span!(format!("Internal error occurred: {:?}", e));
            }

            _ => (),
        }

        (self.status_code(), self.to_string()).into_response()
    }
}

/// Helper trait for converting constraint errors into API errors
pub trait ResultExt<T> {
    fn on_failed_constraint(
        self,
        name: &str,
        f: impl FnOnce(Box<dyn sqlx::error::DatabaseError>) -> Error
    ) -> Result<T, Error>;
}

impl<T, E> ResultExt<T> for Result<T, E>
where
    E: Into<Error>,
{
    fn on_failed_constraint(
        self,
        name: &str,
        map_err: impl FnOnce(Box<dyn sqlx::error::DatabaseError>) -> Error
    ) -> Result<T, Error> {
        self.map_err(|e| match e.into() {
            Error::SqlxError(sqlx::Error::Database(dbe)) if dbe.constraint() == Some(name) => {
                map_err(dbe)
            }
            e => e,
        })
    }
}
