echo
echo \ "DELETE product: X3"

curl -X DELETE \
  http://localhost:9000/product/X3

./getall.sh
