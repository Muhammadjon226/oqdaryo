
async function main () {

    let site = window.location.href
    const response = await fetch(site,{
        method:"VIEW",
    })

    const datas = await response.json()
	console.log(datas)
    
    info = 0
    for (const data of datas) {
        info+=1
        a = document.createElement("a")
		a.setAttribute("href", site +'/'+info)
        spanS = document.createElement("span")
        spanS.classList.add("mystyle")
        spanS.textContent = "Name: "+data.name
        spanR = document.createElement("span")
        spanR.textContent = "Email: "+data.email      
        a.append(spanS,spanR)
        app.appendChild(a)
    }

}

main()
