//use arrow::record_batch::RecordBatch;
use influxdb_iox_client::{connection::Builder, format::QueryOutputFormat, flight};
use influxdb_iox_client::connection::Error;
use futures::stream::TryStreamExt; // Import TryStreamExt trait for try_collect

pub type Result<T, E = Error> = std::result::Result<T, E>;


#[tokio::main]
async fn main() {

    let connection = Builder::default()
        .build("http://127.0.0.1:8082")
        .await
        .unwrap();


    let sql = String::from("select * from cpu limit 1");
    let namespace = String::from("company_sensors");

    let mut client = flight::Client::new(connection);

    client.add_header("bucket", &namespace).unwrap();

    let mut query_results = client.sql(namespace, &sql).await;

    let batches: Vec<_> = match &mut query_results {
        Ok(stream) => stream.try_collect().await.unwrap(),
        Err(err) => {
            eprintln!("Error: {}", err);
            Vec::new()
        }
    };

    let formatx = QueryOutputFormat::Pretty;
    println!("{}", formatx.format(&batches).unwrap());

}
