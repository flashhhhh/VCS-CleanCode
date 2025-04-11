# VCS-CleanCode
## SOLID
### 1. Single responsibility principle
Một class chỉ nên giữ 1 trách nhiệm duy nhất, và chỉ có duy nhất một lý do để thay đổi.

Điều này giúp tăng khả năng bảo trì code và giảm rủi ro phá vỡ chức năng của các hàm khi sửa đổi một class.

### 2. Open/ Close principle
Có thể thoải mái mở rộng 1 class, nhưng không được sửa đổi bên trong class đó

Giúp cho tính năng mới có thể được thêm vào mà không phải sửa đổi code cũ, giảm thiểu lỗi trong các thành phần đã được cố định.

### 3. Liskov substitution principle
Các object của class con có thể thay thế class cha mà không làm thay đổi tính đúng đắn của chương trình.

Tránh việc có những phương thức hoạt động không đúng khi sử dụng class con.

### 4. Interface segregation principle
Mỗi class không bị bắt buộc phải implement cho interface mà nó không sử dụng.

Tránh việc các interface lớn, cồng kềnh chèn các chức năng không liên quan vào một class.

### 5. Dependency inversion principle
Các module cấp cao không nên phụ thuộc vào các module cấp thấp. Thay vào đó, cả 2 chỉ nên phụ thuộc vào các lớp trừu tượng.

Điều này giúp tăng tính linh hoạt bằng cách giảm sự phụ thuộc giữa các lớp được triển khai.

## Dependency Injection
### 1. Tổng quan và mục đích sử dụng
Dependency Injection một kỹ thuật thiết kế phần mềm trong lập trình hướng đối tượng, dùng để giảm thiểu sự phụ thuộc giữa các lớp (class) với nhau, giúp code dễ hiểu, dễ bảo trì và nâng cấp hơn.

Thay vì tự tạo các lớp phụ thuộc vào mình, thì các lớp sẽ nhận các đối tượng từ việc được các lớp khác **inject** vào từ bên ngoài. Điều này giúp đạt được các mục đích sau:
* **Tách biệt mối quan tâm**: Lớp chỉ cần tập trung vào logic riêng của mình, không cần lo việc tạo ra các dependency, khiến code rõ ràng và dễ quản lý hơn.
* **Liên kết lỏng lẻo**: Lớp phụ thuộc ít vào cách thức tạo ra các dependency, dễ dàng thay thế các dependency khác nhau trong các môi trường khác nhau, ví dụ: testing, production…
* **Tính linh hoạt cao**: Kiểm soát tốt hơn cách tạo và quản lý các dependency, thuận tiện cho việc mở rộng và tùy chỉnh hệ thống.

### 2. Phân loại Dependency Injection
#### Constructor Injection
Dependency được truyền thẳng vào thông qua hàm khởi tạo. Cách tiếp cận này được sử dụng nhiều do đảm bảo được dependency luôn tồn tại ngay khi lớp đó được khởi tạo.

Ví dụ:
```go
type Notifier interface {
	Send(msg string) error
}

type EmailNotifier struct {
}

func (e *EmailNotifier) Send(msg string) error {
	fmt.Println("Email: ", email)
	return nil
}

type NotificationService struct {
	notifier Notifier
}

func NewNotificationService(n Notifier) *NotificationService {
	return &NotificationService{notifier: n}
}
```

#### Setter Injection
Dependency được set sau khi khởi tạo object thông qua 1 hàm setter. Điều này giúp cho việc khởi tạo dependency chỉ xảy ra khi cần thiết.

Ví dụ:
```go
/*
	Declare Notifier and EmailNotifier
*/

type NotificationService struct {
	notifier Notifier
}

func (ns *NotificationService) SetNotifier(n Notifier) {
	ns.notifier = n
}
```

#### Method Injection
Dependency được truyền trực tiếp vào như 1 tham số của hàm. Dependency Injection này thường đươc sử dụng khi chỉ có một vài phương thức đặc biệt cần dependency đó.

Ví dụ:
```go
/*
	Declare Notifier and EmailNotifier
*/

type NotificationService struct {
}

func (ns *NotificationService) Send(e *EmailService) {
	e.Send("Hello World!")
}
```

### 3. Ưu điểm
* Giảm sự phụ thuộc giữa các thành phần.
* Dễ dàng tái sử dụng và sửa đổi các thành phần.
* Unit test trở nên dễ dàng hơn với các phụ thuộc giả lập.
* Dễ dàng cập nhật hoặc sửa đổi chức năng.
* Tách biệt mỗi thành phần, giúp cho mỗi thành phần chỉ phải quan tâm đúng một chức năng.

