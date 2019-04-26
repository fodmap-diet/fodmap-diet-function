# fodmap-diet-function
openfaas function for fodmap diet   

### Test live deployment
fodmap diet function is live at [fodmap-diet.o6s.io](https://fodmap-diet.o6s.io/fodmap-diet)
```
curl https://fodmap-diet.o6s.io/fodmap-diet -d '{"items":["banana","cream"]}'
```

### getting started
deploy
```bash
faas deploy -f https://raw.githubusercontent.com/fodmap-diet/fodmap-diet-function/master/stack.yml
```
test
```bash
curl http://192.168.99.100:8080/function/fodmap-diet -d '{"items":["banana","cream"]}'
```
sample output
```bash
{
    "banana": {
        "category": "fruit",
        "fodmap": "low",
        "condition": "unripe",
        "note": "ripe banana has high levels of fructans"
    },
    "cream": {
        "category": "dairy",
        "fodmap": "high"
    }
}
```
to improve performance allow write to `rootfs` and disable `read_only_fs` in `stack.yml`
```yml
    environment:
      read_only_fs: false 
```
