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

if (!isset($_POST['username']) || !isset($_POST['password']) || !isset($_POST['phoneNum']) || !isset($_POST['schoolNum'])) {
    error_log('$_POST array: ' . print_r($_POST, true));
    die("Error: Required field is missing");
}

$username = $_POST['username'];
$password = $_POST['password'];
$phoneNum = $_POST['phoneNum'];
$schoolNum= $_POST['schoolNum'];

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