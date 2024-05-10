package server

const page = `
<!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<title>Orders</title>
	<meta name="description" content="">
	<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<h7>Total: {{.TotalPages}}</h7>
<ul>
	<div style="white-space: pre;">{{.Content}}</div>
</ul>
<div>
	{{if gt .TotalPages 1}}
		{{if gt .CurrPageNum 1}}
			<a href="/?id={{.PrevPageNum}}">Prev</a>{{end}}{{if lt .CurrPageNum .TotalPages}}
			<a href="/?id={{.NextPageNum}}">Next</a>
			<a href="/?id={{.LastPageNum}}">Last</a>
		{{end}}
	{{end}}
</div>
</body>
</html>`

const outOfRangePage = `
<!doctype html>
<html>
<head>
	<meta charset="utf-8">
	<title>Orders</title>
	<meta name="description" content="">
	<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<h7>Index out of range</h7>
<div>
	<a href="/?id=1">Go to first order</a>
</div>
</body>
</html>`
