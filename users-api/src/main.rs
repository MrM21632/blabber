use std::sync::Arc;

use axum::{routing::get, Router};
use clap::Parser;
use config::Config;
use sqlx::{postgres::PgPoolOptions, PgPool};

mod config;
mod database;

async fn shutdown_signal() {
    use tokio::signal;

    let ctrl_c = async {
        signal::ctrl_c().await.expect("Failed to install Ctrl+C handler");
    };

    #[cfg(unix)]
    let terminate = async {
        signal::unix::signal(signal::unix::SignalKind::terminate())
            .expect("Failed to install UNIX signal handler")
            .recv()
            .await;
    };

    #[cfg(not(unix))]
    let terminate = std::future::pending::<()>();

    tokio::select! {
        _ = ctrl_c => {},
        _ = terminate => {},
    }
}

fn create_router(context: config::ApiContext) -> Router {
    Router::new()
        .route("/hello", get(|| async { "Hello, world!" }))
        // TODO: Add middleware layers: Sensitive headers, tracing, timeouts, catching panics
        .with_state(context)
}

async fn serve(config: Config, database: PgPool) {
    let context = config::ApiContext {
        config: Arc::new(config),
        database,
    };

    let app = create_router(context);
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();

    axum::serve(listener, app)
        .with_graceful_shutdown(shutdown_signal())
        .await
        .unwrap();
}

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

    serve(config, database).await;
}
