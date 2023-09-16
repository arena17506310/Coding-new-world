<?php

// Database connection details from signup.php

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
$password_inputted = hash('sha256', $_POST['password']);

$sql =
  'SELECT pw FROM accounts WHERE userName=?';

$stmt=$conn->prepare($sql);
$stmt->bind_param('s', $username);

$stmt->execute();
$result=$stmt->get_result();
$row=$result->fetch_assoc();

if ($row && hash_equals($row['pw'], $password_inputted)) { // 해싱된 비밀번호 검증
  echo 'Success';
} else {
  echo 'Invalid username or password';
}

$conn->close();

?>