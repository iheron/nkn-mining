#! /bin/sh

EINVAL=22

function Usage () {
    local RED="\E[1;31m"
    local GREEN="\E[1;32m"
    local YELLOW="\E[1;33m"
    local BLUE="\E[1;34m"
    local END="\E[0m"
    printf "${RED}Usage${END}: ${BLUE}%s${END} ${GREEN}<NodeNum>${END} ${YELLOW}[dest_dir]${END}\n" $1
    printf "${RED}Example${END}: ${BLUE}%s${END} ${GREEN}10${END} ${YELLOW}test_env${END}\n" $1
}


ECANCELED=125
[ -e './wallet.dat' ] && printf "The wallet.dat already existed\n" && exit ${ECANCELED}


### init wallet.dat
rm -f wallet.dat
RANDOM_PASSWD=$(head -c 1024 /dev/urandom | sha512sum | xxd -p | base64 | head -c 32)
./nknc wallet -c <<EOF
${RANDOM_PASSWD}
${RANDOM_PASSWD}
EOF
[ $? -eq 0 ] || exit $?
echo ${RANDOM_PASSWD} > ./wallet.pswd

