
sendBtn.onclick = async function () {
	site = window.location.href
	console.log(site)
	
	if (sender_name.value != '' && sender_email.value != '' && sender_comments.value!=''){
		
		bodyInfo = JSON.stringify({
			name: document.getElementById("sender_name").value,
    		email: document.getElementById("sender_email").value,
    		comment: document.getElementById("sender_comments").value
		})
		console.log(bodyInfo)
	    let response = await fetch(site, {
	        method: 'POST',
	        body: bodyInfo,
	    })   
	    if (response.status == 201) {

	        console.log('Created!', response)
	    }
	}
}   