export default class TgtgElement extends HTMLElement {
	constructor() {
		super();
	}

	connectedCallback() {
		this.render();
	}

	render() {

	}

	attributeChangedCallback(name, oldValue, newValue) {
		this.render();
	}
}