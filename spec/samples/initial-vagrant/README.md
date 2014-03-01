To create each of these files:

* `pvs` comes from `pvs --units=m --separator=: --nosuffix`

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
