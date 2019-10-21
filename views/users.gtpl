<html>
<head>
</head>
<div id="header_title">
<h1>Profile</h1>
</div>

</nav>
</header>
<link rel="stylesheet" href="./assets/css/styleClear.css" type="text/css"> 
<body>
<div class="center">
</div>

<div class="box">
    <div class="profileImage">
    {{ if .Image }}
    <img src="./assets/img/{{.Image}}" width="400" height="300">
    <h2>{{.Word}}</h2>
    {{ else }}
    <img src="./assets/img/noimage.png" width="400" height="300">
    {{end}}
    </div>
    <div class="profileBox">
    <h1>Name : <t> <i>{{.UserName}}</i></h1>
    <h2>Age : <t> {{.Age}}</h2>
    <h2>Mail : <t> {{.Mail}}</h2>
    <h2>Address : <t> {{.Address}}</h2>
    <h2>Favorite Animal : <t> {{.Animal}}</h2>
    </div>
    <a href="/profile/edit"><button type="button" class="button">Edit</button></a>
    <a href="/profile/edit/image"><button type="button" class="button_changeImage">Change Image</button></a>
    <a href="/profile/changepasswd"><button type="button" class="button_changepasswd">Change Password</button></a>
</div>

</body>
</html>
