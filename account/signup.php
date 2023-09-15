<?php
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

if (!isset($_POST['username']) || !isset($_POST['password']) || !isset($_POST['phoneNum']) || !isset($_POST['schoolNum'])) {
    die("Error: Required field is missing");
}

$username = $_POST['username'];
$password = $_POST['password'];
$phoneNum = $_POST['phoneNum'];
$schoolNum= $_POST['schoolNum'];

$sql = "INSERT INTO accounts (userName, password, phoneNum, schoolNum) VALUES (?, ?, ?, ?)";

$stmt = $conn->prepare($sql);
$stmt->bind_param("ssss", $username, password_hash($password, PASSWORD_DEFAULT), $phoneNum, $schoolNum);

if ($stmt->execute()) {
    header('Location: ccount.html');
    exit;
} else {
    echo "Error: " . $sql . "<br>" . mysqli_error($conn);
}

$conn->close();
?>