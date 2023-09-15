<?php
$host = "svc.sel5.cloudtype.app:30279";
$db   = "alexandria";
$user = "root";
$pass = "camel99";

// Create connection
$conn = new mysqli($host, $user, $pass, $db);

// Check connection
if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}

$username = $_POST['username'];
$password = $_POST['password'];
$phoneNum = $_POST['phoneNum'];
$schoolNum = $_POST['schoolNum'];

$sql = "INSERT INTO accounts (username, pw, phoneNum, schoolNum) VALUES (?, ?, ?, ?)";

$stmt = $conn->prepare($sql);
$stmt->bind_param("ssss", $username, $password, $phoneNum, $schoolNum);

if ($stmt->execute()) {
    echo "New record created successfully";
    header('Location: https://port-9000-coding-new-world-30yyr422almhdj93d.sel5.cloudtype.app/board/main.html');
} else {
    echo "Error: " . $sql . "<br>" . $conn->error;
}

$conn->close();
?>