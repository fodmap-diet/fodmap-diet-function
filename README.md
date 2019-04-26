# fodmap-diet-function
openfaas function for fodmap diet   
live at [o6s.io](https://s8sg.o6s.io/fodmap-diet)

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
