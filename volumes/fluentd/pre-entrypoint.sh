#!/bin/bash

export DOCKER_POSTGRES_QUERYLOG=$(\
    for i in $(ls -1 /var/lib/docker/containers); \
    do \
      if [[ \
          "$(jq .Config.Image /var/lib/docker/containers/$i/config.v2.json)" \
          == \
          "\"mrecco/pglog-collector:v1.0.0\"" \
      ]]; \
      then \
        jq .LogPath /var/lib/docker/containers/$i/config.v2.json | tr -d '"'; \
      fi; \
    done; \
)
echo "Using ${DOCKER_POSTGRES_QUERYLOG} for scrape nginx logs"
sed -ie "s#@@@DOCKER_POSTGRES_QUERYLOG@@@#"$DOCKER_POSTGRES_QUERYLOG"#g" /fluentd/etc/*.conf
