EnglishLearningCenter/
├── cmd/
│   └── api/
│       └── main.go         # Điểm khởi chạy (Entry point) của ứng dụng
├── config/
│   └── config.go           # Đọc và quản lý cấu hình (env, yaml)
├── internal/               # Chứa mã nguồn cốt lõi (bảo mật, không thể import từ ngoài)
│   ├── app/                # Khởi tạo kết nối (DB, Router, Server)
│   ├── delivery/           # Tầng Giao tiếp (HTTP Handlers / gRPC Server)
│   │   └── http/
│   │       ├── v1/
│   │       └── middleware/
│   ├── repository/         # Tầng Cơ sở dữ liệu (Database interactions)
│   ├── usecase/            # Tầng Nghiệp vụ (Business Logic)
│   └── model/              # Định nghĩa Structs / Entities
├── pkg/                    # Chứa mã nguồn dùng chung (có thể chia sẻ sang dự án khác)
│   └── logger/
├── api/                    # Chứa file đặc tả API (Swagger, OpenAPI, Proto)
├── deployments/            # Cấu hình triển khai (Docker, Kubernetes)
├── go.mod                  # Quản lý dependency của Go
├── go.sum
└── Makefile                # Các câu lệnh tắt để build/run dự án nhanh
