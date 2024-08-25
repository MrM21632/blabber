use std::{sync::Arc, time::Duration};

use anyhow::Context;
use axum::{extract::MatchedPath, http::{header::AUTHORIZATION, Request}, routing::get, Router};
use clap::Parser;
use config::Config;
use sqlx::{postgres::PgPoolOptions, PgPool};
use tower_http::{catch_panic::CatchPanicLayer, sensitive_headers::SetSensitiveHeadersLayer, timeout::TimeoutLayer, trace::TraceLayer};
use tracing::{info_span, level_filters::LevelFilter};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt, EnvFilter};

mod config;
mod error;

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
    tracing_subscriber::registry()
        .with(
            EnvFilter::builder()
                .with_default_directive(LevelFilter::ERROR.into())
                .from_env_lossy()
        )
        .with(tracing_subscriber::fmt::layer())
        .init();
    Router::new()
        .route("/hello", get(|| async { "Hello, world!" }))
        .layer((
            TraceLayer::new_for_http()
                .make_span_with(|req: &Request<_>| {
                    let matched_path = req.extensions()
                        .get::<MatchedPath>()
                        .map(MatchedPath::as_str);

                    info_span!(
                        "http_request",
                        method = ?req.method(),
                        matched_path,
                    )
                }),
            SetSensitiveHeadersLayer::new([AUTHORIZATION]),
            TimeoutLayer::new(Duration::from_secs(30)),
            CatchPanicLayer::new(),
        ))
        .with_state(context)
}

async fn serve(config: Config, database: PgPool) -> anyhow::Result<()> {
    let context = config::ApiContext {
        config: Arc::new(config),
        database,
    };

    let app = create_router(context);
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();

    axum::serve(listener, app)
        .with_graceful_shutdown(shutdown_signal())
        .await
        .context("Error occurred starting HTTP server")
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    dotenvy::dotenv().ok();
    let config = Config::parse();

    let database = PgPoolOptions::new()
        .max_connections(10)
        .connect(&config.database_url)
        .await
        .expect("Error establishing database connection");
    
    sqlx::migrate!()
        .run(&database)
        .await?;

    serve(config, database).await?;

    Ok(())
}
