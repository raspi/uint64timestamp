# uint64timestamp

Timestamp fitted into uint64. 
There are two different versions: Base10 which is for humans and more binary packed Base2. 



|          | Base10       | Base2       |
|----------|--------------|-------------|
| Accuracy | microseconds | nanoseconds |



## Base10 - human parsable

See [pkg/base10](pkg/base10) for details.

Example:
`2023051514142573247` is 2023-05-15 14:14:25.73247

```go
now := time.Now()

bin, err := base10.TimeToUint64(now)
if err != nil {
    panic(err)
}

humanTime, err := base10.Uint64ToTime(bin, now.Location())
if err != nil {
    panic(err)
}
```


## Base2 - binary

See [pkg/base2](pkg/base2) for details.

Example:
`2278081318148840948` is 2023-05-15 14:14:25.732479

```go
now := time.Now()

bin, err := base2.TimeToUint64(now)
if err != nil {
    panic(err)
}

binTime, err := base2.Uint64ToTime(bin, now.Location())
if err != nil {
    panic(err)
}
```
