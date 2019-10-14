<html>
<head>
    <title>Buy now</title>

    <script src="//code.jquery.com/jquery-1.11.1.min.js"></script>

    <link href="//netdna.bootstrapcdn.com/bootstrap/3.0.1/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">

    <style>
            @import url("//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css");

            /*********************************************
                            Theme Elements
            *********************************************/

            .gold{
                color: #FFBF00;
            }

            /*********************************************
                                PRODUCTS
            *********************************************/

            .product{
                border: 1px solid #dddddd;
                height: 321px;
            }

            .product>img{
                max-width: 230px;
            }

            .product-rating{
                font-size: 20px;
                margin-bottom: 25px;
            }

            .product-title{
                font-size: 20px;
            }

            .product-desc{
                font-size: 14px;
            }

            .product-price{
                font-size: 22px;
            }

            .product-stock{
                color: #74DF00;
                font-size: 20px;
                margin-top: 10px;
            }

            .product-info{
                margin-top: 50px;
            }

            /*********************************************
                                VIEW
            *********************************************/

            .content-wrapper {
                max-width: 1140px;
                background: #fff;
                margin: 0 auto;
                margin-top: 25px;
                margin-bottom: 10px;
                border: 0px;
                border-radius: 0px;
            }

            .container-fluid{
                max-width: 1140px;
                margin: 0 auto;
            }

            .view-wrapper {
                float: right;
                max-width: 70%;
                margin-top: 25px;
            }

            .container {
                padding-left: 0px;
                padding-right: 0px;
                max-width: 100%;
            }

            .service1-items {
                padding: 0px 0 0px 0;
                float: left;
                position: relative;
                overflow: hidden;
                max-width: 100%;
                height: 321px;
                width: 130px;
            }

            .service1-item {
                height: 107px;
                width: 120px;
                display: block;
                float: left;
                position: relative;
                padding-right: 20px;
                border-right: 1px solid #DDD;
                border-top: 1px solid #DDD;
                border-bottom: 1px solid #DDD;
            }

            .service1-item > img {
                max-height: 110px;
                max-width: 110px;
                opacity: 0.6;
                transition: all .2s ease-in;
                -o-transition: all .2s ease-in;
                -moz-transition: all .2s ease-in;
                -webkit-transition: all .2s ease-in;
            }

            .service1-item > img:hover {
                cursor: pointer;
                opacity: 1;
            }

            .service-image-left {
                padding-right: 50px;
            }

            .service-image-right {
                padding-left: 50px;
            }

            .service-image-left > center > img,.service-image-right > center > img{
                max-height: 155px;
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

<div class="container-fluid">
    <div class="content-wrapper">
        <div class="item-container">
            <div class="container">
                <div class="col-md-12">
                    <div class="product col-md-3 service-image-left">

                        <center>
                            <img id="item-display" src="https://www.corsair.com/medias/sys_master/images/images/h81/h97/8843459657758/-CMPSU-600G-Gallery-gs600-001.png" alt=""></img>
                        </center>
                    </div>

                    <div class="container service1-items col-sm-2 col-md-2 pull-left">
                        <center>
                            <a id="item-1" class="service1-item">
                                <img src="https://www.corsair.com/corsairmedia/sys_master/productcontent/CP-9020063-NA-GS600_PSU_01.png" alt=""></img>
                            </a>
                            <a id="item-2" class="service1-item">
                                <img src="https://www.corsair.com/corsairmedia/sys_master/productcontent/CP-9020063-NA-GS600_PSU_01.png" alt=""></img>
                            </a>
                            <a id="item-3" class="service1-item">
                                <img src="https://www.corsair.com/corsairmedia/sys_master/productcontent/CP-9020063-NA-GS600_PSU_01.png" alt=""></img>
                            </a>
                        </center>
                    </div>
                </div>

                <div class="col-md-7">
                    <div class="product-title">{{.Name}}</div>
                    <div class="product-desc">The Corsair Gaming Series GS600 is the ideal price/performance choice for mid-spec gaming PC</div>
                    <div class="product-rating"><i class="fa fa-star gold"></i> <i class="fa fa-star gold"></i> <i class="fa fa-star gold"></i> <i class="fa fa-star gold"></i> <i class="fa fa-star-o"></i> </div>
                    <hr>
                    <div class="product-price">$ {{.Price}}</div>
                    <div class="product-stock">In Stock</div>
                    <hr>
                    <div class="btn-group cart">
                    {{ if .InCart }}
                        <a type="button"  href="/cart" class="btn btn-primary">
                            In cart (Go to cart)
                        </a>
                    {{ else }}
                        <button type="button" id="to-cart" class="btn btn-success" data-product-id={{.Id}}>
                            Add to cart
                        </button>
                    {{ end }}
                    </div>
                </div>
            </div>
        </div>
        <div class="container-fluid">
            <div class="col-md-12 product-info">
                <ul id="myTab" class="nav nav-tabs nav_tabs">

                    <li class="active"><a href="#service-one" data-toggle="tab">DESCRIPTION</a></li>
                    <li><a href="#service-two" data-toggle="tab">PRODUCT INFO</a></li>
                    <li><a href="#service-three" data-toggle="tab">REVIEWS</a></li>

                </ul>
                <div id="myTabContent" class="tab-content">
                    <div class="tab-pane fade in active" id="service-one">

                        <section class="container product-info">
                            The Corsair Gaming Series GS600 power supply is the ideal price-performance solution for building or upgrading a Gaming PC. A single +12V rail provides up to 48A of reliable, continuous power for multi-core gaming PCs with multiple graphics cards. The ultra-quiet, dual ball-bearing fan automatically adjusts its speed according to temperature, so it will never intrude on your music and games. Blue LEDs bathe the transparent fan blades in a cool glow. Not feeling blue? You can turn off the lighting with the press of a button.

                            <h3>Corsair Gaming Series GS600 Features:</h3>
                            <li>It supports the latest ATX12V v2.3 standard and is backward compatible with ATX12V 2.2 and ATX12V 2.01 systems</li>
                            <li>An ultra-quiet 140mm double ball-bearing fan delivers great airflow at an very low noise level by varying fan speed in response to temperature</li>
                            <li>80Plus certified to deliver 80% efficiency or higher at normal load conditions (20% to 100% load)</li>
                            <li>0.99 Active Power Factor Correction provides clean and reliable power</li>
                            <li>Universal AC input from 90~264V — no more hassle of flipping that tiny red switch to select the voltage input!</li>
                            <li>Extra long fully-sleeved cables support full tower chassis</li>
                            <li>A three year warranty and lifetime access to Corsair’s legendary technical support and customer service</li>
                            <li>Over Current/Voltage/Power Protection, Under Voltage Protection and Short Circuit Protection provide complete component safety</li>
                            <li>Dimensions: 150mm(W) x 86mm(H) x 160mm(L)</li>
                            <li>MTBF: 100,000 hours</li>
                            <li>Safety Approvals: UL, CUL, CE, CB, FCC Class B, TÜV, CCC, C-tick</li>
                        </section>

                    </div>
                    <div class="tab-pane fade" id="service-two">

                        <section class="container">

                        </section>

                    </div>
                    <div class="tab-pane fade" id="service-three">

                    </div>
                </div>
                <hr>
            </div>
        </div>
    </div>
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
    <script src="//netdna.bootstrapcdn.com/bootstrap/3.0.1/js/bootstrap.min.js"></script>
    <script>
        let button = $('#to-cart')
        button.on('click', (e) => {
            e.preventDefault();
            let self = $(this);
            button.addClass('disabled');
            let productId = button.data('product-id');

            $.ajax({
              type: "POST",
              url: "/add-to-cart",
              data: {productId: productId},
              cache: false,
              success: function(data){
                 $('.cart').html('')
                 $('.cart').html('<a type="button"  href="/cart" class="btn btn-primary">In cart (Go to cart)</a>')
              }
            });
        })
    </script>
</body>
</html>