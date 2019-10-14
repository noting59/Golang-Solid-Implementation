<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
  <meta http-equiv="x-ua-compatible" content="ie=edge">
  <title>Pay</title>
  <!-- Font Awesome -->
  <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.8.2/css/all.css">
  <!-- Bootstrap core CSS -->
        <link href="//netdna.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
</head>

<body class="grey lighten-3">

  <!-- Navbar -->
  <nav class="navbar fixed-top navbar-expand-lg navbar-light white scrolling-navbar">
    <div class="container">

      <!-- Brand -->
      <a class="navbar-brand waves-effect" href="https://mdbootstrap.com/docs/jquery/" target="_blank">
        <strong class="blue-text">MDB</strong>
      </a>

    </div>
  </nav>
  <!-- Navbar -->

  <!--Main layout-->
  <div>
    <iframe src="https://pay.sp-stage.us/api/v1/purchase/{{.PayForm.Token}}" width="100%" height="500"></iframe>
  </div>
  <div id="main-page" style="display: none">
    <a href="/" class="btn btn-success">Back to shop</a>
  </div>
  <!--Main layout-->

  <!--Footer-->
  <footer class="page-footer text-center font-small mt-4 wow fadeIn">

    <hr class="my-4">

    <!--Copyright-->
    <div class="footer-copyright py-3">
      © 2019 Copyright:
      <a href="https://mdbootstrap.com/education/bootstrap/" target="_blank"> MDBootstrap.com </a>
    </div>
    <!--/.Copyright-->

  </footer>
  <!--/.Footer-->

  <!-- Bootstrap core JavaScript -->
    <script src="//code.jquery.com/jquery-1.11.1.min.js"></script>

    <script src="//netdna.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>

  <script>

      function listener(event) {
        if (event.origin == "https://pay.sp-stage.us") {
            if(event.data.response.type == "orderStatus") {
                if (vent.data.response.order.status == "approved") {
                    $("#main-page").show()
                }
            }
        console.log(event)
          $.ajax({
            type: "POST",
            url: "/pay/process",
            data: JSON.stringify(event.data.response),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            cache: false,
            success: function(data){}
          });
        }
      }

      if (window.addEventListener) {
        window.addEventListener("message", listener);
      } else {
        // IE8
        window.attachEvent("onmessage", listener);
      }
  </script>
</body>

</html>
