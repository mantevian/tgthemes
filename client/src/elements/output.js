import TgtgElement from "./element";

export default class OutputElement extends TgtgElement {
	static observedAttributes = ["platform", "info", "href", "file-name", "can-copy", "params"];

	render() {
		let platform = this.getAttribute("platform");
		let info = this.getAttribute("info");
		let href = this.getAttribute("href");
		let fileName = this.getAttribute("file-name");
		let canCopy = this.hasAttribute("can-copy");
		let params = this.getAttribute("params");

		this.innerHTML = /*html*/`
			<h3>${platform}</h3>

			<p class="info">${info}</p>

			<div class="actions">
				<p>
					<a
						${href ? `href=${href}` : ""}
						${fileName ? `download=${fileName}` : ""}
					>
						download
					</a>
				</p>
				<button ${canCopy ? "" : "disabled"}>copy code</button>
			</div>
		`;

		this.querySelector("button").addEventListener("click", _ => {
			fetch(`/${platform}?${params}`, {
				method: "GET"
			})
				.then(res => res.text())
				.then(text => {
					navigator.clipboard.writeText(text);
				});
		});
	}
}