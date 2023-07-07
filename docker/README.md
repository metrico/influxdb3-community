<img src="https://pbs.twimg.com/profile_banners/1967601206/1682514855/1500x500" width=800>

# IOx Compose
This recipe starts an IOx API served by Nginx and backed by Minio/S3 _(data)_ and Postgres _(metadata)_


#### Docker
To start the IOx bundle, review the [settings](docker-compose.yml) and initialize using `docker-compose`

```
docker-compose up -d
```

Your local IOx endpoint should be ready on port `8086`

<br>
