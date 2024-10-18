# Log Store

A store which allows you to store information on Disk in the log format


# Log format
```
-- 1 byte: operation
    insert: 01
    delete: 02
-- 4 byte: key-length: kn
-- 4 byte: value-length: vn
-- kn bytes: key
-- vn bytes: value
...
```

Keys and Values are both strings
