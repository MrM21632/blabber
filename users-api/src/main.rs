use clap::Parser;
use config::Config;
use sqlx::postgres::PgPoolOptions;

mod config;

#[tokio::main]
async fn main() {
    dotenvy::dotenv().ok();
    let config = Config::parse();

    let database = PgPoolOptions::new()
        .max_connections(10)
        .connect(&config.database_url)
        .await
        .unwrap();

    let _ = sqlx::migrate!().run(&database).await.unwrap();
}
