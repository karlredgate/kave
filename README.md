kave
====

 * create vm
 * create vcpus
 * create memory region
 * set registers of vcpus

enable nested
FUSE /vm/... for managment?
No under 9p

```
# grep .  /sys/module/kvm_intel/parameters/*
/sys/module/kvm_intel/parameters/allow_smaller_maxphyaddr:N
/sys/module/kvm_intel/parameters/dump_invalid_vmcs:N
/sys/module/kvm_intel/parameters/emulate_invalid_guest_state:Y
/sys/module/kvm_intel/parameters/enable_apicv:Y
/sys/module/kvm_intel/parameters/enable_shadow_vmcs:Y
/sys/module/kvm_intel/parameters/enlightened_vmcs:N
/sys/module/kvm_intel/parameters/ept:Y
/sys/module/kvm_intel/parameters/eptad:Y
/sys/module/kvm_intel/parameters/fasteoi:Y
/sys/module/kvm_intel/parameters/flexpriority:Y
/sys/module/kvm_intel/parameters/nested:Y
/sys/module/kvm_intel/parameters/nested_early_check:N
/sys/module/kvm_intel/parameters/ple_gap:128
/sys/module/kvm_intel/parameters/ple_window:4096
/sys/module/kvm_intel/parameters/ple_window_grow:2
/sys/module/kvm_intel/parameters/ple_window_max:4294967295
/sys/module/kvm_intel/parameters/ple_window_shrink:0
/sys/module/kvm_intel/parameters/pml:Y
/sys/module/kvm_intel/parameters/preemption_timer:Y
/sys/module/kvm_intel/parameters/sgx:N
/sys/module/kvm_intel/parameters/unrestricted_guest:Y
/sys/module/kvm_intel/parameters/vmentry_l1d_flush:not required
/sys/module/kvm_intel/parameters/vnmi:Y
/sys/module/kvm_intel/parameters/vpid:Y
```


