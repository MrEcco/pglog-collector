FROM fluent/fluentd-kubernetes-daemonset:v1.3.3-debian-graylog-1.3

RUN apt-get update                                         \
   && apt-get install -y --no-install-recommends           \
         sudo             \
         jq               \
         net-tools        \
         curl             \
         nano             \
         ca-certificates  \
         logrotate        \
         xz-utils         \
   && apt-get clean -y                                     \
   && rm -rf /var/lib/apt/lists/*                          \
   && mkdir -p /var/log/fluentd-pos                        \
   && touch /var/lib/logrotate/status                      \
   && chmod 777 /var/lib/logrotate/status                  \
   && fluent-gem install fluent-plugin-multi-format-parser
