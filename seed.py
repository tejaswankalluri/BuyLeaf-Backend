import requests

url = "http://localhost:8000/"


def getToken():
    user = {
        "username": "John",
        "email": "joh@email.com",
        "password": "Tejaswan@123",
        "role": "ADMIN"
    }
    requests.post(url + "signup/", user)

    # login
    usertokenreq = requests.post(url + "signin/", user)
    usertoken = usertokenreq.json()
    print("token done")
    return usertoken["token"]


def setProduct(token, cookie):
    if token == "":
        return exit(1)

    product = {
        "product_name": "Python language",
        "product_desc": "this is Python language",
        "product_price": 600,
        # "product_images": [
        #     {
        #         "image_url": "https://google.com"
        #     },
        #     {
        #         "image_url": "https://google.com"
        #     },
        # ],
        "category": {
            "name": "Books",
            "desc": "This is books"
        },
        "inventory": {
            "quantity": 50
        }
    }
    r = requests.post(url + "product", json=product, cookies=cookie)
    if not r.ok:
        print(r.text)
        return exit(1)
    print("set products done")


token = getToken()
cookie = {
    "Authorization": token
}
setProduct(token, cookie)
