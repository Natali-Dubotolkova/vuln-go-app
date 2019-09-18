<html>
<link rel="stylesheet" href="./assets/css/style.css" type="text/css">
<head>
<title></title>
</head>
<body>
<div class="title">
<h1>Vulnerability Apps</h1>
</div>
<div class="register_form">
<h2>New User Register</h2>
<form action="/new" method="post">
    <p>Mail-Address:<input type="email" name="mail" required></p>
    <p>Name:<input type="text" name="name" onsubmit="return validateForm()" required></p>
    <p>Age:<input type="number" name="age" min="0" max="125" pattern=" 0+\.[0-9]*[1-9][0-9]*$" required></p>
    <br>
    <p>Password:<input type="password" name="passwd" id="passwd" required></p>
    <p>Confirm Password:<input type="password" name="confirm" oninput="CheckPassword(this)" required></p>
    <input type="submit" value="Register">
</form>
</div>
</body>
</html>

<style type="text/css">
input[type="number"]::-webkit-outer-spin-button,
input[type="number"]::-webkit-inner-spin-button {
     -webkit-appearance: none;
     margin: 0;
}

input[type="number"] {
     -moz-appearance:textfield;
}
</style>

<script>
	function CheckPassword(confirm){
		var input1 = passwd.value;
		var input2 = confirm.value;
		if(input1 != input2){
			confirm.setCustomValidity("Password not matched!!");
		}else{
			confirm.setCustomValidity('');
		}
	}
</script>
