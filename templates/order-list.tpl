<html>
    <head>
        <title>Cart</title>
        <link href="//netdna.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">

        <style type="text/css">
            @import url("//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css");

            .table>tbody>tr>td, .table>tfoot>tr>td{
                vertical-align: middle;
            }
            @media screen and (max-width: 600px) {
                table#cart tbody td .form-control{
                    width:20%;
                    display: inline !important;
                }
                .actions .btn{
                    width:36%;
                    margin:1.5em 0;
                }

                .actions .btn-info{
                    float:left;
                }
                .actions .btn-danger{
                    float:right;
                }

                table#cart thead { display: none; }
                table#cart tbody td { display: block; padding: .6rem; min-width:320px;}
                table#cart tbody tr td:first-child { background: #333; color: #fff; }
                table#cart tbody td:before {
                    content: attr(data-th); font-weight: bold;
                    display: inline-block; width: 8rem;
                }



                table#cart tfoot td{display:block; }
                table#cart tfoot td .btn{display:block;}

            }
        </style>
    </head>
    <body>

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
    <!------ Include the above in your HEAD tag ---------->

    <div class="container">
        {{ range $key, $value := . }}
        <table id="cart" class="table table-hover table-condensed">
            <thead>
            <tr>
                <th style="width:10%">Order Id</th>
                <th style="width:50%">Product</th>
                <th style="width:10%">Price</th>
                <th style="width:22%" class="text-center">Status</th>
                <th style="width:20%" class="text-center">Actions</th>
            </tr>
            </thead>
            <tbody>
            <tr>
                <td data-th="Quantity">{{$value.Id}}</td>
                <td data-th="Product">
                    <div class="row">
                        <div class="col-sm-2 hidden-xs"><img src="https://www.corsair.com/corsairmedia/sys_master/productcontent/CP-9020063-NA-GS600_PSU_01.png" alt="..." class="img-responsive"/></div>
                        <div class="col-sm-10">
                            <h4 class="nomargin">{{$value.Product.Name}}</h4></div>
                    </div>
                </td>
                <td data-th="Price">$ {{$value.Product.Price}}</td>
                <td data-th="Subtotal" class="text-center">{{$value.Status}}</td>
                {{if eq $value.Status "approved"}}
                    <td><a href='/pay/refund/{{$value.Id}}'>Refund</a></td>
                {{end}}
            </tr>
            </tbody>
        </table>
        {{end}}
    </div>

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
    <script src="//code.jquery.com/jquery-1.11.1.min.js"></script>

    <script src="//netdna.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
    </body>
</html>