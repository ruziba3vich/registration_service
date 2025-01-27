Uy vazifa
Home Assignment: Simplified Microservices with Inter-Service Communication
Objective (Maqsad)
Ushbu topshiriqning maqsadi Go foydalangan holda ikkita microservisni ishlab chiqishdir. User service foydalanuvchi ma'lumotlarini boshqaradi va Order service yangi foydalanuvchi yaratilganda foydalanuvchi xizmatidan bildirishnomalarni boshqaradi.
Tasks (Vazifalar)
Siz ikkita mikroservis yaratasiz: User Service va Order Service. Har bir xizmat o'zining maxsus domenini boshqaradi va HTTP usullari (POST, PUT, DELETE) yordamida boshqasi bilan bog'lanadi.
Requirements (Talablar)
User va order servicelarni yaratish.
Servicelarni o'zaro aloqani yaratish
Detailed Instructions (Batafsil ko'rsatmalar)
1. User Service
Port: 9000
Endpointsregistration_ms
Create User
URL: /user
Method: POST
Request Body:
{
  "id": "string",
  "data": "string"
}
Response:
201: User created
400: User already exists or invalid request
Update User
URL: /user/:id
Method: PUT
Request Body:
{
  "data": "string"
}
Response:
200: User updated
404: User not found
400: Invalid request
Delete User
URL: /user/:id
Method: DELETE
Response:
200: User deleted
404: User not found
Notification (Bildirishnoma)
User yaratilishi haqida Order service ga xabar bering.
URL: http://localhost:9001/order/notify
Method: POST
Request Body:
{
  "user_id": "string",
  "action": "created"
}
