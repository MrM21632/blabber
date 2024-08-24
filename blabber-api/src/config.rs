use std::sync::Arc;

use clap::Parser;
use sqlx::PgPool;

#[derive(Parser)]
pub struct Config {
    #[clap(long, env)]
    pub database_url: String,
}

#[derive(Clone)]
pub(crate) struct ApiContext {
    pub config: Arc<Config>,
    pub database: PgPool,
}