### 4. Nhược điểm
* Tăng tính phức tạp của hệ thống, do tốn thời gian khởi tạo cấu trúc file.
* Việc debug trở nên khó khăn hơn nhiều do các phụ thuộc đều chỉ biết được khi trong runtime.
* Chỉ phù hợp với các dự án lớn, phức tạp đối với các dự án nhỏ.

## Design Pattern
### Định nghĩa
Design pattern là một giải pháp tiêu biểu cho các vấn đề phổ biến trong thiết kế phần mềm. Nó giống như một bản thiết kế sẵn mà bạn có thể tham khảo để giải quyết một vấn đề đang xảy ra cho code của bạn. Mỗi pattern không phải đoạn code hoàn chỉnh, nhưng nó chứa một ý tưởng để giải quyết một vấn đề cụ thể.

###
* **Tăng khả năng tái sử dụng code:** Tránh việc viết lại logic code nhiều lần.
* **Tăng khả năng bảo trì code**: Làm quá trình debug và mở rộng tính năng dễ dàng hơn.
* **Tăng khả năng hợp tác**: Hiểu được code giúp cho lập trình viên có thể teamwork hiệu quả hơn.

### Phân loại design pattern
* **Creational pattern:** Cung cấp các cơ chế khởi tạo đối tượng giúp tăng tính linh hoạt và khả năng tái sử dụng code.
* **Structure pattern:** Cung cấp cho việc kết hợp các lớp và đối tượng lên một cấu trúc lớn hơn, trong khi vẫn giữ cho cấu trúc đó linh hoạt và hiệu quả.
* **Behavior pattern:** Đảm nhận việc giao tiếp hiệu quả và phân phối trách nhiệm giữa các đối tượng.

### Một số pattern ví dụ
#### 1. Singleton
Singleton là một mẫu thiết kế khởi tạo giúp đảm bảo rằng 1 class chỉ có thể khởi tạo duy nhất 1 lần, trong khi class đó vẫn có thể được gọi từ bất cứ đâu.

Các vấn đề Singleton giải quyết:
* Đảm bảo mỗi class chỉ có đúng 1 instance: Sau khi khởi tạo 1 instance, lần khởi tao tiếp theo vẫn sẽ là instance đó mà không phải 1 instance mới.
* Cung cấp một điểm truy cập toàn cục đến instance đó: Người dùng có thể lấy được thông tin từ instance đó, nhưng không thể sửa đổi instance đó.

Giải pháp của Singleton:
* Hàm khởi tạo của class phải là private, để tránh các thực thể khác có thể khởi tạo 1 class mới hoàn toàn.
* Tạo một hàm static hoạt động như 1 constructor, hàm này sẽ gọi đến private constructor để khởi tạo class mới, hoặc trả về class đã được tạo từ trước.

**Mẫu thiết kế:**

Ứng dụng:
* Khi có một class nên có duy nhất 1 instance của nó (ví dụ như 1 kết nối database dùng chung cho mọi class)
* Cần có quyền kiểm soát chặt chẽ hơn cho các biến toàn cục (Không có class nào ngoài class đó có thể sửa đổi instance của nó)

Ưu điểm:
* Đảm bảo mỗi class chỉ có duy nhất một instance.
* Đạt được truy cập toàn cục tới instance đó.
* Instance chỉ được khởi tao duy nhất ở lần đầu tiên.

Nhược điểm:
* Vi phạm quy tắc **Single Responsibility**: Singleton giải quyết tận 2 vấn đề cùng một lúc.
* Nếu quá nhiều lời gọi khởi tạo cùng lúc, phải có cơ chế Lock để chỉ có phép duy nhất 1 khởi tạo.

#### 2. Abstract Factory
Abstract Factory là mẫu thiết kế khởi tạo cho phép tạo ra các họ đối tượng có liên quan mà không cần chỉ định các lớp cụ thể của chúng.

Vấn đề giải quyết:
* Tạo các họ liên quan đến đối tượng đó mà không phải trộn lẫn chúng. Ví dụ: Khi tương tác với nhiều database hay cache, ta có thể kết hợp giữa 2 thành phần không tương thích với nhau (ví dụ UserServiceDatabase và OrderServiceCache).
* Tránh việc liên kết chặt chẽ với các lớp khác: Khi 1 lớp cần sửa đổi như thêm bớt thuộc tính, sẽ làm ảnh hưởng đến tất cả các hàm gọi đến nó.
* Lặp lại logic code ở nhiều nơi: Khi không có cơ chế khởi tạo tập trung, việc khởi tạo object sẽ bị lặp lại và dễ gây ra sự không đồng nhất.
* Khó khăn trong việc quay trở lại họ của các lớp đó (VD: UserServiceDatabase và Database là 2 lớp riêng biệt)

