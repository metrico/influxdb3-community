# IOx 
### FlightSQL Rust
Example gRPC Flight SQL API client using rust

#### Usage
```
cargo build
cargo run
```

#### Static
```bash
RUSTFLAGS='-C target-feature=+crt-static' cargo build --release --target x86_64-unknown-linux-gnu
```
```bash
ldd ./target/x86_64-unknown-linux-gnu/release/iox-client
	statically linked

```
