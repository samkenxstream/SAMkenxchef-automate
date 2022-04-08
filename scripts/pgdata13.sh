#!/bin/bash
#shellcheck disable=SC1128
#
shopt -s extglob

exec 2>&1

# *** WARNING ***
# Please put potentially long-running and/or complex operations in the init hook rather
# than the run hook until the issue described in
#
# https://github.com/habitat-sh/habitat/issues/1973
#
# has been resolved.
# Currently, the Habitat `init` and `health_check` hooks run directly from the main loop
# of the Habitat supervisor. If these hooks hang or take too long to run, they can block
# execution of the supervisor.
#


#shellcheck disable=SC1091
source /hab/svc/automate-postgresql/config/functions.sh

mkdir -p /hab/svc/automate-postgresql/config/conf.d
mkdir -p /hab/svc/automate-postgresql/var/pg_stat_tmp

mkdir -p /hab/svc/automate-postgresql/data/pgdata13
mkdir -p /hab/svc/automate-postgresql/data/archive

ensure_dir_ownership
ensure_key_ownership

if [[ ! -f "/hab/svc/automate-postgresql/data/pgdata13/PG_VERSION" ]]; then
  echo " Database does not exist, creating with 'initdb'"
    /hab/pkgs/core/postgresql13/13.5/20220120092917/bin/initdb -U automate \
    -E utf8 \
    -D /hab/svc/automate-postgresql/data/pgdata13 \
    --locale POSIX \
    --data-checksums
fi
