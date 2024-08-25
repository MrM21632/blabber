use std::sync::Arc;

use clap::Parser;
use sqlx::PgPool;

#[derive(Parser)]
pub struct Config {
    #[clap(long, env)]
    pub database_url: String,

    #[clap(long, env)]
    pub server_port: u32,
}

#[derive(Clone)]
pub(crate) struct ApiContext {
    pub config: Arc<Config>,
    pub database: PgPool,
}
