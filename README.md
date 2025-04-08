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

### 5 Dependency inversion principle
Các module cấp cao không nên phụ thuộc vào các module cấp thấp. Thay vào đó, cả 2 chỉ nên phụ thuộc vào các lớp trừu tượng.

Điều này giúp tăng tính linh hoạt bằng cách giảm sự phụ thuộc giữa các lớp được triển khai.

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