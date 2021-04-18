# go-email-verification
this simple app that mimic email verification process
# usage
-  email registration link : http://localhost:8000/mail/{yourmail}
-  it will print {"verified": false}
- it will send email to the  registered {yourmail} that contain verification link
- visit the verification link and then refresh  http://localhost:8000/mail/{yourmail} and it will print {"verified": true}