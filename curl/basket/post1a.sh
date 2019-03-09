echo
echo \ "POST basket: X3"

curl -X POST \
  http://localhost:9001/basket \
  -d '{"productId":"1A","quantity":1.2}
'
./getall.sh
