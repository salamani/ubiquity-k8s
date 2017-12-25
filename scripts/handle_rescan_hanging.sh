#!/bin/bash -xe
# Description : If hanging $cmd, then clean faulty devices, 
#                      If the hanging $cmd no longer exist(after 20 seconds) script finish with success else with exit code >0

function abort() { echo $1; exit $2; }
function clean_up_faulty_mpath_devices() {
   mp_fault_devices=`multipath -ll | grep "dm-[0-9]\+" | grep -v IBM | awk '{print $1}'`
   [ -z "$mp_fault_devices" ] && abort "There are no faulty mpath to clean, exit; " 5 || echo "cleaning the following faulty devices : $mp_fault_devices"
   for mpath in $mp_fault_devices; do
        dmsetup message /dev/mapper/$mpath  0 fail_if_no_path || echo "Warning: dmsetup message on $mpath failed with exit code [$?]. Ignore and continue clean up"
        multipath -f $mpath || echo "Warning: multipath -f $mpath command failed with exit code [$?]. Ignore and continue clean up"
   done
}

too_long_running=180   # 3 minutes
[ -n "$1" ] && cmd=$1 || cmd=rescan-scsi-bus.sh
pid=`ps -ef| grep -w $cmd | grep -v grep | awk '{print $2}'`
[ -z "$pid" ] && abort "$cmd process not found" 1 || :
running_time=`ps -ef -o etimes= -p $pid| awk '{print $1}'`

echo "$running_time" | grep "^[0-9]\+$" || abort "$cmd mtime is not valid $running_time. exit" 2
[ $running_time -lt $too_long_running ] && abort "$cmd running $running_time seconds, its ok, not reached max time of $too_long_running" 3  || :
echo "$cmd is running too long $running_time > $too_long_running. Found hanging $cmd"

clean_up_faulty_mpath_devices

echo "sleeping 20 seconds to let $cmd terminating..."; sleep 20 ;
pid2=`ps -ef| grep -w $cmd | grep -v grep | awk '{print $2}'`
[ "$pid2" != "$pid" ] && { echo "fixed $cmd hanging by cleaning mpath devices"; exit 0; } || abort "$cmd still hanging even after cleanup the faulty devices" 7

