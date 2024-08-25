use anyhow::Context;
use argon2::{
    password_hash::{rand_core::OsRng, SaltString},
    Algorithm, Argon2, Params, PasswordHash, Version,
};

use crate::util::error::{Error, Result};

fn create_argon2_context() -> Argon2<'static> {
    Argon2::new(
        Algorithm::default(),
        Version::default(),
        Params::new(65536, 3, 2, Some(32)).unwrap(),
    )
}

pub async fn hash_password(password: String) -> Result<String> {
    tokio::task::spawn_blocking(move || -> Result<String> {
        let salt = SaltString::generate(&mut OsRng);
        let algo_config = create_argon2_context();

        Ok(
            PasswordHash::generate(algo_config, password, salt.as_salt())
                .map_err(|e| anyhow::anyhow!("Failed to generate password hash: {}", e))?
                .to_string(),
        )
    })
    .await
    .context("Panic occurred generating password hash")?
}

pub async fn verify_password(password: String, password_hash: String) -> Result<()> {
    tokio::task::spawn_blocking(move || -> Result<()> {
        let hash = PasswordHash::new(&password_hash)
            .map_err(|e| anyhow::anyhow!("Invalid password hash: {}", e))?;

        hash.verify_password(&[&create_argon2_context()], password)
            .map_err(|e| match e {
                argon2::password_hash::Error::Password => Error::Unauthorized,
                _ => anyhow::anyhow!("Failed to verify password hash: {}", e).into(),
            })
    })
    .await
    .context("Panic occurred verifying password hash")?
}
