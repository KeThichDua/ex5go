## bài tập sử dụng grpc và http, redis
- các bạn đã tới bài này. Sau bài này chúng ta có thể làm đc việc rồi
-  đây là bài sẽ tổng hợp toàn bộ kiến thức mà các bạn đã có và đã học được từ những bài trước.
### Chuẩn bị kiến thức:
<!-- - Đọc kỹ bài sau:[chương 3](https://zalopay-oss.github.io/go-advanced/ch3-rpc/ch3-01-rpc-go.html) -->
- Tài liệu [grpc concepts](https://grpc.io/docs/what-is-grpc/core-concepts/)
- Xem video bài học dưới đây. [TECH SCHOOL](https://www.youtube.com/watch?v=2Sm_O75I7H0&list=PLy_6D98if3UJd5hxWNfAqKMr15HZqFnqf&index=1)


###  Cài đặt
- Tải xuống bản build sẵn của protobuf complier  [url](https://github.com/protocolbuffers/protobuf/releases). Tải xuống bản phù hợp, và đưa vào thư mục /usr/local/

- sau đó thực hiện lệnh `which protoc`
- đọc bài [protoc gogo](https://github.com/gogo/protobuf)
- chay cụm lệnh này.
```bash
go get github.com/gogo/protobuf/protoc-gen-gofast \
        github.com/gogo/protobuf/protoc-gen-gogofaster

go install github.com/gogo/protobuf/protoc-gen-gofast \
        github.com/gogo/protobuf/protoc-gen-gogofaster
```

1. bài 1: Viết một `message` grpc sao cho khi generate ta được struct sau.

```go
type UserPartner struct {
	Id string `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PartnerId string `protobuf:"bytes,3,opt,name=partner_id,json=partnerId,proto3" json:"partner_id,omitempty"`
	AliasUserId string           `protobuf:"bytes,5,opt,name=alias_user_id,json=aliasUserId,proto3" json:"alias_user_id,omitempty"`
	Apps        map[string]int64 `protobuf:"bytes,6,rep,name=apps,proto3" json:"apps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Phone       string           `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Created     int64            `protobuf:"varint,10,opt,name=created,proto3" json:"created,omitempty"`
	UpdatedAt   int64            `protobuf:"varint,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}
```

2. bài 2: Viết một message `UserPartnerRequest` nhằm tạo 1 query `xorm`. Bao gồm lấy userpartner theo `user_id`, `phone`, với `limit` là số lượng row lớn nhất được quét ra. Với id được genere ngẫu nhiên với [xid](https://github.com/rs/xid)

3. bài 3: Sử dụng kiến thức đã đọc tao 1 server `net/http` hoặc `mux` sử dụng port `3001` mở 1 url. Viết 1 route `/` trả về `pong`

4. bài 4: kết hợp với kiến thức trên viết 1 server sử dụng `grpc` generate 2 message ở bài 1, kết hợp server bài 2 và kiến thức từ bài trước `xorm`.
 Viết 1 route: POST `/user-partner` tạo mới 1 partner, GET `/user-partner` lấy danh sách partner, DELETE `/user-partner/{id}` theo id,  GET `/user-partner/{id}` lấy theo 1 id cụ thể.

 5. bài 5: Tài liệu [grpc](https://grpc.io/docs/what-is-grpc/core-concepts/)
 Tạo 1 service gen code. Tạo 1 grpc server với message `UserPartner`. Nhằm getlist, create, update
 Tạo 1 grpc client để thực hiện

6. Sử dụng thư viện [go-redis](https://github.com/go-redis/redis). Tạo 1 chương trình
thêm một key `demo_key` có giá trị là `time.Now().Unix()` với thời gian expire 10s
- Sleep chương trình 12s đọc key `demo_key` ra nếu ko có `key_demo` thì lưu lại với giá trị thời gian hiện tại tính theo giây.

7. tạo 1 list trong redis để lưu dữ liệu sau.

```txt
鈴木 陽一
Leslie Moreno
Emma Williams
渡辺 裕美子
Marcantonio Galuppi
Martha Davis
Kristen Turner
中津川 春香
Ashley Castillo
山田 桃子

'Elda Palumbo'
'Pacifico Giordano'
'Sig. Avide Guerra'
'Yago Amato'
'Eustachio Messina'
'Dott. Violante Lombardo'
'Sig. Alighieri Monti'
'Costanzo Costa'
'Nazzareno Barbieri'
'Max Coppola'
'سديم'
'وفائي'
'عبد الغفور'
'صاحب'
'قصيد'
```
- Lưu đúng số dòng trên vào file `name.txt` đọc file ra theo từng dòng và Lpush vào list `names`. Nếu dòng ko có dữ liệu hoặc dữ liệu là `"\s+"` thì ko lưu vào redis.

- Lpop và inra gía trị lấy được.

- Set thời gian expire cho list là 5s. sau đó thêm 1 phần tử vào bên phải list là `Dota2vn`. Sau đó in ra danh sách.


8. Sử dụng thư viện (gin)[https://github.com/gin-gonic/gin] để tạo 1 server http đơn giản.
- viết api get `/names` để lấy dữ liệu từ trong redis ra.(nhớ mở cái expire 5s của list).
- viết api post `/name` để tạo 1 name theo `struct Name {Name string}` Lpush vào redis.
- viết api get `/name/{index}` để lấy name theo vị trí trong redis.
- viết api post `/name/{index}` để update tên theo vị trí trong redis.
- viết api delete `/name/{index}` để xóa theo 1 index trong redis.

9. Viết một đoạn chương trình sử dụng net/http:
+ get dữ liệu của api sau: `https://jsonplaceholder.typicode.com/posts`. Tạo 1 struct để hứng kết quả trả về.

+ post một đối tượng từ struct trên để tạo mới dữ liệu với api: `https://jsonplaceholder.typicode.com/posts`. Log ra kết quả ra màn hình.
