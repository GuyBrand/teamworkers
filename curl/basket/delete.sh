echo
echo \ "DELETE basket: X3"

curl -X DELETE \
  http://localhost:9001/basket/X3

./getall.sh
