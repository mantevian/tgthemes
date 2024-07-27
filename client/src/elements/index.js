import OutputElement from "./output";

const prefix = "tgtg";

const elements = {
	"output": OutputElement,
};

for (let e of Object.entries(elements)) {
	customElements.define(`${prefix}-${e[0]}`, e[1]);
}