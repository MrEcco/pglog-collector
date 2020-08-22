FROM ubuntu:focal

RUN export DEBIAN_FRONTEND=noninteractive          && \
    ln -fs /usr/share/zoneinfo/Etc/UTC /etc/localtime && \
    apt-get update                                 && \
    apt-get install -y --no-install-recommends        \
        ca-certificates                               \
        curl                                          \
        gnupg2                                        \
        wget                                       && \
    apt-get clean -y                               && \
    apt-get autoclean -y                           && \
    rm -rf /var/lib/apt/lists/*

ARG JOBBER_DEB_PACKAGE="https://github.com/dshearer/jobber/releases/download/v1.4.4/jobber_1.4.4-1_amd64.deb"
ENV JOBBER_VERSION="1.4.4"

RUN curl -L ${JOBBER_DEB_PACKAGE}                     \
        -o /root/jobber.deb                        && \
    dpkg -i /root/jobber.deb                       && \
    rm /root/jobber.deb                            && \
    mkdir -p /var/lib/jobber /etc/jobber           && \
    useradd jobber                                    \
            -d /var/lib/jobber -M                     \
            -s /usr/sbin/nologin                      \
            -u 999 -U                              && \
    chown jobber:jobber /var/lib/jobber

USER jobber

ENTRYPOINT [ "/usr/lib/x86_64-linux-gnu/jobberrunner", "/etc/jobber/jobber.yml" ]