Giải pháp:
* Định nghĩa các interface cho các họ đối tượng đó (VD: database, cache, objectstorage), sau đó tạo ra tất cả các đối tượng liên quan khác implement các interface chung này. VD: Tất cả các database của service khác như UserServiceDatabase hay OrderServiceDatabase đều implement được interface Database.
* Tiếp theo, khởi tạo lớp AbstractFactory với một loạt các hàm khởi tạo cho các lớp nằm trong họ đối tượng trên (VD: createDatabase(), createCache(), ...). Các hàm này bắt buộc phải trả về lớp abstract (như Database, Cache, ...).

**Mẫu thiết kế:**

Ứng dụng:
* Sử dụng Abstract Factory khi cần hoạt động với nhiều họ sản phẩm liên quan khác nhau, nhưng không muốn nó phụ thuộc vào các lớp cụ thể của những sản phẩm đó - có thể chúng chưa được biết trước hoặc muốn mở rộng thêm trong tương lai.
* Có thể cài đặt Abstract Factory khi có một lớp chứa một số hàm khởi tạo Factory, giúp nó ẩn đi trách nhiệm chính của lớp đó.

Ưu điểm:
* Đảm bảo được các lớp sinh ra từ Abstract Factory tương thích với nhau.
* Tránh sự phụ thuộc giữa các lớp với client.
* Đảm bảo được nguyên tắc **Single Responsibility** và **Open/Close**, do hoàn toàn có thể thêm 1 lớp khác mà không cần phải sửa code cũ.

Nhược điểm:
* Làm cho code trở nên phức tạp hơn, do nhiều interfaces và lớp khác được thêm vào.

#### Proxy
Proxy là một mẫu thiết kế cấu trúc cho phép tạo ra một bản sao cho một đối tượng khác. Proxy có nhiệm vụ kiểm soát truy cập đến đối tượng gốc, cho phép thực hiện một số thay đổi trước và sau khi request đó đến được đối tượng gốc.

Vấn đề giải quyết:
* Khi đối tượng gốc quá tiêu tốn tài nguyên để khởi tạo hoặc sử dụng: Ví dụ như khi cuộn xuống từ youtube thì các video mới dần dần được load.
* Cần thêm các hành động (như cache, authorization, ...) mà không cần động vào thực thể gốc.
* Cần hạn chế hoặc kiểm soát lượng truy cập.

Giải pháp:
* Cài đặt một lớp Proxy có interface giống hệt như đối tượng gốc.
* Trong Proxy có một tham số đại diện cho chính đối tượng gốc.
* Khi có một request đến Proxy, tiến hành chỉnh sửa request đó trước khi bắt đầu khởi tạo đối tượng gốc rồi truyền request đó vào.

**Mẫu thiết kế:**

Ứng dụng:
* Khởi tạo lười (Virtual Proxy): Thay vì khởi tạo một đối tượng tốn rất nhiều tài nguyên nhưng không được sử dụng, ta chỉ bắt đầu khởi tạo khi cần thiết.
* Kiểm soát truy cập (Protection Proxy): Khi ta muốn chỉ một vài client được chỉ định có thể truy cập vào đối tượng gốc đó.
* Xử lý cục bộ cho một dịch vụ từ xa (Remote Proxy): Proxy sẽ được xử lý trên local và tiến hành gửi nhận và xử lý các vấn đề về network.
* Ghi log lại các request (Logging Proxy): Proxy sẽ ghi lại lịch sử của các request đến service đó.
* Cache reponse (Caching Proxy): Proxy sẽ cache lại các request và response tương ứng trước đó, khi request sau đến sẽ tiến hành kiểm tra cache trước.
* Kiểm soát đối tượng gốc: Sau một khoảng thời gian không có request nào đến sẽ tiến hành ngắt đối tượng gốc để giải phóng tài nguyên.

Ưu điểm:
* Có thể kiểm soát đối tượng gốc mà client không biết về nó.
* Kiểm soát vòng đời của đối tượng gốc.
* Proxy có thể hoạt động kể cả khi đối tượng gốc chưa được khởi tạo.
* Đảm bảo nguyên tắc **Open/Close** do có thể thêm xóa sửa mà không ảnh hưởng đến code.

Nhược điểm:
* Code trở nên phức tạp hơn do phải thêm nhiều lớp và hàm.
* Mất thêm thời gian phản hồi cho request do phải thông qua proxy.