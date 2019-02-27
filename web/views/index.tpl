<!DOCTYPE html>
<html>
  <head>
<style>
    html, body{
        width: 100%;
        height: 100%;
    }
    #map{
        width: 100%;
        height: 100%;
    }
</style>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/css/bootstrap.min.css">

<!-- jQuery library -->
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

<!-- Latest compiled JavaScript -->
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/js/bootstrap.min.js"></script>


<!---Multiple Select----->
<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/multiple-select/1.2.2/multiple-select.min.css">
<!-- Latest compiled and minified JavaScript -->
<script src="//cdnjs.cloudflare.com/ajax/libs/multiple-select/1.2.2/multiple-select.min.js"></script>
</head>
  <body>
<span class="metadata-marker" style="display: none;" data-region_tag="html-body"></span>    
  <h1>法拍屋價格查詢</h1>
    <!--The div element for the map -->

    <form class="form-inline" action="/action_page.php">
      <div class="form-group">
        <label for="city">縣市:</label>
        <select class="form-control" id="sel1">
            <option selected>台北市</option>
            <option>新北市</option>
          </select>

      </div>
      <div class="form-group">
        <label for="area">地區:</label>
        <select id="select" multiple="multiple">
            <option value="1">January</option>
            <option value="12">December</option>
        </select>
      </div>
      <div class="checkbox">
        <label><input type="checkbox"> Remember me</label>
      </div>
      <button type="submit" class="btn btn-default">Submit</button>
    </form>
    <br/>
    <br/>
    <div id="map" display = "width:70%;"></div>
    <script>

$('#select').multipleSelect();

//<span class="metadata-marker" style="display: none;" data-region_tag="script-body"></span>
function initMap() {
  // The location of Uluru
//Lat:25.0329693, Lng:121.5654073 maps.LatLng{Lat:25.124722, Lng:121.5081544}
//maps.LatLng{Lat:25.1219282, Lng:121.5027871}

//Lat:25.0485466, Lng:121.5550682
  var uluru = {lat: 25.0485466, lng:121.5550682};
  // The map, centered at Uluru
  var map = new google.maps.Map(
      document.getElementById('map'), {zoom: 17, center: uluru});
  // The marker, positioned at Uluru
  var marker = new google.maps.Marker({position: uluru, map: map});
}
    </script>
    <!--Load the API from the specified URL
    * The async attribute allows the browser to render the page while the API loads
    * The key parameter will contain your own API key (which is not needed for this tutorial)
    * The callback parameter executes the initMap() function
    -->
    <script async defer
    src="https://maps.googleapis.com/maps/api/js?key=AIzaSyAK3QC8cDO94CpT7h-mQeEZNEix4gvDMlo&callback=initMap">
    </script>
  </body>
</html>
