<?php
  http_response_code(418);
  header("X-Some-Response-Header: some-value");

  $headers = array();
  foreach (getallheaders() as $name => $value) {
    $headers[$name] = $value;
  }
  echo json_encode(
    array(
      "request" => $_REQUEST,
      "headers" => $headers,
    )
  );
?>