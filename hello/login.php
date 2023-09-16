<?php

// Database connection details from signup.php

$host = "svc.sel5.cloudtype.app:30279";
$db   = "alexandria";
$user = "root";
$pass = "camel99";

$conn = new mysqli($host, $user, $pass, $db);

if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}

if (!isset($_POST['username']) || !isset($_POST['password'])) {
    die("Error: Required field is missing");
}

$username = $_POST['username'];
$password = $_POST['password'];
$password = password_hash($_POST['password'], PASSWORD_DEFAULT);
$sql = "SELECT pw FROM accounts WHERE userName=?";
$stmt = $conn->prepare($sql);
$stmt->bind_param("s", $username);

$stmt->execute();
$result = $stmt->get_result();
$row = $result->fetch_assoc();

if ($row && password_verify($password, $row['pw'])) {
    echo 'Success';
} else {
    echo 'Invalid username or password';
}

$conn->close();

?>