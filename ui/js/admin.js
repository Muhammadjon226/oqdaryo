const ordersElement = document.getElementById("ordersElement")

async function main () {
	site = window.location.href

	const response = await fetch(site)

	const orders = await response.json()

	for (const order of orders) {

		const liElement = document.createElement("li")

		ordersElement.appendChild(liElement)

		const pElement = document.createElement("p")

		liElement.appendChild(pElement)

		const userElement = document.createElement("span")
		userElement.textContent = order.Name

		const emailElement = document.createElement("span")
		emailElement.textContent = order.Email
		
		const commentElement = document.createElement("span")
		commentElement.textContent = order.Email

		pElement.appendChild(userElement)
		pElement.appendChild(emailElement)
		pElement.appendChild(commentElement)
	}
}