echo
echo \ "POST basket: X3"

curl -X POST \
  http://localhost:9001/basket \
  -d '{"productId":"2B","quantity":5}
'
./getall.sh
