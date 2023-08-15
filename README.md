# Wope Case
Bu uygulama GoFiber, Gorm ve PostgreSQL kullanarak basit bir web uygulaması örneğidir.

Docker imajı multi-stage build ile oluşturuldu, ilk aşamada kaynak kodlar kopyalanıp uygulama build ediliyor, ikinci aşamada ise binary dosyası ve uygulamanın çalışması için gerekli kütüphaneler kopyalanıp yeni bir imaj oluşturuluyor.

## Kullanım
Uygulamayı çalıştırmak için, projenin ana dizininde `docker compose up` komutunu çalıştırmanız yeterli olacaktır. 2 adet container oluşturulacak, biri uygulama için diğeri ise veritabanı için. Uygulama 8080 portunda çalışacaktır.

## Helm Chart
Uygulama için bir helm chart oluşturdum. Chart' ı deploy etmek için, projenin ana dizininde `helm install wope-case ./helm` komutunu çalıştırmanız yeterli olacaktır. Chart, Cluster içerisinde bir deployment ve service oluşturacaktır. Deployment' ın replica sayısı 3 olarak ayarlandı. Uygulama 8080 portunda çalışacaktır.

### Endpoints
| Endpoint       | Açıklama                   |
|----------------|----------------------------|
| `GET /`        | Tüm taskları döndürür.     |
| `GET /:id`     | Bir taskı döndürür.        |
| `POST /`       | Yeni bir task oluşturur.   |
| `PATCH /:id`   | Bir taskı günceller.       |
| `DELETE /:id`  | Bir taskı siler.           |

### Aldığım Notlar
Uygulamanın daha iyi olması için;
- bir front end yapılabilir
- bir authentication sistemi eklenebilir
- pre-commit hookları ile test ve lint işlemleri yapılabilir

Helm chart yazarken çok fazla değişken kullanmayı denedim ve bir noktada yönetimi zor bir hale geldi. Bü yüzden daha az değişken kullanarak bazı değerleri statik olarak ekledim.

> Terraform ile bir AWS EKS cluster oluşturup uygulamayı orada kaldırabilirim. Uygulamanın güncellenmesi için, main branch' e kod push'landığında bir docker imajı oluşturulabilir, Deployment objesinde bir secret tanımlayarak imajın güncel versiyonunu alabilir. Bu şekilde uygulama sürekli güncel kalır.

> Uygulama için bash script yerine go ile bir executable yazılabilir.