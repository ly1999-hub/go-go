# go-go
Dự án cung cấp API cho chức năng:
- Admin quản lý chuỗi các nhà hàng và người dùng dựa vào quyền được cấp bởi tài khoản root.
- Các nhà hàng có thể tự quản lý các sản phẩn cung cấp.
- người dùng có thể xem danh sách các nhà hàng, lựa chọn món và đặt bàn.
- xác thực và phân quyền cho tất cả các đối tượng sử dụng hệ thống.

## Công Nghệ.
1.Dự án sử dụng framework echo Golang xây dựng hệ thống back-end.
  - Echo là một framework tập trung vào hiệu suất, tính mở rộng và đơn giản.
  - Echo cung cấp nhiều tính năng như ràng buộc và hiển thị dữ liệu, hỗ trợ TLS tự động, tính mở rộng, hỗ trợ   HTTP/2, middlewares, tối ưu hóa router, và các tính năng khác.
  ![echo](https://echo.labstack.com/assets/images/terminal-6c7d48f3f4012a643e3e5a49b23e9308.png)
  [echo golang](https://echo.labstack.com/) 

2.MongoDB cho việc lưu trữ và thao tác với dữ liệu.
  - Ưu điểm đầu tiên của MongoDB chính là sử dụng lưu trữ dữ liệu dưới dạng Document JSON.
  - Sử dụng MongoDB các bạn có thể mở rộng dễ dàng hơn.
  - Có hiệu suất cao.
  ![mongoDB](https://www.mongodb.com/fr-fr)
  [MongoDB](https://intech.vietnamworks.com/media/gallery/2023/04/10/6433c077d1c64.jpg)

3. Firebase.
  - Realtime Database là một cơ sở dữ liệu thời gian thực.
  - Cloud Storage là tính năng cho phép lưu trữ và quản lý nội dung đã tạo ra như ảnh, video, nội dung, văn bản,...
    ![Firebase](https://teky.edu.vn/blog/wp-content/uploads/2021/10/Google-Firebase-la-gi.jpg)
    [Firebase](https://firebase.google.com/)
5. Sendgrid.
  - dịch vụ gửi email thông qua sendGrid.
    ![Sendgrid](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTsLMxcU_YbGPfbHcvF_PszxZ21AWIoDe3GgWpCw24lUA&s)
    [Sendgrid](https://sendgrid.com/en-us)
    

