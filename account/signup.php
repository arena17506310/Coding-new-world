<?php
ini_set('display_errors', 1);
error_reporting(E_ALL);
$host = "svc.sel5.cloudtype.app:30279";
$db   = "alexandria";
$user = "root";
$pass = "camel99";

// Create connection to MariaDB
$conn = new mysqli($host, $user, $pass, $db);

// Check connection
if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}
//닉네임 안쓰면 작동 안함 ㅅㄱ
if (!isset($_POST['username'])) {
    error_log('$_POST array: ' . print_r($_POST, true));
    die("Error: Required field is missing");
}
//비밀번호 안쓰면 작동 안함 ㅅㄱ
if (!issset($_POST['password'])) {
    error_log('$_POST array: ' . print_r($_POST, true));
    die("Error: Required field is missing");
}
//전화번호 안쓰면 작동 안함 ㅅㄱ
if (!issset($_POST['phoneNum'])) {
    error_log('$_POST array: ' . print_r($_POST, true));
    die("Error: Required field is missing");
}
//학번 안쓰면 작동 안함 ㅅㄱ
if (!issset($_POST['schoolNum'])) {
    error_log('$_POST array: ' . print_r($_POST, true));
    die("Error: Required field is missing");
}

$username = mysqli_real_escape_string($conn, $_POST['username']);
$password = mysqli_real_escape_string($conn, $_POST['password']);
$phoneNum = mysqli_real_escape_string($conn, $_POST['phoneNum']);
$schoolNum= mysqli_real_escape_string($conn, $_POST['schoolNum']);

$password = hash('sha256', $password);

$sql = "INSERT INTO accounts (userName, pw, phoneNum, schoolNum) VALUES (?, ?, ?, ?)";

$stmt = $conn->prepare($sql);
$stmt->bind_param("ssss", $username, $password, $phoneNum, $schoolNum);

if ($stmt->execute()) {
    echo "Success";
} else {
    echo "Error: " . $sql . "<br>" . mysqli_error($conn);
}

$conn->close();
?>