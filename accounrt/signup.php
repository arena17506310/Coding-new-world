<?php
// DB 연결 설정
$servername = "localhost";
$username = "사용자명";
$password = "비밀번호";
$dbname = "데이터베이스명";

$conn = new mysqli($servername, $username, $password, $dbname);
if ($conn->connect_error) {
    die("DB 연결 실패: " . $conn->connect_error);
}

// POST 요청으로부터 회원가입 폼 데이터 받기
$username = $_POST["username"];
$password = $_POST["password"];
$phone_number = $_POST["phone_number"];
$student_id = $_POST["student_id"];

// SQL 쿼리 실행하여 데이터베이스에 회원정보 저장
$sql = "INSERT INTO users (username, password, phone_number, student_id)
        VALUES ('$username', '$password', '$phone_number', '$student_id')";

if ($conn->query($sql) === TRUE) {
    echo "회원가입 성공!";
} else {
    echo "오류: " . $sql . "<br>" . $conn->error;
}

$conn->close();
?>