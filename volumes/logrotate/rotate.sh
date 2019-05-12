#!/bin/bash

# Use envs:
# PGUID=$(id -u postgres)
# PGGID=$(id -g postgres)

function checkPasswd {
   if [[ $(cat /etc/passwd | grep -c postgres) -lt 1 ]]
   then
      echo "postgres:x:$PGUID:$PGGID::/:/bin/false" >> /etc/passwd
      echo "postgres:x:$PGGID:" >> /etc/group
   fi
}

checkPasswd
sudo -u postgres logrotate -f /etc/logrotate/logrotate.conf | true
