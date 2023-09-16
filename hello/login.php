<?php

$host = 'svc.sel5.cloudtype.app';
$db   = 'alexandria';
$user = 'root';
$pass = 'camel99';

$conn = new mysqli($host, $user, $pass, $db);

if ($conn->connect_error) {
  die('Connection failed: '.$conn->connect_error);
}

if (!isset($_POST['username']) || !isset($_POST['password'])) {
  die('Error: Required field is missing');
}

$username = $_POST['username'];
$password = hash('sha256', $_POST['password']); // Hash the password using SHA256

$sql = "SELECT pw FROM accounts WHERE userName=?";
$stmt = $conn->prepare($sql);
$stmt->bind_param("s", $username);

$stmt->execute();
$result = $stmt->get_result();
$row = $result->fetch_assoc();

if ($row && strcmp($row['pw'], $password) === 0) { // Compare the hashed password with the one in the database
    echo 'Success';
} else {
    echo 'Invalid username or password';
}

$conn->close();

?>