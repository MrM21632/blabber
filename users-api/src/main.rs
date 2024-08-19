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

    let row: (i64,) = sqlx::query_as("SELECT $1")
        .bind(150_i64)
        .fetch_one(&database)
        .await
        .expect("Error running test query");

    println!("Got: {:?}", row.0);
}
