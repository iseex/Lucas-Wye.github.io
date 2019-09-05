#! /bin/bash
#==============================================================#
#   Description:  Unixbench script                             #
#   Author: Teddysun <i@teddysun.com>                          #
#   Intro:  https://teddysun.com/245.html                      #
#==============================================================#

# Check System
[[ $EUID -ne 0 ]] && echo 'Error: This script must be run as root!' && exit 1
[[ -f /etc/redhat-release ]] && os='centos'
[[ ! -z "`egrep -i debian /etc/issue`" ]] && os='debian'
[[ ! -z "`egrep -i ubuntu /etc/issue`" ]] && os='ubuntu'
[[ "$os" == '' ]] && echo 'Error: Your system is not supported to run it!' && exit 1

# Install necessary libaries
if [ "$os" == 'centos' ]; then
    yum -y install make automake gcc autoconf gcc-c++ time perl-Time-HiRes
else
    apt-get -y update
    apt-get -y install make automake gcc autoconf time perl
fi

#Run unixbench
make
./Run

echo ''
echo ''
echo "======= Script description and score comparison completed! ======= "
echo ''
echo ''
