#
# Regular cron jobs for the gocrypto package
#
0 4	* * *	root	[ -x /usr/bin/gocrypto_maintenance ] && /usr/bin/gocrypto_maintenance
