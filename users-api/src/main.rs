use axum::{routing::get, Router};
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
        .expect("Error establishing database connection");
    
    sqlx::migrate!()
        .run(&database)
        .await
        .expect("Migration failed");

    let app = Router::new().route("/hello", get(|| async { "Hello, world!" }));
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();

    axum::serve(listener, app).await.unwrap();
}
