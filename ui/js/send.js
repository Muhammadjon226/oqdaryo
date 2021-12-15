
sendBtn.onclick = async function () {
	
	if (sender_name.value != '' && sender_email.value!='' && sender_comment.value!=''){
		
		let commentInfo = JSON.stringify({
	       
	            name: sender_name.value,
	            email:   sender_email.value,
	            comments: sender_comment.value
	    })

	    let response = await fetch('http://localhost:8080', {
	        method: 'POST',
	        body: commentInfo,
	        
	    })   
	    if (response.status == 201) {

	        console.log('Created!', response)
	    }
	}
}   