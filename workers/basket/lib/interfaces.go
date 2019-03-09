package basket

//mmm...autonomous vs  library coupling ("dont share model")
import (
	"fmt"
	"teamworkers/workers/crud"
	pb "teamworkers/workers/product/lib/proto"
	"time"
)

func GetProductFromCatalog(ProductId string) (*pb.Product, error) {
	t := time.Now()
	p := pb.Product{}
	if err := crud.GetEntityFromService("product", ProductId, &p); err != nil {
		return nil, err
	} else {
		fmt.Printf("getting product %s took %s\n", ProductId, time.Now().Sub(t))
		return &p, nil
	}
}

//other solutions
/*
func GetProductLocally(ProductId string) (*pb.Product, error, bool) {
	if f := crud.GetLocalGetter("product"); f != nil {
		p := f(ProductId)
		ret, ok := p.(*pb.Product)
		//fmt.Printf("ok: %v, ret %v", ok, ret)
		if ok {
			return ret, nil, true
		} else {
			return nil, fmt.Errorf("crud registration type error for product entity "), true
		}
	}
	return nil, nil, false
}

func GetProductFromCatalog(ProductId string) (*pb.Product, error) {
	//if p, err, isLocal := GetProductLocallyA(ProductId); isLocal {
	//	return p, err
	//}

	p := pb.Product{}
	if err := crud.GetEntityFromService("product", ProductId, &p); err != nil {
		return nil, err
	} else {
		return &p, nil
	}
}


func GetProductLocallyA(ProductId string) (*pb.Product, error, bool) {
	fmt.Println("x")
	if f := crud.GetLocalGetter("product"); f != nil {
		getter := crud.MakeFunction(f).(func(string) *pb.Product)
		p:= getter(ProductId)
		return p,nil,true
	}
	return nil,nil,false
}

func GetProductLocallyA(ProductId string) (*pb.Product, error, bool) {
	fmt.Println("x")
	if f := crud.GetLocalGetter("product"); f != nil {
		getter := crud.MakeFunction(f).(func(string) *pb.Product)
		p:= getter(ProductId)
		return p,nil,true
	}
	return nil,nil,false
}
*/
