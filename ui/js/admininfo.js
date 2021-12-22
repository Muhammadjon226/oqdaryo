async function main () {
    const site = window.location.href

    const response = await fetch(site,{
        method:"VIEW",
    })
    const data = await response.json()
    

    sfname.textContent = data.name
    sfemail.textContent = data.email
    comment.textContent = data.comment

    console.log(data)
   
}
main()
