#!/bin/bash -xe

sudo syslogd
sleep 1
sudo service ssh restart

ifconfig

tail -f /var/log/syslog