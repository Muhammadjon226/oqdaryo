package main

var POSTINFO = `
	INSERT INTO comments (name, email, comment)	values ($1, $2, $3)
`
var GETALLINFO = `
	SELECT name, email, comment FROM comments
`