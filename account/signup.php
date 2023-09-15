<?php
  echo "DB연동 결과는?!<BR>";
$pgsql = pg_Connect("host=svc.sel5.cloudtype.app:30544 dbname=template0 user=root password=camel99") or die('접속실패');

  $row = pg_fetch_array($pgsql);
  echo "접속성공" .$row;
?>