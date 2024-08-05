package main

import (
	"encoding/json"
	"fmt"
	pb "github.com/phthaocse/protobuf_json/proto"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

func TestJsonMarshal(data any) []byte {
	start := time.Now()
	out, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("Failed to encode json:", err)
	}
	elapsedTime := time.Since(start)
	fmt.Printf("elapsed time marshal json %dns with size %d bytes \n", elapsedTime.Nanoseconds(), len(out))
	return out
}

func TestProtoMarshal(data proto.Message) []byte {
	start := time.Now()
	out, err := proto.Marshal(data)
	if err != nil {
		log.Fatalln("Failed to encode proto:", err)
	}
	elapsedTime := time.Since(start)
	fmt.Printf("elapsed time marshal proto %dns with size %d bytes \n", elapsedTime.Nanoseconds(), len(out))
	return out
}

func TestJsonUnmarshal(data []byte, v any) {
	start := time.Now()
	err := json.Unmarshal(data, v)
	if err != nil {
		fmt.Println("Failed to decode json:", err)
	}
	elapsedTime := time.Since(start)
	fmt.Printf("elapsed time unmarshal json %dns \n", elapsedTime.Nanoseconds())
}

func TestProtoUnmarshal(data []byte, v proto.Message) {
	start := time.Now()
	err := proto.Unmarshal(data, v)
	if err != nil {
		fmt.Println("Failed to decode json:", err)
	}
	elapsedTime := time.Since(start)
	fmt.Printf("elapsed time unmarshal proto %dns \n", elapsedTime.Nanoseconds())
}

