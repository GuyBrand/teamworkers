echo
echo \ "PUT product: X3 - price = 999"

curl -X PUT \
  http://localhost:9000/product \
  -d '{"id":"X3","description":"Tuna","price":999
}
'
./getall.sh
