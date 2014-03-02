To create each of these files:

* `pvs` comes from `pvs --units=m --separator=: --nosuffix --noheadings`

Other examples:

```
# pvs --units=k --noheadings --separator=:
  /dev/sda5:precise64:lvm2:a-:83632128.00k:0k
# pvs --units=k --noheadings --separator=: --nosuffix
  /dev/sda5:precise64:lvm2:a-:83632128.00:0

# pvs --units=k --separator=: --nosuffix
  PV:VG:Fmt:Attr:PSize:PFree
  /dev/sda5:precise64:lvm2:a-:83632128.00:0

# pvs --units=m --separator=: --nosuffix
  PV:VG:Fmt:Attr:PSize:PFree
  /dev/sda5:precise64:lvm2:a-:81672.00:0
```

* `vgs` comes from `vgs --units=m --separator=: --nosuffix --noheadings`

```
# vgs --units=m --separator=: --nosuffix --noheadings
  precise64:1:2:0:wz--n-:81672.00:0
```

* `lvs` comes from `lvs --units=m --separator=: --nosuffix --noheadings`

```
# lvs --units=m --separator=: --nosuffix --noheadings
  root:precise64:-wi-ao:80904.00::::::
  swap_1:precise64:-wi-ao:768.00::::::

# lvs
  LV     VG        Attr   LSize   Origin Snap%  Move Log Copy%  Convert
  root   precise64 -wi-ao  79.01g
  swap_1 precise64 -wi-ao 768.00m
```

The lv_attr bits are:

1.  Volume type: (m)irrored, (M)irrored without initial sync, (o)rigin, (O)rigin with merging snapshot, (s)napshot,  merging  (S)napshot,  (p)vmove,  (v)irtual,
   mirror (i)mage, mirror (I)mage out-of-sync, under (c)onversion
2.  Permissions: (w)riteable, (r)ead-only
3.  Allocation  policy:  (c)ontiguous,  c(l)ing,  (n)ormal,  (a)nywhere,  (i)nherited
    This  is capitalised if the volume is currently locked against allocation changes, for example during pvmove (8).
4.  fixed (m)inor
5.  State: (a)ctive, (s)uspended, (I)nvalid snapshot, invalid (S)uspended snapshot, mapped (d)evice present without tables, mapped device present  with  (i)nactive table
6.  device (o)pen
