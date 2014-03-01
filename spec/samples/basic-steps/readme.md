```
# mkfs.ext3 /dev/sdb
mke2fs 1.42 (29-Nov-2011)
/dev/sdb is entire device, not just one partition!
Proceed anyway? (y,n) y
Filesystem label=
OS type: Linux
Block size=4096 (log=2)
Fragment size=4096 (log=2)
Stride=0 blocks, Stripe width=0 blocks
320000 inodes, 1280000 blocks
64000 blocks (5.00%) reserved for the super user
First data block=0
Maximum filesystem blocks=1312817152
40 block groups
32768 blocks per group, 32768 fragments per group
8000 inodes per group
Superblock backups stored on blocks:
	32768, 98304, 163840, 229376, 294912, 819200, 884736

Allocating group tables: done
Writing inode tables: done
Creating journal (32768 blocks): done
Writing superblocks and filesystem accounting information: done
```

```
# pvcreate /dev/sdb
  Physical volume "/dev/sdb" successfully created
```

```
# pvdisplay -c
  /dev/sda5:vg0:84254720:-1:8:8:-1:4096:10284:0:10284:IKGNO5-Dx7w-2UBv-rUzw-ekJg-e496-9RQ5cP
  "/dev/sdb" is a new physical volume of "4.88 GiB"
  /dev/sdb:#orphans_lvm2:10240000:-1:8:8:-1:0:0:0:0:ByVn3c-aI8G-BYYI-CS84-QBPO-cOsa-34kHWu
```

```
# pvdisplay
  --- Physical volume ---
  PV Name               /dev/sda5
  VG Name               vg0
  PV Size               40.18 GiB / not usable 4.00 MiB
  Allocatable           yes (but full)
  PE Size               4.00 MiB
  Total PE              10284
  Free PE               0
  Allocated PE          10284
  PV UUID               IKGNO5-Dx7w-2UBv-rUzw-ekJg-e496-9RQ5cP

  "/dev/sdb" is a new physical volume of "4.88 GiB"
  --- NEW Physical volume ---
  PV Name               /dev/sdb
  VG Name
  PV Size               4.88 GiB
  Allocatable           NO
  PE Size               0
  Total PE              0
  Free PE               0
  Allocated PE          0
  PV UUID               ByVn3c-aI8G-BYYI-CS84-QBPO-cOsa-34kHWu
```

```
# vgcreate vg1 /dev/sdb
  Volume group "vg1" successfully created
```

```
# vgdisplay -c
  vg1:r/w:772:-1:0:0:0:-1:0:1:1:5115904:4096:1249:0:1249:SBnFzt-Kqor-llgF-NODX-MrB4-QS7S-VVk6wI
  vg0:r/w:772:-1:0:2:2:-1:0:1:1:42123264:4096:10284:10284:0:TKXx5M-QtcZ-deWE-4TSP-G1Ec-uRFF-MtWTyH
```

```
# vgdisplay
  --- Volume group ---
  VG Name               vg1
  System ID
  Format                lvm2
  Metadata Areas        1
  Metadata Sequence No  1
  VG Access             read/write
  VG Status             resizable
  MAX LV                0
  Cur LV                0
  Open LV               0
  Max PV                0
  Cur PV                1
  Act PV                1
  VG Size               4.88 GiB
  PE Size               4.00 MiB
  Total PE              1249
  Alloc PE / Size       0 / 0
  Free  PE / Size       1249 / 4.88 GiB
  VG UUID               SBnFzt-Kqor-llgF-NODX-MrB4-QS7S-VVk6wI

  --- Volume group ---
  VG Name               vg0
  System ID
  Format                lvm2
  Metadata Areas        1
  Metadata Sequence No  3
  VG Access             read/write
  VG Status             resizable
  MAX LV                0
  Cur LV                2
  Open LV               2
  Max PV                0
  Cur PV                1
  Act PV                1
  VG Size               40.17 GiB
  PE Size               4.00 MiB
  Total PE              10284
  Alloc PE / Size       10284 / 40.17 GiB
  Free  PE / Size       0 / 0
  VG UUID               TKXx5M-QtcZ-deWE-4TSP-G1Ec-uRFF-MtWTyH
```

## Logical volumes

```
# lvcreate -l 20 -n test1 vg1
  Logical volume "test1" created
```

```
# lvdisplay
  --- Logical volume ---
  LV Name                /dev/vg1/test1
  VG Name                vg1
  LV UUID                pf9gtT-Rsia-5CSU-QFkc-tr5t-UNfG-qRGIjx
  LV Write Access        read/write
  LV Status              available
  # open                 0
  LV Size                80.00 MiB
  Current LE             20
  Segments               1
  Allocation             inherit
  Read ahead sectors     auto
  - currently set to     256
  Block device           252:2
...
```

```
# lvdisplay -c
  /dev/vg1/test1:vg1:3:1:-1:0:163840:20:-1:0:-1:252:2
  /dev/vg0/root:vg0:3:1:-1:1:78118912:9536:-1:0:-1:252:0
  /dev/vg0/swap:vg0:3:1:-1:2:6127616:748:-1:0:-1:252:1
```

```
# mkfs.ext4 /dev/vg1/test1
mke2fs 1.42 (29-Nov-2011)
Filesystem label=
OS type: Linux
Block size=1024 (log=0)
Fragment size=1024 (log=0)
Stride=0 blocks, Stripe width=0 blocks
20480 inodes, 81920 blocks
4096 blocks (5.00%) reserved for the super user
First data block=1
Maximum filesystem blocks=67371008
10 block groups
8192 blocks per group, 8192 fragments per group
2048 inodes per group
Superblock backups stored on blocks:
	8193, 24577, 40961, 57345, 73729

Allocating group tables: done
Writing inode tables: done
Creating journal (4096 blocks): done
Writing superblocks and filesystem accounting information: done
```
