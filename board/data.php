<?php

$pageNumber = $_GET['page']; 

$itemsPerPage=10;

$startFrom=($pageNumber-1)*$itemsPerPage;

$data=array();

for($i=$startFrom+1;$i<=$startFrom+$itemsPerPage;$i++){
   $data[]=$i;
}

echo json_encode($data);

?>