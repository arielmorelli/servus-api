import requests

# Create data
payload = dict(
    route="a/b/c",
    methods=["get", "post"],
    response_code=200,
    headers={"H": "A"},
    parameters={"ppp": "aaa"},
    match_headers=True,
    response={"data": "111111111"},
)
res = requests.post("http://127.0.0.1:8080/_register/", json=payload)
print(res)

payload = dict(
    route="a/b/c",
    methods=["get", "post"],
    response_code=200,
    headers={"H": "b"},
    parameters={"ppp": "aaa"},
    match_headers=True,
    response={"data": "aaaaaaaa"},
)
res = requests.post("http://127.0.0.1:8080/_register/", json=payload)
print(res)

# Match data
res = requests.get("http://127.0.0.1:8080/a/b/c/", params={"ppp": "aaa"}, headers={"H": "A"})
print(res, res.json())

res = requests.get("http://127.0.0.1:8080/a/b/c/", params={"ppp": "aaa"}, headers={"H": "b"})
print(res, res.json())