func main() {
	t := &pb.TestVarInt{
		//Id:     50000,
		Number: 1000000000,
		//Str:  "chỉ là test thôi mà",
		//Str2: "Fixed-Size Numeric Types: Protocol Buffers support fixed-size numeric types such as fixed32 and fixed64. These types use a fixed number of bytes to represent the value, regardless of the magnitude of the number. For example, a fixed32 field always occupies 4 bytes, and a fixed64 field always occupies 8 bytes.",
	}
	// JSON data
	jsonData := `{
        "_id": "66b077ecfa4afeeb66be600f",
        "shop_id": 53579,
        "client_id": 9,
        "return_name": "0949900024",
        "return_phone": "0949900024",
        "return_address": "180 Đ. Nguyễn Thị Minh Khai, Phường 2, Quận 3, Thành phố Hồ Chí Minh",
        "return_ward_code": "20302",
        "return_district_id": 1444,
        "return_location": {},
        "from_name": "0949900024",
        "from_phone": "0949900024",
        "from_hotline": "",
        "from_address": "180 Đ. Nguyễn Thị Minh Khai, Phường 2, Quận 3, Thành phố Hồ Chí Minh",
        "from_ward_code": "20302",
        "from_district_id": 1444,
        "from_location": {},
        "deliver_station_id": 0,
        "to_name": "Liêu Sao Linh",
        "to_phone": "0909000002",
        "to_address": "7/28, Thành Thái, Phường 14, Quận 10, Hồ Chí Minh",
        "to_ward_code": "21014",
        "to_district_id": 1452,
        "to_location": {
            "lat": 10.7694164,
            "long": 106.6641113,
            "cell_code": "AJIAENA7",
            "place_id": "ChIJo2NkAegudTERcMqLszDgRVM",
            "trust_level": 5,
            "wardcode": "21014"
        },
        "weight": 40000,
        "length": 10,
        "width": 10,
        "height": 10,
        "converted_weight": 200,
        "calculate_weight": 40000,
        "image_ids": "",
        "service_type_id": 5,
        "service_id": 100039,
        "payment_type_id": 2,
        "payment_type_ids": [
            2
        ],
        "custom_service_fee": 0,
        "sort_code": "000-C-00-00",
        "cod_amount": 50000,
        "is_cod_transferred": false,
        "is_cod_collected": false,
        "insurance_value": 100000,
        "order_value": 0,
        "pick_station_id": 0,
        "client_order_code": "",
        "cod_failed_amount": 0,
        "cod_failed_collect_date": "",
        "required_note": "KHONGCHOXEMHANG",
        "content": "tuan [1 kiện]",
        "note": "",
        "employee_note": "",
        "seal_code": "",
        "pickup_time": "2024-08-05T06:57:47.686Z",
        "items": [{
            "name": "tuan",
            "code": "",
            "quantity": 1,
            "price": 0,
            "length": 10,
            "width": 10,
            "height": 10,
            "category": {
                "level1": "",
                "level2": "",
                "level3": ""
            },
            "weight": 40000,
            "image_ids": [],
            "status": "",
            "item_order_code": "L6A8Q9_1"
        }],
        "coupon": "",
        "coupon_campaign_id": 0,
        "order_code": "L6A8Q9",
        "version_no": "e40d0e6b-9fea-4365-bf39-a0965eba315e",
        "updated_ip": "127.0.0.1",
        "updated_employee": 0,
        "updated_client": 9,
        "updated_source": "online--stream-corev2.ShippingOrder--UpdateOrder",
        "updated_date": "2024-08-05T06:57:48.663Z",
        "updated_warehouse": 0,
        "created_ip": "127.0.0.1",
        "created_employee": 0,
        "created_client": 9,
        "created_source": "",
        "created_date": "2024-08-05T06:57:40.120Z",
        "status": "ready_to_pick",
        "internal_process": {
            "status": "",
            "type": ""
        },
        "pick_warehouse_id": 20751002,
        "deliver_warehouse_id": 20751000,
        "current_warehouse_id": 20751002,
        "return_warehouse_id": 20751002,
        "next_warehouse_id": 0,
        "current_transport_warehouse_id": 0,
        "leadtime": "2024-08-07T23:59:59.000Z",
        "order_date": "2024-08-05T06:57:47.686Z",
        "data": {
            "ahamove_seller_address_detail": "180 Đ. Nguyễn Thị Minh Khai, Phường 2, Quận 3, Thành phố Hồ Chí Minh, Phường 2, Quận 3, Hồ Chí Minh",
            "ahamove_recipient_address_detail": "7/28, Thành Thái, Phường 14, Quận 10, Hồ Chí Minh, Phường 14, Quận 10, Hồ Chí Minh",
            "province_id": 202,
            "ahamove_order_code": "24WC65PG"
        },
        "action": "",
        "soc_id": "66b077ecfa4afeeb66be600e",
        "tag": [
            "bulky",
            "truck"
        ],
        "legacy": false,
        "is_partial_return": false,
        "is_document_return": false,
        "-": {},
        "updated_date_pick_shift": "",
        "updated_date_deliver_shift": "",
        "updated_date_return_shift": "",
        "transaction_ids": [
            "1e36a2a6-59f6-4d62-8cb7-bf9bc834c536"
        ],
        "transportation_status": "",
        "transportation_phase": "",
        "extra_service": {
            "document_return": {
                "flag": false
            },
            "double_check": false
        },
        "config_fee_id": "663c886196f0eb6f57ee884c",
        "extra_cost_id": "668ba9a47c8d5f119036bf47",
        "standard_config_fee_id": "663c886196f0eb6f57ee884c",
        "standard_extra_cost_id": "668ba9a47c8d5f119036bf47",
        "operation_partner": "ahamove",
        "type_order": "freight",
        "type_order_code": "FSME",
        "updated_time": "2024-08-05T06:57:49.197Z"
    }`

	// Parse JSON data
	order := new(pb.ShippingOrder)
	err := json.Unmarshal([]byte(jsonData), order)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	fmt.Println("test marshal")
	protoByte := TestProtoMarshal(order)
	jsonByte := TestJsonMarshal(order)
	fmt.Println("test unmarshal")
	TestProtoUnmarshal(protoByte, t)
	TestJsonUnmarshal(jsonByte, t)
}
