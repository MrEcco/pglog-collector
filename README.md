# Postgres query log collector

This is little app for collect logs from postgres csv logs to json. This is implement like pipeline-to-pipeline formatter. Logs ready for inject to graylog (via fluent, by example).

## Configuration

### Postgres

We must enable **csvlog** behavior and disable integrated rotation (because we use **logrotate** for this). All changes on **postgresql.conf** file.

1. Add 'csvlog' to **log_destination**.
2. Set **logging_collector** to **on**. This is required for logging to CSV format.
3. Define **log_directory** and **log_filename**. You must know where you store logs. Set filename without any dynamic parameters.
4. Set **log_file_mode** to **600**.
5. Set **log_rotation_size** and **log_rotation_age** to **0**.
6. Set **log_truncate_on_rotation** to **off**

See **example.postgresql.conf**: this contain compliant log-related parameters.

### Logrotate

See **volumes/logrotate/logrotate.conf**.

### Fluentd

All configured and ready to send to graylog. See **volumes/fluentd** directory.

## Solution (why this way and dont otherwise)

My code just convert ugly CSV format to JSON. JSON can be pushed to graylog by other applications and I satisfied how fluentd work for this.

## Reference

### Based

Based on PostgreSQL Server. And this all for this!

https://www.postgresql.org/docs/11/runtime-config-logging.html

### Dependencies

https://github.com/hpcloud/tail

## Tips

Give me 0.1 BTC if you want me to never write code.
Give me 1 BTC if you want me to believe in God.
Give me 10 BTC - I will become a monk.

BTC: 32gLMtXkEsYK8nKbYTuhUMAAQxBTj7WUqr

Glad each Satoshi :)
