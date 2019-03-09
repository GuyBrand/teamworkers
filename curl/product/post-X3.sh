echo
echo \ "POST product: X3"

curl -X POST \
  http://localhost:9000/product \
  -d '{"id":"X3","description":"Tuna","price":6.78
      }'
./getall.sh
