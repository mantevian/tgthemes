const form = document.querySelector("form");
const formInputs = form.querySelectorAll("input");

function updateOutputs() {
	let values = htmx.values(form);

	let params = new URLSearchParams(
		Object.fromEntries([...values.entries()]),
	).toString();

	let androidOutput = htmx.find("tgtg-output[platform='android']");
	androidOutput.setAttribute("can-copy", "");
	androidOutput.setAttribute("params", params);
	androidOutput.setAttribute("href", `/android?decimal=true&${params}`);
	androidOutput.setAttribute("file-name", `${values.get("name") || "my_theme"}.attheme`);

	let desktopOutput = htmx.find("tgtg-output[platform='desktop']");
	desktopOutput.setAttribute("params", params);
	desktopOutput.setAttribute("can-copy", "");
}

updateOutputs();

formInputs.forEach(input => {
	input.addEventListener("input", e => {
		updateOutputs();
	});
});

const colorInputs = document.querySelectorAll("input[type=color]");

colorInputs.forEach(input => {
	input.addEventListener("input", e => {
		document.documentElement.style.setProperty(`--${e.currentTarget.name}`, e.target.value);
	});
});